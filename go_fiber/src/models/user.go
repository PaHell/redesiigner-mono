package models

import (
	"errors"
	"time"

	"github.com/PaHell/redesiigner-mono/database"
	"github.com/PaHell/redesiigner-mono/forms"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	// keys
	ID uint `gorm:"primaryKey" json:"id"`
	// props
	Email    string
	Password []byte
	FName    string
	LName    string
	Username string
	IsAdmin  bool
	// timestamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// relational
}

type UserModel struct{}

var authModel = new(AuthModel)

//Login ...
func (m UserModel) Login(form forms.LoginForm) (user *User, token *Token, err error) {
	// get username match
	database.InstanceGorm.First(&user, "Username = ?", form.EmailUsername)
	// otherwise get email match
	if user == nil {
		database.InstanceGorm.First(&user, "Email = ?", form.EmailUsername)
	}
	if user == nil {
		return nil, nil, errors.New("User does not exist")
	}

	//Compare the password form and database if match
	bytePassword := []byte(form.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err != nil {
		return user, token, err
	}

	//Generate the JWT auth token
	tokenDetails, err := authModel.CreateToken(user.ID)
	if err != nil {
		return user, token, err
	}

	saveErr := authModel.CreateAuth(user.ID, tokenDetails)
	if saveErr == nil {
		token.AccessToken = tokenDetails.AccessToken
		token.RefreshToken = tokenDetails.RefreshToken
	}

	return user, token, nil
}

//Register ...
func (m UserModel) Register(form forms.RegisterForm) (user *User, err error) {
	// check if user already exists
	database.InstanceGorm.First(&user, "Username = ?", form.Username)
	if user != nil {
		return nil, errors.New("Username already taken")
	}
	database.InstanceGorm.First(&user, "Email = ?", form.Email)
	if user != nil {
		return nil, errors.New("Email already taken")
	}
	// create password hash
	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return user, errors.New("something went wrong, please try again later")
	}
	// create user
	database.InstanceGorm.Create(&User{
		Email:    form.Email,
		Password: hashedPassword,
		FName:    form.FirstName,
		LName:    form.LastName,
		Username: form.Username,
		IsAdmin:  false,
	})
	return user, nil
}

//One ...
func (m UserModel) One(userID int64) (user *User, err error) {
	database.InstanceGorm.First(&user, userID)
	if user == nil {
		return nil, errors.New("User does not exist")
	}
	return user, nil
}
