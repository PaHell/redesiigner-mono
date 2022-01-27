package models

import (
	"errors"
	"time"

	"github.com/PaHell/redesiigner-mono/database"
	"github.com/PaHell/redesiigner-mono/util"

	"github.com/go-playground/validator/v10"
)

type Application struct {
	// keys
	ID     uint `gorm:"primaryKey" json:"id"`
	UserID uint `gorm:"index" json:"user_id"`
	// props
	Title       string `json:"title" gorm:"unique" form:"title" validate:"required,ascii,min=2,max=128"`
	Description string `json:"description" form:"description" validate:"required,min=12,max=512"`
	Url         string `json:"url" gorm:"unique" form:"url" validate:"required,url"`
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
	err = database.InstanceGorm.Order("title asc, url asc").Find(&list).Error
	if err != nil {
		return list, errors.New(util.RESPONSE_MSG_CANNOT_GET)
	}
	return list, nil
}

func (m ApplicationModel) Create(form *Application) (application *Application, errors []string) {
	// validate
	v := validator.New()
	var err error
	err = v.StructPartial(form, "Title", "Url", "Description")
	if err != nil {
		errors = append(errors, err.Error())
	}
	// check unique
	err = database.InstanceGorm.First(&Application{}, "Url = ?", form.Url).Error
	if err == nil {
		errors = append(errors, "Url needs to be unique")
	}
	err = database.InstanceGorm.First(&Application{}, "Title = ?", form.Title).Error
	if err == nil {
		errors = append(errors, "Title needs to be unique")
	}
	// return collected errors
	if len(errors) != 0 {
		return nil, errors
	}
	// create
	newItem := &Application{
		Title:       form.Title,
		Description: form.Description,
		Url:         form.Url,
	}
	err = database.InstanceGorm.Create(&newItem).Error
	if err != nil {
		errors = append(errors, err.Error())
		return nil, errors
	}
	return newItem, nil
}

func (m ApplicationModel) Update(ID uint, form *Application) (item *Application, errors []string) {
	// get item
	err := database.InstanceGorm.First(&item, "ID = ?", ID).Error
	if err != nil {
		errors = append(errors, err.Error())
		return nil, errors
	}
	// update props
	v := validator.New()
	if form.Title != "" {
		err = v.StructPartial(form, "Title")
		if err != nil {
			errors = append(errors, err.Error())
		}
		item.Title = form.Title
	}
	if form.Description != "" {
		err = v.StructPartial(form, "Description")
		if err != nil {
			errors = append(errors, err.Error())
		}
		item.Description = form.Description
	}
	if form.Url != "" {
		err = v.StructPartial(form, "Url")
		if err != nil {
			errors = append(errors, err.Error())
		}
		item.Url = form.Url
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

func (m ApplicationModel) Delete(ID uint) (deleted *Application, err error) {
	err = database.InstanceGorm.First(&deleted, "ID = ?", ID).Error
	if err != nil {
		return nil, errors.New(util.RESPONSE_MSG_DOES_NOT_EXIST)
	}
	database.InstanceGorm.Delete(&deleted)
	return deleted, nil
}
