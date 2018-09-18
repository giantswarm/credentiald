package lister

// Response is the response format our endpoint will return.
// TODO: include the actual metadata (aws and azure specific)
type Response struct {
	ID string `json:"id"`
}
