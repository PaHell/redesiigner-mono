package util

import (
	"errors"
	"strconv"
)

func ParseID(str string) (parsed uint, err error) {
	// parse id
	id, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, errors.New("Given ID is not a valid integer")
	}
	return uint(id), nil
}
