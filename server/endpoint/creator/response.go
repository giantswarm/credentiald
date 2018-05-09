package creator

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`

	CredentialID string `json:"-"`
	Organization string `json:"-"`
}
