package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.elastic.co/apm/module/apmechov4"
	"interaction-api/api/middlewares"
	"interaction-api/config"
	"interaction-api/module"
	"net/http"
	"strings"
)

func RegisterPath() *echo.Echo {
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://examplePublicKey@o0.ingest.sentry.io/0",
		//Dsn: config.AppConfig.AdsLink,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(apmechov4.Middleware())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${time_rfc3339} [${method}] ${uri} ; remote_ip=${remote_ip} ; error=${error}\n",
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Once it's done, you can attach the handler as one of your middleware
	//e.Use(sentryecho.New(sentryecho.Options{}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if hub := sentryecho.GetHubFromContext(ctx); hub != nil {
				hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
			}
			return next(ctx)
		}
	})

	// get jwt secret
	var jwtSecret string
	if config.AppConfig.JwtRedis.Enable {
		jwtSpec, err_ := module.AppRedis.GetParseCustomClient(context.TODO(), module.AppRedis.JwtRedisClient, config.AppConfig.JwtRedis.Key)
		if err_ != nil {
			panic(err_)
		}
		if val, ok := jwtSpec["secret"]; ok {
			if val_, ok_ := val.(string); ok_ {
				jwtSecret = val_
			} else {
				panic(err_)
			}
		} else {
			panic(errors.New("JWT not configured"))
		}
	} else {
		jwtSecret = config.AppConfig.JWTSecret
	}

	// initialize repo, service, & controller
	checkerHandler := NewCheckerController()
	likeHandler := NewLikeController()
	viewHandler := NewViewController()
	followHandler := NewFollowController()
	consumerHandler := NewConsumerController()

	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	// equivalent of auto.offset.reset: earliest
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               strings.Split(config.AppConfig.KafkaConfig.Brokers, ","),
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         "interaction-group",
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	messages, err := subscriber.Subscribe(context.Background(), "like")
	if err != nil {
		panic(err)
	}

	unmessages, err := subscriber.Subscribe(context.Background(), "unlike")
	if err != nil {
		panic(err)
	}

	viewsMsg, err := subscriber.Subscribe(context.Background(), "view")
	if err != nil {
		panic(err)
	}

	followMsg, err := subscriber.Subscribe(context.Background(), "follow")
	if err != nil {
		panic(err)
	}

	unFollowMsg, err := subscriber.Subscribe(context.Background(), "unfollow")
	if err != nil {
		panic(err)
	}

	go consumerHandler.Process(messages)
	go consumerHandler.UnProcess(unmessages)
	go consumerHandler.Views(viewsMsg)
	go consumerHandler.Follow(followMsg)
	go consumerHandler.UnFollow(unFollowMsg)

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong")
	})

	e.GET("/health-check", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	basePathUrlV1 := fmt.Sprintf("%sv1", config.AppConfig.BasePathURL)
	g := e.Group(basePathUrlV1)
	{
		appChecker := g.Group("/checker")
		{
			appChecker.GET("", checkerHandler.TestChecker)
		}

		like := g.Group("/likes")
		middlewares.AuthenticationMiddleware(like, jwtSecret)
		{
			like.POST("/like", likeHandler.AddLike, middlewares.ValidateUserJWT)
			like.POST("/unlike", likeHandler.Unlike, middlewares.ValidateUserJWT)
			like.GET("/total/:service/:type/:id", likeHandler.TotalLike)
		}

		views := g.Group("/views")
		middlewares.AuthenticationMiddleware(views, jwtSecret)
		{
			views.POST("/add", viewHandler.AddView)
			views.GET("/total/:service/:type/:id", viewHandler.GetTotalView)
		}

		follow := g.Group("/follows")
		middlewares.AuthenticationMiddleware(follow, jwtSecret)
		{
			follow.POST("/follow", followHandler.AddFollow)
			follow.POST("/unfollow", followHandler.UnFollow, middlewares.ValidateUserJWT)
			follow.GET("/isfollow/:id", followHandler.StatusFollow, middlewares.ValidateUserJWT)
			follow.GET("/total/:id", followHandler.GetTotalFollow)
		}

	}

	return e
}
