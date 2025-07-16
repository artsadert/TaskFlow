package mongodb

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestPing(t *testing.T) {
	err := godotenv.Load("./../../../../.env")
	if err != nil {
		t.Errorf("could  not load dotenv!, %s", err.Error())
	}

	connection, err := NewConnection()

	if err != nil {
		t.Errorf("could nnot connect to db, %s", err.Error())
	}

	collection, err := connection.GetCollectionUsers()
	fmt.Println(collection.Name())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := collection.InsertOne(ctx, bson.M{"test": "true"})
	if err != nil {
		t.Errorf("cannot insertone, %s", err.Error())
	}
	fmt.Println(res)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = connection.client.Ping(ctx, nil)
	if err != nil {
		t.Errorf("cannot ping db, %s", err.Error())
	}
}
