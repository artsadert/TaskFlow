package command

import (
	"github.com/artsadert/TaskFlow/internal/application/common"
	"github.com/google/uuid"
)

type UpdateMovieCommand struct {
	Id          uuid.UUID
	Name        string
	Year        int
	Genre       []string
	Description string
	Poster_url  string
}

type UpdateMovieCommandResult struct {
	Result *common.MovieResult
}
