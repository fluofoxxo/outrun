package responses

import (
	"github.com/fluofoxxo/outrun/logic/conversion"
	"github.com/fluofoxxo/outrun/netobj"
	"github.com/fluofoxxo/outrun/obj"
	"github.com/fluofoxxo/outrun/obj/constobjs"
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

func WeeklyLeaderboardOptions(base responseobjs.BaseInfo, mode, ltype, param, startTime, resetTime int64) WeeklyLeaderboardOptionsResponse {
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
	//ltype := int64(1)
	ltype := int64(0)
	//param := int64(0)
	param := int64(5)
	return WeeklyLeaderboardOptions(base, mode, ltype, param, startTime, resetTime)
}

type WeeklyLeaderboardEntriesResponse struct {
	BaseResponse
	PlayerEntry  obj.LeaderboardEntry   `json:"playerEntry"`
	LastOffset   int64                  `json:"lastOffset"`
	StartTime    int64                  `json:"startTime"`
	ResetTime    int64                  `json:"resetTime"`
	StartIndex   int64                  `json:"startIndex"`
	Mode         int64                  `json:"mode"`
	TotalEntries int64                  `json:"totalEntries"`
	EntriesList  []obj.LeaderboardEntry `json:"entriesList"`
}

func WeeklyLeaderboardEntries(base responseobjs.BaseInfo, pe obj.LeaderboardEntry, lo, st, rt, si, m, te int64, el []obj.LeaderboardEntry) WeeklyLeaderboardEntriesResponse {
	baseResponse := NewBaseResponse(base)
	out := WeeklyLeaderboardEntriesResponse{
		baseResponse,
		pe,
		lo,
		st,
		rt,
		si,
		m,
		te,
		el,
	}
	return out
}

func DefaultWeeklyLeaderboardEntries(base responseobjs.BaseInfo, player netobj.Player, mode int64) WeeklyLeaderboardEntriesResponse {
	startTime := now.BeginningOfDay().UTC().Unix()
	resetTime := startTime + 86400 // +1 Day
	myEntry := conversion.PlayerToLeaderboardEntry(player)
	return WeeklyLeaderboardEntries(
		base,
		//obj.DefaultLeaderboardEntry(uid),
		myEntry,
		-1,
		startTime,
		resetTime,
		1,
		mode,
		0,
		[]obj.LeaderboardEntry{},
	)
}

type LeagueDataResponse struct {
	BaseResponse
	LeagueData obj.LeagueData `json:"leagueData"`
	Mode       int64          `json:"mode"`
}

func LeagueData(base responseobjs.BaseInfo, leagueData obj.LeagueData, mode int64) LeagueDataResponse {
	baseResponse := NewBaseResponse(base)
	out := LeagueDataResponse{
		baseResponse,
		leagueData,
		mode,
	}
	return out
}

func DefaultLeagueData(base responseobjs.BaseInfo, mode int64) LeagueDataResponse {
	var leagueData obj.LeagueData
	if mode == 0 {
		leagueData = constobjs.DefaultLeagueDataMode0
	} else if mode == 1 {
		leagueData = constobjs.DefaultLeagueDataMode1
	}
	return LeagueData(base, leagueData, mode)
}
