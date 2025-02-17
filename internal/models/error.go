package models

type Errors map[string]string

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Errors  Errors `json:"errors,omitempty"`
}
