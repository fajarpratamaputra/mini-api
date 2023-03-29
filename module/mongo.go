package module

import (
	"context"
	"fmt"
	"go.elastic.co/apm/module/apmmongo/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"interaction-api/config"
)

type DBMongo struct {
	Client *mongo.Client
}

var AppMongo *DBMongo

func init() {
	var dbMongo DBMongo
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", config.AppConfig.MongoConfig.User, config.AppConfig.MongoConfig.Pass, config.AppConfig.MongoConfig.Host, config.AppConfig.MongoConfig.Port)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri), options.Client().SetMonitor(apmmongo.CommandMonitor()))
	if err != nil {
		panic(err)
	}

	dbMongo.Client = client
	AppMongo = &dbMongo
}

func (m *DBMongo) InsertOneDocument(ctx context.Context, col string, value interface{}) (*mongo.InsertOneResult, error) {
	db := m.Client.Database(config.AppConfig.MongoConfig.DB)
	coll := db.Collection(col)
	result, err := coll.InsertOne(ctx, value)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m *DBMongo) RemoveDocument(ctx context.Context, col string, filter interface{}) (*mongo.DeleteResult, error) {
	db := m.Client.Database(config.AppConfig.MongoConfig.DB)
	coll := db.Collection(col)
	result, err := coll.DeleteMany(ctx, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m *DBMongo) Count(ctx context.Context, col string, filter interface{}) (int64, error) {
	db := m.Client.Database(config.AppConfig.MongoConfig.DB)
	coll := db.Collection(col)
	result, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return result, nil
}
