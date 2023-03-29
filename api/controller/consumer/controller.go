package consumer

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill/message"
	"interaction-api/domain/consumer"
	"interaction-api/domain/follow"
	"interaction-api/domain/like"
	"interaction-api/domain/view"
	"log"
)

type Controller struct {
	service       consumer.Service
	viewService   view.Service
	likeService   like.Service
	followService follow.Service
}

func NewController(service consumer.Service, viewSvc view.Service, likeSvc like.Service, followSvc follow.Service) *Controller {
	return &Controller{
		service:       service,
		viewService:   viewSvc,
		likeService:   likeSvc,
		followService: followSvc,
	}
}

func (c *Controller) Process(messages <-chan *message.Message) {
	ctx := context.Background()
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))
		if err := c.likeService.InsertLike(ctx, msg.Payload); err != nil {
			fmt.Printf("Error : %v", err)
		}
		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}

func (c *Controller) Views(messages <-chan *message.Message) {
	ctx := context.Background()
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))
		if err2 := c.viewService.InsertViewSvc(ctx, msg.Payload); err2 != nil {
			fmt.Printf("Error : %v", err2)
		}
		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}

func (c *Controller) UnProcess(messages <-chan *message.Message) {
	ctx := context.Background()
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))
		if err := c.likeService.ExtractLike(ctx, msg.Payload); err != nil {
			fmt.Printf("Error : %v", err)
		}
		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}

func (c *Controller) Follow(messages <-chan *message.Message) {
	ctx := context.Background()
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))
		if err2 := c.followService.InsertFollowSvc(ctx, msg.Payload); err2 != nil {
			fmt.Printf("Error : %v", err2)
		}
		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}

func (c *Controller) UnFollow(messages <-chan *message.Message) {
	ctx := context.Background()
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))
		if err2 := c.followService.ExtractFollow(ctx, msg.Payload); err2 != nil {
			fmt.Printf("Error : %v", err2)
		}
		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
