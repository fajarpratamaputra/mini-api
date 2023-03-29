package like

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill/message"
	"interaction-api/domain"
	"interaction-api/module"
)

type service struct {
	kafkaPub   *module.KafkaPub
	repository Repository
	rds        *module.RedisWrapper
}

func NewService(pub *module.KafkaPub, repository Repository, rds *module.RedisWrapper) Service {
	return &service{
		kafkaPub:   pub,
		repository: repository,
		rds:        rds,
	}
}

func (s *service) Like(ctx context.Context, interaction domain.Interaction) error {
	return s.kafkaPub.PublishLike(ctx, interaction.Action, interaction)
}

func (s *service) Unlike(ctx context.Context, interaction domain.Interaction) error {
	return s.kafkaPub.PublishLike(ctx, interaction.Action, interaction)
}

func (s *service) InsertLike(ctx context.Context, payload message.Payload) error {
	//Unmarshal message to Interaction Struct
	var intr domain.Interaction
	err := json.Unmarshal(payload, &intr)
	if err != nil {
		return err
	}

	//Insert Data to Mongo DB
	if err = s.repository.AddLikeMongo(ctx, intr); err != nil {
		return err
	}

	key := fmt.Sprintf("master:%s:content_type:%s:{%d}:like", intr.Service, intr.ContentType, intr.ContentID)

	//Increment Redis
	if err = s.rds.Increment(ctx, key); err != nil {
		return err
	}

	return nil
}

func (s *service) ExtractLike(ctx context.Context, payload message.Payload) error {
	//Unmarshal message to Interaction Struct
	var intr domain.DeleteMongo
	err := json.Unmarshal(payload, &intr)
	if err != nil {
		return err
	}

	//Remove Data from Mongo DB
	if err = s.repository.RemoveLikeMongo(ctx, intr); err != nil {
		return err
	}

	key := fmt.Sprintf("master:%s:content_type:%s:{%d}:like", intr.Service, intr.ContentType, intr.ContentID)

	//Decrement Redis Like
	if err = s.rds.Decrement(ctx, key); err != nil {
		return err
	}

	return nil
}

func (s *service) TotalLike(ctx context.Context, get domain.RequestGet) (int, error) {
	return s.repository.TotalLike(ctx, get)
}
