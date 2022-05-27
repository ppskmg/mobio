package report

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Report struct {
	Email string `json:"email"`
}

func (s *Report) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Email, validation.Required),
	)
}
