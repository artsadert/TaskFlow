package user

import (
	"context"

	"github.com/artsadert/TaskFlow/internal/domain/entities"
	"github.com/artsadert/TaskFlow/internal/domain/repository"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoUserRepository struct {
	db *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Collection) repository.UserRepo {
	return &MongoUserRepository{db: db}
}

func (u *MongoUserRepository) GetUser(id uuid.UUID) (*entities.User, error) {
	var result DBUser
	err := u.db.FindOne(context.TODO(), bson.M{"uuid": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return fromDBUser(&result), nil

}

func (u *MongoUserRepository) GetUsers() ([]*entities.User, error) {
	cursor, err := u.db.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var db_result []DBUser
	if err = cursor.All(context.TODO(), &db_result); err != nil {
		return nil, err
	}

	var result []*entities.User
	for _, dbuser := range db_result {
		result = append(result, fromDBUser(&dbuser))
	}

	return result, nil
}

func (u *MongoUserRepository) CreateUser(user *entities.User) error {
	_, err := u.db.InsertOne(context.TODO(), toDBUser(user))

	return err
}

func (u *MongoUserRepository) UpdateUser(user *entities.User) error {
	db_user := toDBUser(user)

	filter := bson.M{"uuid": db_user.Uuid}

	update := bson.M{"$set": bson.M{"name": db_user.Name, "email": db_user.Email, "update_at": db_user.Update_at.String()}}
	_, err := u.db.UpdateOne(context.TODO(), filter, update)

	return err
}

func (u *MongoUserRepository) DeleteUser(id uuid.UUID) error {
	filter := bson.M{"uuid": id}
	_, err := u.db.DeleteOne(context.TODO(), filter)

	return err
}
