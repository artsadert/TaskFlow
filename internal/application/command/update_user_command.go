package command

import (
	"github.com/artsadert/TaskFlow/internal/application/common"
	"github.com/google/uuid"
)

type UpdateUserCommand struct {
	Id    uuid.UUID
	Name  string
	Email string
}

type UpdateUserCommandResult struct {
	Result *common.UserResult
}
