package responses

import "github.com/jinzhu/now"

type LeaderboardResponse struct {
    BaseResponse
    Type      int64 `json:"type"`
    Param     int64 `json:"param"`
    ResetTime int64 `json:"resetTime"`
    StartTime int64 `json:"startTime"`
    Mode      int64 `json:"mode"`
}

func NewLeaderboardResponse(base BaseInfo, lType, param, resetTime, startTime, mode int64) LeaderboardResponse {
    br := NewBaseResponse(base)
    lbr := LeaderboardResponse{
        br,
        lType,
        param,
        resetTime,
        startTime,
        mode,
    }
    return lbr
}

func DefaultLeaderboardResponse(base BaseInfo, mode int64) LeaderboardResponse {
    startTime := now.BeginningOfDay().UTC().Unix()
    resetTime := startTime + 86400 // + 1 Day
    param := int64(5)
    lType := int64(0)
    lbr := NewLeaderboardResponse(
        base,
        lType,
        param,
        resetTime,
        startTime,
        mode,
    )
    return lbr
}
