package models

type Auth struct {
	Code         string `json:"code"`
	CodeVerifier string `json:"codeVerifier"`
}

type PoeToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
	TokenType string `json:"bearer"`
	Scope string `json:"scope"`
	UserName string `json:"username"`
}

type PoeUser struct {
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	Locale string `json:"locale"`
}

func (a *Auth) Validate() Errors {
	problems := make(Errors)

	if a.Code == "" {
		problems["code"] = "Authorization code is required"
	}
	if a.CodeVerifier == "" {
		problems["codeVerifier"] = "code verifier is required for PKCE"
	}

	return problems
}
