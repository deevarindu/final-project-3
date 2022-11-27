package models

type Task struct {
	ID          *int   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"notNull"`
	Description string `json:"description" gorm:"notNull"`
	Status      bool   `json:"status"`
	UserID      int    `json:"user_id" gorm:"foreignKey:UserID"`
	CategoryID  int    `json:"category_id" gorm:"foreignKey:CategoryID"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
