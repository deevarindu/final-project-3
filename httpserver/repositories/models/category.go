package models

import "time"

type Category struct {
	ID        *int      `json:"id" gorm:"primaryKey"`
	Type      string    `json:"type" gorm:"notNull"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tasks     []Task    `gorm:"foreignkey:CategoryId"`
}
