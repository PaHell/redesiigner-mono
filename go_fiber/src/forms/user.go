package forms

type UserPasswordForm struct {
	OldPassword string `form:"old" validate:"required"`
	NewPassword string `form:"new" validate:"required"`
}

type UserEmailForm struct {
	OldEmail string `form:"old" validate:"required"`
	NewEmail string `form:"new" validate:"required"`
}
