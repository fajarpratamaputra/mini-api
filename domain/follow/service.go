package follow

import (
	"context"
	"encoding/json"
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

func (s *service) Follow(ctx context.Context, interaction domain.InteractionFollow) error {
	return s.kafkaPub.PublishFollow(ctx, "follow", interaction)
}

func (s *service) UnFollow(ctx context.Context, interaction domain.InteractionFollow) error {
	return s.kafkaPub.PublishFollow(ctx, "unfollow", interaction)
}

func (s *service) InsertFollowSvc(ctx context.Context, payload message.Payload) error {
	//Unmarshal message to Interaction Struct
	var intr domain.InteractionFollow
	err := json.Unmarshal(payload, &intr)
	if err != nil {
		return err
	}

	//Insert Data to Mongo DB
	if err = s.repository.InsertFollow(ctx, intr); err != nil {
		return err
	}

	return nil
}

func (s *service) ExtractFollow(ctx context.Context, payload message.Payload) error {
	//Unmarshal message to Interaction Struct
	var intr domain.DeleteFollowMongo
	err := json.Unmarshal(payload, &intr)
	if err != nil {
		return err
	}

	//Remove Data from Mongo DB
	if err = s.repository.RemoveFollowMongo(ctx, intr); err != nil {
		return err
	}

	return nil
}

func (s *service) GetStatusFollow(ctx context.Context, userId, followId int) (bool, error) {
	return s.repository.IsFollow(ctx, userId, followId)
}

func (s *service) TotalFollow(ctx context.Context, get domain.RequestGetFollow) (int, error) {
	return s.repository.TotalFollows(ctx, get)
}

func (s *service) TotalFollower(ctx context.Context, get domain.RequestGetFollower) (int, error) {
	return s.repository.TotalFollower(ctx, get)
}
