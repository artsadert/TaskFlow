package common

import (
	"time"

	"github.com/google/uuid"
)

type UserResult struct {
	Id        uuid.UUID
	Name      string
	Email     string
	Create_at time.Time
	Update_at time.Time
}
