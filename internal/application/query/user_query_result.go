package query

import "github.com/artsadert/TaskFlow/internal/application/common"

type UserQueryResult struct {
	Result *common.UserResult
}

type UserQueryListResult struct {
	Result []*common.UserResult
}
