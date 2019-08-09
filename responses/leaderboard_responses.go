package responses

import (
    "encoding/json"

    "github.com/fluofoxxo/outrun/consts"
    "github.com/fluofoxxo/outrun/objects"
    "github.com/jinzhu/now"
)

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

type LeagueDataResponse struct {
    BaseResponse
    LeagueData objects.LeagueData `json:"leagueData"`
    Mode       int64              `json:"mode"`
}

func NewLeagueDataResponse(base BaseInfo, leagueData objects.LeagueData, mode int64) LeagueDataResponse {
    br := NewBaseResponse(base)
    ldr := LeagueDataResponse{
        br,
        leagueData,
        mode,
    }
    return ldr
}

func DefaultLeagueDataResponse(base BaseInfo, mode int64) LeagueDataResponse {
    var leagueData objects.LeagueData
    if mode == 0 {
        err := json.Unmarshal([]byte(consts.JSON_DEFAULT_LEAGUEDATA_MODE0), &leagueData)
        if err != nil {
            panic(err)
        }
    } else if mode == 1 {
        err := json.Unmarshal([]byte(consts.JSON_DEFAULT_LEAGUEDATA_MODE1), &leagueData)
        if err != nil {
            panic(err)
        }
    }
    return NewLeagueDataResponse(base, leagueData, mode)
}
