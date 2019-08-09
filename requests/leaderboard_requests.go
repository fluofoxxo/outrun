package requests

type LeaderboardRequest struct {
	BasicRequest
	Mode int64 `json:"mode,string"`
}

type LeagueDataRequest LeaderboardRequest // are the same
