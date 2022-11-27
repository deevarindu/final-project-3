package models

import "time"

type User struct {
	ID        *int      `json:"id" gorm:"primaryKey"`
	FullName  string    `json:"full_name" gorm:"notNull"`
	Email     string    `json:"email" gorm:"uniqueIndex;notNull"`
	Password  string    `json:"password" gorm:"notNull" valid:"minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Role      string    `json:"role" gorm:"notNull"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
