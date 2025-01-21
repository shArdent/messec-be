package question

import (
	"time"

	"github.com/shardent/messec-be/internal/answer"
)

type Question struct {
	ID        uint          `gorm:"primaryKey" json:"id"`
	UserID    uint          `json:"user_id"`
	Body      string        `json:"body"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Answer    answer.Answer `gorm:"foreignKey:PostID" json:"comments"`
}
