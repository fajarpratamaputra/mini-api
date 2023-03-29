package api

import (
	checkerController "interaction-api/api/controller/checker"
	consumerController "interaction-api/api/controller/consumer"
	followController "interaction-api/api/controller/follow"
	likeController "interaction-api/api/controller/likes"
	viewController "interaction-api/api/controller/view"
	"interaction-api/config"
	checkerService "interaction-api/domain/checker"
	consumerService "interaction-api/domain/consumer"
	followService "interaction-api/domain/follow"
	likeService "interaction-api/domain/like"
	viewService "interaction-api/domain/view"
	"interaction-api/module"
	checkerRepository "interaction-api/module/checker"
	consumerRepository "interaction-api/module/consumer"
	followRepository "interaction-api/module/follow"
	likeRepository "interaction-api/module/like"
	viewRepository "interaction-api/module/view"
)

func NewCheckerController() *checkerController.Controller {
	checkerRepo := checkerRepository.NewRepository(module.AppDB, config.AppConfig, module.AppMongo)
	checkerSvc := checkerService.NewService(checkerRepo)
	return checkerController.NewController(checkerSvc)
}

func NewLikeController() *likeController.Controller {
	likeRepo := likeRepository.NewRepository(module.AppDB, config.AppConfig, module.AppMongo)
	kfk := module.Configure()
	likeSvc := likeService.NewService(kfk, likeRepo, module.AppRedis)
	return likeController.NewController(likeSvc)
}

func NewViewController() *viewController.Controller {
	viewRepo := viewRepository.NewRepository(module.AppDB, config.AppConfig, module.AppMongo)
	kfk := module.Configure()
	viewSvc := viewService.NewService(kfk, viewRepo)
	return viewController.NewController(viewSvc)
}

func NewFollowController() *followController.Controller {
	followRepo := followRepository.NewRepository(module.AppDB, config.AppConfig, module.AppMongo)
	kfk := module.Configure()
	followSvc := followService.NewService(kfk, followRepo, module.AppRedis)
	return followController.NewController(followSvc)
}

func NewConsumerController() *consumerController.Controller {
	pub := module.Configure()
	consumerRepo := consumerRepository.NewRepository(module.AppDB, config.AppConfig, module.AppMongo)
	consumerSvc := consumerService.NewService(consumerRepo)
	likeRepo := likeRepository.NewRepository(module.AppDB, config.AppConfig, module.AppMongo)
	likeSvc := likeService.NewService(pub, likeRepo, module.AppRedis)
	viewRepo := viewRepository.NewRepository(module.AppDB, config.AppConfig, module.AppMongo)
	viewSvc := viewService.NewService(pub, viewRepo)
	followRepo := followRepository.NewRepository(module.AppDB, config.AppConfig, module.AppMongo)
	followSvc := followService.NewService(pub, followRepo, module.AppRedis)
	return consumerController.NewController(consumerSvc, viewSvc, likeSvc, followSvc)
}
