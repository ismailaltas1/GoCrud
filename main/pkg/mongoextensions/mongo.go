package mongoextensions

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func NewDatabase(uri, databaseName string) (db *mongo.Database, err error) {
	return getDatabase(uri, databaseName)
}

func NewDatabaseWithCursor(uri, databaseName, collectionName string, filter interface{}) (*mongo.Cursor, error, context.Context) {

	db, err := getDatabase(uri, databaseName)

	if err != nil {
		log.Printf("Mongo: mongo client couldn't connect with new database client: %v", err)
		panic(err.Error())
	}

	collection := db.Collection(collectionName)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		log.Printf("Mongo: cursor coud not initialized: %v", err)

		panic(err.Error())
	}

	return cursor, err, ctx
}

func getDatabase(uri, databaseName string) (db *mongo.Database, err error) {

	client, err := mongo.NewClient(options.
		Client().
		ApplyURI(uri))

	if err != nil {

		return db, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Mongo: mongo client couldn't connect with background context: %v", err)
		return db, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return db, err
	}

	db = client.Database(databaseName)

	return db, err
}
