package comment

import "time"

type Comment struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    *uint     `json:"user_id"`
	PostID    uint      `json:"post_id"`
	Body      *string   `gorm:"type:text" json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
