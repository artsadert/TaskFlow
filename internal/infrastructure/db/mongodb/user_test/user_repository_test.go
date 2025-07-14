package usertest

import (
	"context"
	"testing"

	"github.com/artsadert/TaskFlow/internal/domain/entities"
	"github.com/artsadert/TaskFlow/internal/infrastructure/db/mongodb"
	"github.com/artsadert/TaskFlow/internal/infrastructure/db/mongodb/user"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestGetUser(t *testing.T) {
	err := godotenv.Load("./../../../../../.env.test")
	if err != nil {
		t.Errorf("could not load dotenv!, %s", err)
	}

	collection, err := mongodb.GetCollectionUsers()
	if err != nil {
		t.Errorf("cann't create collection: %s", err.Error())
	}

	test_user, err := entities.NewUser("arthur", "email")
	if err != nil {
		t.Errorf("cann't create entity user: %s", err.Error())
	}

	_, err = collection.InsertOne(context.TODO(),
		bson.M{
			"uuid":      test_user.Id,
			"name":      test_user.Name,
			"email":     test_user.Email,
			"create_at": test_user.Email,
			"update_at": test_user.Update_at,
		},
	)
	if err != nil {
		t.Errorf("cannot create user mannually: %s", err)
	}

	// actual test
	repo := user.NewMongoUserRepository(collection)
	get_user, err := repo.GetUser(test_user.Id)
	if err != nil {
		t.Errorf("cannot get user from repo: %s", err)
	}

	if get_user.Create_at != test_user.Create_at {
		t.Errorf("created and written users are not the same!")
	}

}
