package role

type updateRequest struct {
	ID   int64  `json:"-" binding:"required"`
	Name string `json:"name" binding:"required"`
}
