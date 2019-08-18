package requests

type LeaderboardRequest struct {
    Base
    Mode int64 `json:"mode,string"`
}
