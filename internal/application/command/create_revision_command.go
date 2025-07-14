package command

import (
	"github.com/artsadert/TaskFlow/internal/application/common"
	"github.com/google/uuid"
)

type CreateRevisionCommand struct {
	Rating  int
	Review  string
	UserId  uuid.UUID
	MovieId uuid.UUID
}

type CreateRevisionCommandResult struct {
	Result *common.RevisionResult
}
