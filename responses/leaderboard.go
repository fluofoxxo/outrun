package responses

import (
    "github.com/fluofoxxo/outrun/responses/responseobjs"
    "github.com/jinzhu/now"
)

type WeeklyLeaderboardOptionsResponse struct {
    BaseResponse
    Mode      int64 `json:"mode"`      // 0 == ENDLESS, 1 == QUICK
    Type      int64 `json:"type"`      // 0 == RankingScoreType.HIGH_SCORE, else == RankingScoreType.TOTAL_SCORE
    Param     int64 `json:"param"`     // seemingly unused
    StartTime int64 `json:"startTime"` // both times are also seemingly unused...
    ResetTime int64 `json:"resetTime"`
}

func NewWeeklyLeaderboardOptions(base responseobjs.BaseInfo, mode, ltype, param, startTime, resetTime int64) WeeklyLeaderboardOptionsResponse {
    baseResponse := NewBaseResponse(base)
    return WeeklyLeaderboardOptionsResponse{
        baseResponse,
        mode,
        ltype,
        param,
        startTime,
        resetTime,
    }
}

func DefaultWeeklyLeaderboardOptions(base responseobjs.BaseInfo, mode int64) WeeklyLeaderboardOptionsResponse {
    startTime := now.BeginningOfDay().UTC().Unix()
    resetTime := startTime + 86400 // + 1 Day
    ltype := int64(1)
    param := int64(0)
    return NewWeeklyLeaderboardOptions(base, mode, ltype, param, startTime, resetTime)
}
