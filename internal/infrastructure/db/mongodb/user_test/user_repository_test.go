package usertest

import (
	"context"
	"fmt"
	"testing"

	"github.com/artsadert/TaskFlow/internal/domain/entities"
	"github.com/artsadert/TaskFlow/internal/infrastructure/db/mongodb"
	"github.com/artsadert/TaskFlow/internal/infrastructure/db/mongodb/user"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestGetUser(t *testing.T) {
	err := godotenv.Load("./../../../../../.env")
	if err != nil {
		t.Errorf("could not load dotenv!, %s", err)
	}
	connection, err := mongodb.NewConnection()
	if err != nil {
		t.Errorf("cannt connect to db, %s", err.Error())
	}

	collection, err := connection.GetCollectionUsers()
	if err != nil {
		t.Errorf("cann't create collection: %s", err.Error())
	}

	test_user, err := entities.NewUser("arthur", "email")
	if err != nil {
		t.Errorf("cann't create entity user: %s", err.Error())
	}

	fmt.Println(test_user.Id, test_user.Name, test_user.Email, test_user.Create_at, test_user.Update_at)

	_, err = collection.InsertOne(context.TODO(),
		bson.M{
			"uuid":      test_user.Id,
			"name":      test_user.Name,
			"email":     test_user.Email,
			"create_at": test_user.Create_at,
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
	fmt.Println("created", test_user.Create_at)
	fmt.Println("", get_user.Create_at)

	if get_user.Create_at.Equal(test_user.Create_at) {
		t.Errorf("created and written users are not the same!")
	}

}
