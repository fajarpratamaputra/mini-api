package middlewares

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"interaction-api/config"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func AuthenticationMiddleware(g *echo.Group, secret string) {
	g.Use(setBearerRule)
	g.Use(echojwt.WithConfig(echojwt.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(secret),
	}))
	//g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	//	SigningMethod: "HS256",
	//	SigningKey:    []byte(secret),
	//}))
	g.Use(validateJWTClient)
}

func setBearerRule(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if strings.HasPrefix(tokenString, "Bearer") == false {
			c.Request().Header.Set("Authorization", "Bearer "+tokenString)
		}
		return next(c)
	}
}

func validateJWTClient(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_data", map[string]interface{}(claims))
			return next(c)
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
}

type TokenClaim struct {
	VID       int    `json:"vid"`
	Token     string `json:"token"`
	Pl        string `json:"pl"`
	UserAgent string `json:"user_agent"`
	DeviceID  string `json:"device_id"`
	jwt.Claims
}

func DecodeToken(token string) (*TokenClaim, error) {
	claim := new(TokenClaim)
	_, err := jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return claim, nil
}

func ValidateUserJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)
		//claims, ok := token.Claims.(jwt.MapClaims)
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID, err := strconv.Atoi(fmt.Sprintf("%0.f", claims["vid"]))
			if err != nil {
				log.Println("ERROR at ValidateJWTclient: ", err.Error())
			}
			if userID != 0 {
				c.Set("user_data", map[string]interface{}(claims))
				return next(c)
			}
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "Please Log In First!")
	}
}
