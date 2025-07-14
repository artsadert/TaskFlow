package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Connection struct {
	client *mongo.Client
}

func NewConnection() (*Connection, error) {
	mongo_url := fmt.Sprintf("mongodb://%s:%s@localhost:27017", os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD"))

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongo_url).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, err
	}
	return &Connection{client: client}, nil
}

func (c Connection) getCollection(name string) (*mongo.Collection, error) {
	mongo_url := fmt.Sprintf("mongodb://%s:%s@localhost:27017", os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD"))

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongo_url).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("TaskFlow").Collection(name)

	return collection, nil

}

func (c *Connection) GetCollectionUsers() (*mongo.Collection, error) {
	return c.getCollection("users")
}

func GetCollectionMovies() (*mongo.Collection, error) {
	return getCollection("movies")

}

func GetCollectionRevisions() (*mongo.Collection, error) {
	return getCollection("revisions")

}
