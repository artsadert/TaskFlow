package repository

import (
	"github.com/artsadert/TaskFlow/internal/domain/entities"
	"github.com/google/uuid"
)

type UserRepo interface {
	GetUser(uuid.UUID) (*entities.User, error)
	GetUsers() ([]*entities.User, error)
	CreateUser(*entities.User) error
	UpdateUser(*entities.User) error
	DeleteUser(uuid.UUID) error
}
