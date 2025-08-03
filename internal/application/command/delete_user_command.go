package command

import (
	"github.com/artsadert/TaskFlow/internal/application/common"
	"github.com/google/uuid"
)

type DeleteUserCommand struct {
	Id uuid.UUID
}

type DeleteUserCommandResult struct {
	Result *common.UserResult
}
