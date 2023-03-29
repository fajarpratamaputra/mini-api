package like

import (
	"context"
	"fmt"
	"interaction-api/config"
	"interaction-api/domain"
	"interaction-api/module"
	"time"
)

var collection = "likes"

type Repository struct {
	DbRepository *module.DBRepository
	config       *config.Config
	DbMongo      *module.DBMongo
}

func NewRepository(db *module.DBRepository, cfg *config.Config, dbMongo *module.DBMongo) *Repository {
	return &Repository{
		DbRepository: db,
		config:       cfg,
		DbMongo:      dbMongo,
	}
}

func (r *Repository) AddLikeMongo(ctx context.Context, interaction domain.Interaction) error {
	interaction.CreatedAt = time.Now()
	result, err := r.DbMongo.InsertOneDocument(ctx, collection, interaction)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return nil
}

func (r *Repository) RemoveLikeMongo(ctx context.Context, interaction domain.DeleteMongo) error {
	result, err := r.DbMongo.RemoveDocument(ctx, collection, interaction)
	if err != nil {
		return err
	}

	fmt.Printf("Success Delete %v Documents\n", result.DeletedCount)
	return nil
}

func (r *Repository) TotalLike(ctx context.Context, get domain.RequestGet) (int, error) {
	result64, err := r.DbMongo.Count(ctx, collection, get)
	if err != nil {
		return 0, err
	}
	return int(result64), nil
}
