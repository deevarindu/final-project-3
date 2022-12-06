package params

type CategoryCreateRequest struct {
	Type string `json:"type" binding:"required"`
}

type CategoryUpdateRequest struct {
	Type string `json:"type"`
}
