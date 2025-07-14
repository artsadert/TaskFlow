package movie

import (
	"time"

	"github.com/google/uuid"
)

type DBMovie struct {
	Uuid        uuid.UUID `bson:"uuid"`
	Name        string    `bson:"name"`
	Year        int       `bson:"year"`
	Genre       []string  `bson:"genre"`
	Description string    `bson:"description"`
	Poster_url  string    `bson:"poster_url"`
	Update_at   time.Time `bson:"update_at"`
	Create_at   time.Time `bson:"create_at"`
}
