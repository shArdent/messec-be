package post

import (
	"time"

	"github.com/shardent/messec-be/internal/comment"
)

type Post struct {
	ID        uint              `gorm:"primary_key" json:"id"`
	Body      string            `gorm:"type:text" json:"body"`
	UserID    uint              `json:"user_id"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
    Comment   []comment.Comment `gorm:"foreignKey:PostID,constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"comments"`
}
