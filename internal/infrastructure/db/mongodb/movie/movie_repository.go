package movie

import (
	"context"

	"github.com/artsadert/TaskFlow/internal/domain/entities"
	"github.com/artsadert/TaskFlow/internal/domain/repository"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoMovieRepository struct {
	db *mongo.Collection
}

func NewMongoMovieRepository(db *mongo.Collection) repository.MovieRepo {
	return &MongoMovieRepository{db: db}
}

func (m *MongoMovieRepository) GetMovie(id uuid.UUID) (*entities.Movie, error) {
	var result DBMovie
	err := m.db.FindOne(context.TODO(), bson.M{"uuid": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return fromDBMovie(&result), nil
}

func (m *MongoMovieRepository) GetMovies() ([]*entities.Movie, error) {
	cursor, err := m.db.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	var db_results []*DBMovie
	if err = cursor.All(context.TODO(), &db_results); err != nil {
		return nil, err
	}

	var result []*entities.Movie
	for _, movie := range db_results {
		result = append(result, fromDBMovie(movie))
	}

	return result, nil
}

func (m *MongoMovieRepository) CreateMovie(movie *entities.Movie) error {
	_, err := m.db.InsertOne(context.TODO(), toDBMovie(movie))
	return err
}

func (m *MongoMovieRepository) UpdateMovie(movie *entities.Movie) error {
	db_movie := toDBMovie(movie)

	filter := bson.M{"uuid": db_movie.Uuid}

	update := bson.M{"$set": bson.M{
		"name":        db_movie.Name,
		"year":        db_movie.Year,
		"genre":       db_movie.Genre,
		"description": db_movie.Description,
		"poster_url":  db_movie.Poster_url,
		"update_at":   db_movie.Update_at.String(),
	}}
	_, err := m.db.UpdateOne(context.TODO(), filter, update)

	return err
}

func (m *MongoMovieRepository) DeleteMovie(id uuid.UUID) error {
	_, err := m.db.DeleteOne(context.TODO(), bson.M{"uuid": id})
	return err
}
