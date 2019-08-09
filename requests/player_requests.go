package requests

type SetUsernameRequest struct {
	BasicRequest
	Username string `json:"userName"`
}
