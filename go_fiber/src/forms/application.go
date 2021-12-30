package forms

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type ApplicationForm struct{}

type CreateApplicationForm struct {
	Title       string
	Description string
	Url         string
}

func (f ApplicationForm) Title(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the article title"
		}
		return errMsg[0]
	case "min", "max":
		return "Title should be between 3 to 100 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

//Content ...
func (f ApplicationForm) Content(tag string, errMsg ...string) (message string) {
	switch tag {
	case "required":
		if len(errMsg) == 0 {
			return "Please enter the article content"
		}
		return errMsg[0]
	case "min", "max":
		return "Content should be between 3 to 1000 characters"
	default:
		return "Something went wrong, please try again later"
	}
}

//Create ...
func (f ApplicationForm) Create(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Title" {
				return f.Title(err.Tag())
			}
			if err.Field() == "Content" {
				return f.Content(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}

//Update ...
func (f ApplicationForm) Update(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:

		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return "Something went wrong, please try again later"
		}

		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Title" {
				return f.Title(err.Tag())
			}
			if err.Field() == "Content" {
				return f.Content(err.Tag())
			}
		}

	default:
		return "Invalid request"
	}

	return "Something went wrong, please try again later"
}
