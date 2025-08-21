package validator

import (
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Collect all error messages
		errors := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			var msg string
			switch err.Tag() {
			case "required":
				msg = fmt.Sprintf("%s is required", err.Field())
			case "min":
				msg = fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
			case "url":
				msg = fmt.Sprintf("%s must be a valid URL", err.Field())
			case "uuid":
				msg = fmt.Sprintf("%s must be a valid UUID", err.Field())
			case "after_2020": // eta chai date ko validator
				msg = fmt.Sprintf("%s must be on or after 2020-01-01)", err.Field())
			default:
				msg = fmt.Sprintf("%s is not valid", err.Field())
			}
			errors = append(errors, msg)
		}
		return fmt.Errorf("%v", errors)
	}
	return nil
}

func NewValidator() *CustomValidator {
	v := validator.New()
	v.RegisterValidation("after-2020", func(fl validator.FieldLevel) bool {
		dateStr := fl.Field().String()
		parsed, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Println("Error parsing date(invalid date format):", err)
			return false
		}
		//Yo chai minimum date rakhdeko 2020 jan 1
		minDate, err := time.Parse("2006-01-02", "2020-01-01")
		if err != nil {
			log.Println("Date must be on or after 2020-01-01:", err)
			return false
		}

		return !parsed.Before(minDate)
	})

	return &CustomValidator{validator: v}
}
