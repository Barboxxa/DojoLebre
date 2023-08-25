package domain

type SignRequest struct {
	Image string `json:"image" validate:"required,base64"`
}
