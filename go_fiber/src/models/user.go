package models

import (
	"errors"
	"time"

	"github.com/PaHell/redesiigner-mono/database"
	"github.com/PaHell/redesiigner-mono/util"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	// keys
	ID uint `gorm:"primaryKey" json:"id"`
	// props
	Email     string `json:"email" gorm:"unique" form:"email" validate:"required,email,max=255"`
	Password  string `json:"-" form:"password" validate:"required,min=10,max=255"`
	FirstName string `json:"fname" form:"fname" validate:"required,ascii,min=2,max=127"`
	LastName  string `json:"lname" form:"lname" validate:"required,ascii,min=2,max=255"`
	Username  string `json:"username" gorm:"unique" form:"username" validate:"required,ascii,min=2,max=127"`
	IsAdmin   bool   `json:"admin"`
	// timestamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// relational
	//Tokens []Token `json:"tokens" gorm:"ForeignKey:UserID"`
}

type UserModel struct{}

// Helper Functions

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (m UserModel) CheckPasswordHash(password string, user *User) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

// Controller Functions

func (m UserModel) OneByID(ID uint) (user *User, err error) {
	err = database.InstanceGorm.First(&user, "ID = ?", ID).Error
	if err != nil {
		return nil, errors.New(util.RESPONSE_MSG_CANNOT_GET)
	}
	return user, nil
}

func (m UserModel) OneByEmailUsername(emailUsername string) (user *User, err error) {
	err = database.InstanceGorm.First(&user, "Username = ?", emailUsername).Error
	if err == nil {
		return user, nil
	}
	err = database.InstanceGorm.First(&user, "Email = ?", emailUsername).Error
	if err == nil {
		return user, nil
	}
	return nil, errors.New(util.RESPONSE_MSG_CANNOT_GET)
}

func (m UserModel) All() (list []User, err error) {
	err = database.InstanceGorm.Order("id desc, username").Find(&list).Error
	if err != nil {
		return list, errors.New(util.RESPONSE_MSG_CANNOT_GET)
	}
	return list, nil
}

func (m UserModel) Create(form *User) (user *User, errors []string) {
	// validate
	v := validator.New()
	var err error
	err = v.Struct(form)
	if err != nil {
		errors = append(errors, err.Error())
	}
	// create password
	passwordHash, err := HashPassword(form.Password)
	if err != nil {
		errors = append(errors, err.Error())
	}
	// return collected errors
	if len(errors) != 0 {
		return nil, errors
	}
	// create
	newItem := &User{
		Email:     form.Email,
		Password:  passwordHash,
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Username:  form.Username,
		IsAdmin:   false,
	}
	err = database.InstanceGorm.Create(&newItem).Error
	if err != nil {
		errors = append(errors, err.Error())
		return nil, errors
	}
	return newItem, nil
}

func (m UserModel) UpdatePublic(ID uint, form *User) (item *User, errors []string) {
	// get item
	err := database.InstanceGorm.First(&item, "ID = ?", ID).Error
	if err != nil {
		errors = append(errors, err.Error())
		return nil, errors
	}
	// update props
	v := validator.New()
	if form.FirstName != "" {
		err = v.StructPartial(form, "FirstName")
		if err != nil {
			errors = append(errors, err.Error())
		}
		item.FirstName = form.FirstName
	}
	if form.LastName != "" {
		err = v.StructPartial(form, "LastName")
		if err != nil {
			errors = append(errors, err.Error())
		}
		item.LastName = form.LastName
	}
	if form.Username != "" {
		err = v.StructPartial(form, "Username")
		if err != nil {
			errors = append(errors, err.Error())
		}
		item.Username = form.Username
	}
	if len(errors) != 0 {
		return nil, errors
	}
	// save and return updated
	err = database.InstanceGorm.Save(&item).Error
	if err != nil {
		errors = append(errors, err.Error())
		return nil, errors
	}
	return item, nil
}

func (m UserModel) UpdateSecret(ID uint, form *User) (item *User, errors []string) {
	// get item
	err := database.InstanceGorm.First(&item, "ID = ?", ID).Error
	if err != nil {
		errors = append(errors, err.Error())
		return nil, errors
	}
	// update props
	v := validator.New()
	if form.Email != "" {
		err = v.StructPartial(form, "Email")
		if err != nil {
			errors = append(errors, err.Error())
		}
		item.Email = form.Email
	}
	if form.Password != "" {
		err = v.StructPartial(form, "Password")
		if err != nil {
			errors = append(errors, err.Error())
		}
		// create password
		passwordHash, err := HashPassword(form.Password)
		if err != nil {
			errors = append(errors, err.Error())
		}
		item.Password = passwordHash
	}
	if len(errors) != 0 {
		return nil, errors
	}
	// save and return updated
	err = database.InstanceGorm.Save(&item).Error
	if err != nil {
		errors = append(errors, err.Error())
		return nil, errors
	}
	return item, nil
}

func (m UserModel) Delete(ID uint) (deleted *User, err error) {
	err = database.InstanceGorm.First(&deleted, "ID = ?", ID).Error
	if err != nil {
		return nil, errors.New(util.RESPONSE_MSG_DOES_NOT_EXIST)
	}
	database.InstanceGorm.Delete(&deleted)
	return deleted, nil
}
