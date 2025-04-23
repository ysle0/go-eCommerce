package mongodb

import (
	"context"
	"github.com/ysle0/go-starter/internal/cfg"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"log"
)

type MongoDb struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoDb(ctx context.Context, cfg cfg.Config) *MongoDb {
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.Mongo.Uri).SetServerAPIOptions(serverApi)
	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	db := client.Database(cfg.Mongo.Db)
	m := &MongoDb{
		client: client,
		db:     db,
	}

	return m
}
