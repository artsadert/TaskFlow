package repository

import (
	"github.com/artsadert/TaskFlow/internal/domain/entities"
	"github.com/google/uuid"
)

type RevisionRepo interface {
	GetRevision(uuid.UUID) (*entities.Revision, error)
	GetRevisions() ([]*entities.Revision, error)
	GetRevisionsByUserId(uuid.UUID) ([]*entities.Revision, error)
	GetRevisionsByMovieId(uuid.UUID) ([]*entities.Revision, error)
	CreateRevision(*entities.Revision) error
	UpdateRevision(*entities.Revision) error
	DeleteRevision(uuid.UUID) error
}
