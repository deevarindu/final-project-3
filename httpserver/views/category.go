package views

import (
	"time"
)

type GetCategories struct {
	ID        int        `json:"id"`
	Type      string     `json:"type"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
	Tasks     []GetTasks `json:"tasks"`
}
