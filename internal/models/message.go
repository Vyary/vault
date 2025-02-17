package models

import (
	"regexp"
	"strings"
)

type Feedback struct {
	Name    string `json:"name,omitempty"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func (f *Feedback) Validate() Errors {
	problems := make(Errors)

	if len(strings.TrimSpace(f.Name)) > 20 {
		problems["name"] = "Name must be less than 20 characters long."
	}

	if f.Email == "" {
		problems["email"] = "Email is required and cannot be empty."
	}

	if !isValidEmail(f.Email) {
		problems["email"] = "Invalid email format"
	}

	if f.Message == "" {
		problems["message"] = "Message is required and cannot be empty."
	}

	return problems
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(strings.ToLower(email))
}
