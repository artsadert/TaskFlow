package revision

import (
	"context"

	"github.com/artsadert/TaskFlow/internal/domain/entities"
	"github.com/artsadert/TaskFlow/internal/domain/repository"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoRevisionRepo struct {
	db *mongo.Collection
}

func NewMongoRevisionRepo(db *mongo.Collection) repository.RevisionRepo {
	return &MongoRevisionRepo{db: db}
}

func (r *MongoRevisionRepo) GetRevision(id uuid.UUID) (*entities.Revision, error) {
	var result DBRevision
	err := r.db.FindOne(context.TODO(), bson.M{"uuid": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return fromDBRevision(&result), nil
}

func (r *MongoRevisionRepo) GetRevisions() ([]*entities.Revision, error) {
	cursor, err := r.db.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var db_results []*DBRevision
	if err = cursor.All(context.TODO(), &db_results); err != nil {
		return nil, err
	}

	var result []*entities.Revision
	for _, revision := range db_results {
		result = append(result, fromDBRevision(revision))
	}

	return result, nil
}

func (r *MongoRevisionRepo) GetRevisionsByUserId(id uuid.UUID) ([]*entities.Revision, error) {
	cursor, err := r.db.Find(context.TODO(), bson.M{"user_id": id})
	if err != nil {
		return nil, err
	}

	var db_results []*DBRevision
	if err = cursor.All(context.TODO(), &db_results); err != nil {
		return nil, err
	}

	var result []*entities.Revision
	for _, revision := range db_results {
		result = append(result, fromDBRevision(revision))
	}

	return result, nil
}

func (r *MongoRevisionRepo) GetRevisionsByMovieId(id uuid.UUID) ([]*entities.Revision, error) {
	cursor, err := r.db.Find(context.TODO(), bson.M{"movie_id": id})
	if err != nil {
		return nil, err
	}

	var db_results []*DBRevision
	if err = cursor.All(context.TODO(), &db_results); err != nil {
		return nil, err
	}

	var result []*entities.Revision
	for _, revision := range db_results {
		result = append(result, fromDBRevision(revision))
	}

	return result, nil
}

func (r *MongoRevisionRepo) CreateRevision(revision *entities.Revision) error {
	_, err := r.db.InsertOne(context.TODO(), toDBRevision(revision))
	return err
}

func (r *MongoRevisionRepo) UpdateRevision(revision *entities.Revision) error {
	db_revision := toDBRevision(revision)

	filter := bson.M{"uuid": db_revision.Uuid}
	update := bson.M{"$set": bson.M{
		"status":    db_revision.Status,
		"rating":    db_revision.Rating,
		"review":    db_revision.Review,
		"update_at": db_revision.Update_at,
	}}

	_, err := r.db.UpdateOne(context.TODO(), filter, update)

	return err
}

func (r *MongoRevisionRepo) DeleteRevision(id uuid.UUID) error {
	_, err := r.db.DeleteOne(context.TODO(), bson.M{"uuid": id})
	return err
}
