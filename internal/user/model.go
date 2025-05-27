package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`

	WatchedAnimeIDs  pq.StringArray `gorm:"type:text[]" json:"watched_anime_ids"`
	FavoriteAnimeIDs pq.StringArray `gorm:"type:text[]" json:"favorite_anime_ids"`
}
