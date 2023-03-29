package view

import (
	"context"
	"github.com/ThreeDotsLabs/watermill/message"
	"interaction-api/domain"
)

type Service interface {
	View(ctx context.Context, interaction domain.InteractionView) error
	InsertViewSvc(ctx context.Context, payload message.Payload) error
	TotalView(ctx context.Context, get domain.RequestGet) (int, error)
}

type Repository interface {
	InsertView(ctx context.Context, interaction domain.InteractionView) error
	TotalViews(ctx context.Context, get domain.RequestGet) (int, error)
}
