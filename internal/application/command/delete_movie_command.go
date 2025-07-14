package command

import (
	"github.com/artsadert/TaskFlow/internal/application/common"
	"github.com/google/uuid"
)

type DeleteMovieCommand struct {
	Id uuid.UUID
}

type DeleteMovieCommandResult struct {
	Result *common.MovieResult
}
