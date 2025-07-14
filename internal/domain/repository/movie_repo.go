package repository

import (
	"github.com/artsadert/TaskFlow/internal/domain/entities"
	"github.com/google/uuid"
)

type MovieRepo interface {
	GetMovie(uuid.UUID) (*entities.Movie, error)
	GetMovies() ([]*entities.Movie, error)
	CreateMovie(*entities.Movie) error
	UpdateMovie(*entities.Movie) error
	DeleteMovie(uuid.UUID) error
}
