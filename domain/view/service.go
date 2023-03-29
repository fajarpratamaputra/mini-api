package view

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
}

func NewService(pub *module.KafkaPub, repository Repository) Service {
	return &service{
		kafkaPub:   pub,
		repository: repository,
	}
}

func (s *service) View(ctx context.Context, interaction domain.InteractionView) error {
	return s.kafkaPub.PublishView(ctx, interaction.Action, interaction)
}

func (s *service) InsertViewSvc(ctx context.Context, payload message.Payload) error {
	var intr domain.InteractionView
	err := json.Unmarshal(payload, &intr)
	if err != nil {
		return err
	}
	return s.repository.InsertView(ctx, intr)
}

func (s *service) TotalView(ctx context.Context, get domain.RequestGet) (int, error) {
	return s.repository.TotalViews(ctx, get)
}
