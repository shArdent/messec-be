package user

import (
	"time"

	"github.com/shardent/messec-be/internal/answer"
	"github.com/shardent/messec-be/internal/comment"
	"github.com/shardent/messec-be/internal/post"
	"github.com/shardent/messec-be/internal/question"
)

type User struct {
	ID        uint                `gorm:"primary_key" json:"id"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	Username  *string             `gorm:"type:varchar(100);uniqueIndex;default:null"`
	Name      *string             `gorm:"type:varchar(100);default:null"`
	Email     string              `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string              `gorm:"type:varchar(255);not null"`
	Bio       *string             `gorm:"type:text;default:null"`
	Post      []post.Post         `gorm:"foreignKey:UserID" json:"posts"`
	Comment   []comment.Comment   `gorm:"foreignKey:UserID" json:"comments"`
	Question  []question.Question `gorm:"foreignKey:UserID" json:"questions"`
	Answer    []answer.Answer     `gorm:"foreignKey:UserID" json:"answer"`
}

type UserDto struct {
	ID        uint        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Username  *string     `gorm:"type:varchar(100);uniqueIndex;default:null"`
	Name      *string     `gorm:"type:varchar(100);default:null"`
	Email     string      `gorm:"type:varchar(100);uniqueIndex;not null"`
	Bio       *string     `gorm:"type:text;default:null"`
	Post      []post.Post `gorm:"foreignKey:UserID" json:"posts"`
}

func (u *User) TableName() string {
	return "user"
}
