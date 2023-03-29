package view

import (
	"context"
	"fmt"
	"interaction-api/config"
	"interaction-api/domain"
	"interaction-api/module"
	"time"
)

var collection = "views"

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

func (r *Repository) InsertView(ctx context.Context, interaction domain.InteractionView) error {
	interaction.CreatedAt = time.Now()
	result, err := r.DbMongo.InsertOneDocument(ctx, collection, interaction)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return nil
}

func (r *Repository) TotalViews(ctx context.Context, get domain.RequestGet) (int, error) {
	result64, err := r.DbMongo.Count(ctx, collection, get)
	if err != nil {
		return 0, err
	}
	return int(result64), nil
}
