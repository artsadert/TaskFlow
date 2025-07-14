package query

import "github.com/artsadert/TaskFlow/internal/application/common"

type MovieQueryResult struct {
	Result *common.MovieResult
}

type MovieQueryListResult struct {
	Result []*common.MovieResult
}
