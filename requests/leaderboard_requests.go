package requests

type LeaderboardRequest struct {
	BasicRequest
	Mode int64 `json:"mode,string"`
}

type LeagueDataRequest LeaderboardRequest // are the same

type LeaderboardEntriesRequest struct {
	BasicRequest
	Mode         int64    `json:"mode,string"`
	First        int64    `json:"first,string"`
	Count        int64    `json:"count,string"`
	Type         int64    `json:"type,string"`
	FriendIDList []string `json:"friendIdList"` // TODO: this is probably not a list of strings. Find out what it is.
}
