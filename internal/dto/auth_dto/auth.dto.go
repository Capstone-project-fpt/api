package auth_dto

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

type ResetPasswordInput struct {
	Password string `json:"password" binding:"required"`
	Token    string `json:"token" binding:"required"`
}

type LoginOutput struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// This used for swagger
type LoginSwaggerOutput struct {
	Code    int         `json:"code"`
	Success bool        `json:"message"`
	Data    LoginOutput `json:"data"`
}
