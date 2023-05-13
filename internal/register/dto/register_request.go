package register

type CreateRegisterRequestBody struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required"`
}