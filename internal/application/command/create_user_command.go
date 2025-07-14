package command

import "github.com/artsadert/TaskFlow/internal/application/common"

type CreateUserCommand struct {
	// we dont need id cause it will create it later
	Name  string
	Email string
}

type CreateUserCommandResult struct {
	Result *common.UserResult
}
