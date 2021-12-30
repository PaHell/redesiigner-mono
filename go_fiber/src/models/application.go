package models

import (
	"errors"
	"strings"
	"time"

	"github.com/PaHell/redesiigner-mono/database"
	"github.com/PaHell/redesiigner-mono/util"

	"github.com/go-playground/validator/v10"
)

var errApp = map[string]string{
	"GET_ONE":      "App with that ID does not exist",
	"GET_ALL":      "Apps could not be resolved",
	"URL_UNIQUE":   "App with same Url already exists",
	"TITLE_UNIQUE": "App with same Title already exists",
	"CREATE":       "App could not be created",
	"UPDATE":       "App could not be updated",
	"DELETE":       "App could not be deleted",
}

type Application struct {
	// keys
	ID     uint `gorm:"primaryKey" json:"id"`
	UserID uint `gorm:"index" json:"user_id"`
	// props
	Title       string `json:"title" gorm:"unique" form:"title" validate:"required,omitempty,ascii,min=2,max=128"`
	Description string `json:"description" form:"description" validate:"required,omitempty,min=32,max=512"`
	Url         string `json:"url" gorm:"unique" form:"url" validate:"required,omitempty,url"`
	// timestamps
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// relational
	User User `json:"user" gorm:"embedded"`
}

type ApplicationModel struct{}

func (m ApplicationModel) One(ID uint) (application *Application, err error) {
	err = database.InstanceGorm.First(&application, "ID = ?", ID).Error
	if err != nil {
		return nil, errors.New(util.RESPONSE_MSG_CANNOT_GET)
	}
	return application, nil
}

func (m ApplicationModel) All() (list []Application, err error) {
	err = database.InstanceGorm.Order("id desc, user_id").Find(&list).Error
	if err != nil {
		return list, errors.New(util.RESPONSE_MSG_CANNOT_GET)
	}
	return list, nil
}

func (m ApplicationModel) Create(form *Application) (application *Application, err error) {
	// validate
	v := validator.New()
	err = v.Struct(form)
	if err != nil {
		return nil, errors.New(util.ValidationErrorsToString(err))
	}
	// check unique
	err = database.InstanceGorm.First(&Application{}, "Url = ?", form.Url).Error
	if err == nil {
		return nil, errors.New("Property url needs to be unique")
	}
	err = database.InstanceGorm.First(&Application{}, "Title = ?", form.Title).Error
	if err == nil {
		return nil, errors.New("Property title needs to be unique")
	}
	// create
	newItem := &Application{
		Title:       form.Title,
		Description: form.Description,
		Url:         form.Url,
	}
	err = database.InstanceGorm.Create(&newItem).Error
	if err == nil {
		return nil, errors.New(util.RESPONSE_MSG_CANNOT_CREATE)
	}
	return newItem, nil
}

func (m ApplicationModel) Update(ID uint, form *Application) (item *Application, err error) {
	// get item
	err = database.InstanceGorm.First(&item, "ID = ?", ID).Error
	if err != nil {
		return nil, errors.New(util.RESPONSE_MSG_DOES_NOT_EXIST)
	}
	// update props
	v := validator.New()
	var propErrors []string
	if len(form.Title) != 0 {
		err = v.StructPartial(form, "Title")
		if err != nil {
			propErrors = append(propErrors, util.ValidationErrorsToString(err))
		}
		item.Title = form.Title
	}
	if len(form.Description) != 0 {
		err = v.StructPartial(form, "Description")
		if err != nil {
			propErrors = append(propErrors, util.ValidationErrorsToString(err))
		}
		item.Description = form.Description
	}
	if len(form.Url) != 0 {
		err = v.StructPartial(form, "Url")
		if err != nil {
			propErrors = append(propErrors, util.ValidationErrorsToString(err))
		}
		item.Url = form.Url
	}
	if err != nil {
		return nil, errors.New(strings.Join(propErrors, ", "))
	}
	// save and return updated
	err = database.InstanceGorm.Save(&item).Error
	if err != nil {
		return nil, errors.New(util.RESPONSE_MSG_CANNOT_UPDATE)
	}
	return item, nil
}

//Delete ...
func (m ApplicationModel) Delete(ID uint) (deleted *Application, err error) {
	err = database.InstanceGorm.First(&deleted, "ID = ?", ID).Error
	if err != nil {
		return nil, errors.New(util.RESPONSE_MSG_DOES_NOT_EXIST)
	}
	database.InstanceGorm.Delete(&deleted)
	return deleted, nil
}
