package common

import (
	"fmt"

	"github.com/go-playground/validator/v10" // go get github.com/go-playground/validator/v10
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

func (cv *CustomValidator) Validate(i any) error {

	if err := cv.validator.Struct(i); err != nil {

		// Optionally return the error to let each route control the status code.
		// return echo.ErrBadRequest.Wrap(err)

		fmt.Println(err)
		return err
	}

	return nil
}
