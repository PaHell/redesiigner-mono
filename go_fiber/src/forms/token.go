package forms

type TokenCreateForm struct {
	EmailUsername string `form:"email_username" validate:"required"`
	Password      string `form:"password" validate:"required"`
}

type TokenUpdateForm struct {
	ID           uint   `form:"id" validate:"required"`
	RefreshToken string `form:"refresh_token" validate:"required"`
}
