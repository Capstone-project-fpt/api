package auth_dto

type InputLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type OutputLogin struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}