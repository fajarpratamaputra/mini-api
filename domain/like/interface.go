package like

import (
	"context"
	"github.com/ThreeDotsLabs/watermill/message"
	"interaction-api/domain"
)

type Service interface {
	Like(ctx context.Context, interaction domain.Interaction) error
	Unlike(ctx context.Context, interaction domain.Interaction) error
	InsertLike(ctx context.Context, payload message.Payload) error
	ExtractLike(ctx context.Context, payload message.Payload) error
	TotalLike(ctx context.Context, get domain.RequestGet) (int, error)
}

type Repository interface {
	AddLikeMongo(ctx context.Context, interaction domain.Interaction) error
	RemoveLikeMongo(ctx context.Context, interaction domain.DeleteMongo) error
	TotalLike(ctx context.Context, get domain.RequestGet) (int, error)
}
