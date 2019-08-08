package requests

type LeaderboardRequest struct {
	BasicRequest
	Mode int64 `json:"mode,string"`
}

type LeagueDataRequest struct {
	LeaderboardRequest                    // Also includes a Mode
	LeagueData         objects.LeagueData `json:"leagueData"`
}
