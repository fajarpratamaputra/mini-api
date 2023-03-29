package follow

import (
	"context"
	"fmt"
	"interaction-api/config"
	"interaction-api/domain"
	"interaction-api/module"
	"time"
)

var collection = "follow"

type Repository struct {
	DbRepository *module.DBRepository
	config       *config.Config
	DbMongo      *module.DBMongo
}

type IsFollow struct {
	UserID   int
	FollowTo int
}

func NewRepository(db *module.DBRepository, cfg *config.Config, dbMongo *module.DBMongo) *Repository {
	return &Repository{
		DbRepository: db,
		config:       cfg,
		DbMongo:      dbMongo,
	}
}

func (r *Repository) InsertFollow(ctx context.Context, interaction domain.InteractionFollow) error {
	interaction.CreatedAt = time.Now()
	result, err := r.DbMongo.InsertOneDocument(ctx, collection, interaction)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return nil
}

func (r *Repository) RemoveFollowMongo(ctx context.Context, interaction domain.DeleteFollowMongo) error {
	result, err := r.DbMongo.RemoveDocument(ctx, collection, interaction)
	if err != nil {
		return err
	}

	fmt.Printf("Success Delete %v Documents\n", result.DeletedCount)
	return nil
}

func (r *Repository) TotalFollows(ctx context.Context, get domain.RequestGetFollow) (int, error) {
	result64, err := r.DbMongo.Count(ctx, collection, get)
	if err != nil {
		return 0, err
	}
	return int(result64), nil
}

func (r *Repository) TotalFollower(ctx context.Context, get domain.RequestGetFollower) (int, error) {
	result64, err := r.DbMongo.Count(ctx, collection, get)
	if err != nil {
		return 0, err
	}
	return int(result64), nil
}

func (r *Repository) IsFollow(ctx context.Context, userId, followId int) (bool, error) {
	isF := IsFollow{
		userId,
		followId,
	}
	result64, err := r.DbMongo.Count(ctx, collection, isF)
	if err != nil {
		return false, err
	}

	if result64 <= 0 {
		return false, nil
	}

	return true, nil
}
