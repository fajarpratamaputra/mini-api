package checker

import (
	"context"
	"errors"
	"interaction-api/config"
	"interaction-api/module"
)

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

func (r *Repository) Bbb() error {
	return errors.New("Bbb")
}

func (r *Repository) CheckDB() error {
	err := r.DbRepository.DB.Error
	return err
}

func (r *Repository) CheckMongoDB(ctx context.Context) error {
	return r.DbMongo.Client.Ping(ctx, nil)
}
