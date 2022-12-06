package views

type GetTasks struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Status      bool     `json:"status"`
	Description string   `json:"description"`
	UserID      int      `json:"user_id"`
	CategoryID  int      `json:"category_id"`
	User        GetUsers `json:"user"`
}
