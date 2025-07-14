package revision

import (
	"time"

	"github.com/google/uuid"
)

type DBRevision struct {
	Uuid       uuid.UUID `bson:"uuid"`
	Status     string    `bson:"status"`
	Rating     int       `bson:"raiting"`
	Review     string    `bson:"review"`
	UserId     uuid.UUID `bson:"user_id"`
	MovieId    uuid.UUID `bson:"movie_id"`
	Date_added time.Time `bson:"create_at"`
	Update_at  time.Time `bson:"update_at"`
}
