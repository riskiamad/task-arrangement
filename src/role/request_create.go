package role

type createRequest struct {
	Name string `json:"name" binding:"required"`
}
