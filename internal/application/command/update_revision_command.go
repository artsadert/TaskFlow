package command

import (
	"github.com/artsadert/TaskFlow/internal/application/common"
	"github.com/google/uuid"
)

type UpdateRevisionCommand struct {
	Id     uuid.UUID
	Status string
	Review string
	Rating int
}

type UpdateRevisionCommandResult struct {
	Result *common.RevisionResult
}
