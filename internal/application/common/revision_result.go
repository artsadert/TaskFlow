package common

import (
	"time"

	"github.com/google/uuid"
)

type RevisionResult struct {
	Id         uuid.UUID
	Status     string
	Rating     int
	Review     string
	UserId     uuid.UUID
	MovieId    uuid.UUID
	Date_added time.Time
	Update_at  time.Time
}
