package interfaces

import (
	"github.com/artsadert/TaskFlow/internal/application/command"
	"github.com/artsadert/TaskFlow/internal/application/query"
	"github.com/google/uuid"
)

type MovieService interface {
	CreateMovie(*command.CreateMovieCommand) (*command.CreateMovieCommandResult, error)
	DeleteMovie(*command.DeleteMovieCommand) (*command.DeleteMovieCommandResult, error)
	UpdateMovie(*command.UpdateMovieCommand) (*command.UpdateMovieCommandResult, error)
	GetMovie(uuid.UUID) (*query.MovieQueryResult, error)
	GetMovies() (*query.MovieQueryListResult, error)
}
