package follow

import (
	"context"
	"github.com/ThreeDotsLabs/watermill/message"
	"interaction-api/domain"
)

type Service interface {
	Follow(ctx context.Context, interaction domain.InteractionFollow) error
	UnFollow(ctx context.Context, interaction domain.InteractionFollow) error
	InsertFollowSvc(ctx context.Context, payload message.Payload) error
	ExtractFollow(ctx context.Context, payload message.Payload) error
	GetStatusFollow(ctx context.Context, userId, followId int) (bool, error)
	TotalFollow(ctx context.Context, get domain.RequestGetFollow) (int, error)
	TotalFollower(ctx context.Context, get domain.RequestGetFollower) (int, error)
}

type Repository interface {
	InsertFollow(ctx context.Context, interaction domain.InteractionFollow) error
	RemoveFollowMongo(ctx context.Context, interaction domain.DeleteFollowMongo) error
	IsFollow(ctx context.Context, userId, followId int) (bool, error)
	TotalFollows(ctx context.Context, get domain.RequestGetFollow) (int, error)
	TotalFollower(ctx context.Context, get domain.RequestGetFollower) (int, error)
}
