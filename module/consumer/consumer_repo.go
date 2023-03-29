package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill/message"
	"interaction-api/config"
	"interaction-api/domain"
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

func (r *Repository) LikeMongo(payload message.Payload) error {
	var intr domain.Interaction
	err := json.Unmarshal(payload, &intr)
	if err != nil {
		return err
	}
	coll := r.DbMongo.Client.Database(config.AppConfig.MongoConfig.DB).Collection("interactions")
	result, err := coll.InsertOne(context.TODO(), intr)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return nil
}

func (r *Repository) UnLikeMongo(payload message.Payload) error {
	var intr domain.Interaction
	err := json.Unmarshal(payload, &intr)
	if err != nil {
		return err
	}
	intr.Action = "like"
	coll := r.DbMongo.Client.Database(config.AppConfig.MongoConfig.DB).Collection("interactions")
	_, err = coll.DeleteMany(context.TODO(), intr)
	if err != nil {
		return err
	}

	fmt.Printf("Delete Document Succedd \n")
	return nil
}
