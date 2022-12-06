package models

import "time"

type Task struct {
	ID          *int      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"notNull"`
	Description string    `json:"description" gorm:"notNull"`
	Status      bool      `json:"status"`
	UserID      int       `json:"user_id"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	User        User      `gorm:"foreignKey:UserID"`
}
