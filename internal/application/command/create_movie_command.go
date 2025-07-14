package command

import "github.com/artsadert/TaskFlow/internal/application/common"

type CreateMovieCommand struct {
	Name        string
	Year        int
	Genre       []string
	Description string
	Poster_url  string
}

type CreateMovieCommandResult struct {
	Result *common.MovieResult
}
