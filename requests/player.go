package requests

type SetUsernameRequest struct {
	Base
    Username string `json:"userName"`
}
