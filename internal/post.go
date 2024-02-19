package internal

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
