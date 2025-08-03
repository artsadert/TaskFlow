package graph

import "github.com/artsadert/TaskFlow/internal/application/interfaces"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	user_service     interfaces.UserService
	movie_service    interfaces.MovieService
	revision_service interfaces.RevisionService
}
