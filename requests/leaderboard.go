package requests

type LeaderboardRequest struct {
    Base
    Mode int64 `json:"mode,string"`
}

type LeaderboardEntriesRequest struct {
    Base
    Mode         int64    `json:"mode,string"`
    First        int64    `json:"first,string"`
    Count        int64    `json:"count,string"`
    Type         int64    `json:"type,string"`
    FriendIDList []string `json:"friendIdList"`
}
