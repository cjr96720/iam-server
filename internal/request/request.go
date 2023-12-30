package request

type CreateApplicationRequest struct {
	Name string `json:"name" binding:"required,min=3"`
}
