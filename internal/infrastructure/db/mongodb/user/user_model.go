package user

import (
	"time"

	"github.com/google/uuid"
)

type DBUser struct {
	Uuid      uuid.UUID `bson:"uuid"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	Create_at time.Time `bson:"create_at"`
	Update_at time.Time `bson:"update_at"`
}
