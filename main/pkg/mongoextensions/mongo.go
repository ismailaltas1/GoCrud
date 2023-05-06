package mongoextensions

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func GetDatabase(uri string, databaseName string) (db *mongo.Database) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		return db
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Mongo: mongo client couldn't connect with background context: %v", err)
		return db
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return db
	}

	db = client.Database(databaseName)

	return db
}
