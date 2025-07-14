package query

import "github.com/artsadert/TaskFlow/internal/application/common"

type RevisionQueryResult struct {
	Result *common.RevisionResult
}

type RevisionQueryListResult struct {
	Result []*common.RevisionResult
}
