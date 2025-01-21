package answer

import "time"

type Answer struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	UserID     uint      `json:"user_id"`
	QuestionID uint      `json:"question_id"`
	Body       string    `gorm:"type:text" json:"body"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
