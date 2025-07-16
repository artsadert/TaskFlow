package mongodb

import (
	"fmt"
	//"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Connection struct {
	client *mongo.Client
}

func NewConnection() (*Connection, error) {
	//mongo_url := fmt.Sprintf("mongodb://%s:%s@localhost:27017/", os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD"))
	mongo_url := fmt.Sprintf("mongodb://admin:12345678@localhost:27017/TaskFlow?authSource=admin")
	fmt.Println(mongo_url)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongo_url).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, err
	}
	return &Connection{client: client}, nil
}

func (c Connection) getCollection(name string) (*mongo.Collection, error) {
	collection := c.client.Database("TaskFlow").Collection(name)

	return collection, nil

}

func (c Connection) GetCollectionUsers() (*mongo.Collection, error) {
	return c.getCollection("users")
}

func (c Connection) GetCollectionMovies() (*mongo.Collection, error) {
	return c.getCollection("movies")

}

func (c Connection) GetCollectionRevisions() (*mongo.Collection, error) {
	return c.getCollection("revisions")

}
