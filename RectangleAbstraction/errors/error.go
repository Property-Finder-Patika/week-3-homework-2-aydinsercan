package errors

import (
	"errors"
)

//check if height is >0
func CheckLength(length float32) error {
	return errors.New("can't work with less than or equal to zero Length values :) ")
}

//check if width is >0
func CheckWidth(width float32) error {
	return errors.New("can't work with less than or equal to zero Width values :) ")
}
