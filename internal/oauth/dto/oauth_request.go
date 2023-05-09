package oauth

type LoginRequestBody struct {
	Email        string `json:"email" binding:"email"`
	Password     string `json:"password" binding:"required"`
	ClientId     string `json:"client_id" binding:"required"`
	ClientSecret string `json:"client_secret" binding:"required"`
}

type RefreshTokenRequestBody struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
