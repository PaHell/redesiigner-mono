package util

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationErrorsToString(err error) (str string) {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		fmt.Println(err)
		return "[(critical)]"
	}
	for _, err := range err.(validator.ValidationErrors) {
		if len(err.Param()) == 0 {
			str += fmt.Sprintf("%s:%s, ", err.Field(), err.Tag())
		} else {
			str += fmt.Sprintf("%s:%s=%s, ", err.Field(), err.Tag(), err.Param())
		}
	}
	str = strings.TrimRight(str, ", ")
	return str
}
