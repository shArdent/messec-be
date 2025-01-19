package post

import (
	"time"
)

type Post struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Body      *string   `gorm:"type:text"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
