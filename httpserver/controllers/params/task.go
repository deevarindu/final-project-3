package params

type TaskCreateRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID  int    `json:"category_id" binding:"required"`
}

type TaskUpdateTitleDescRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type TaskUpdateStatusRequest struct {
	Status bool `json:"status" binding:"required"`
}

type TaskUpdateCategoryRequest struct {
	CategoryID int `json:"category_id" binding:"required"`
}
