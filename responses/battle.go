package responses

import (
    "time"

    "github.com/fluofoxxo/outrun/obj"
    "github.com/fluofoxxo/outrun/responses/responseobjs"
)

type DailyBattleDataResponse struct {
    BaseResponse
    EndTime      int64               `json:"endTime"`
    BattleStatus obj.DailyBattleData `json:"battleStatus"`
}

func DailyBattleData(base responseobjs.BaseInfo, endTime int64, battleStatus obj.DailyBattleData) DailyBattleDataResponse {
    baseResponse := NewBaseResponse(base)
    return DailyBattleDataResponse{
        baseResponse,
        endTime,
        battleStatus,
    }
}

func DefaultDailyBattleData(base responseobjs.BaseInfo) DailyBattleDataResponse {
    battleStatus := obj.DefaultDailyBattleData()
    return DailyBattleData(
        base,
        time.Now().Unix()+80000, // ~22 hours from now
        battleStatus,
    )
}

type DailyBattleStatusResponse struct {
    DailyBattleDataResponse
    RewardFlag              bool `json:"rewardFlag"`
    obj.DailyBattleDataPair      // embedded in root
}

func DailyBattleStatus(dbdr DailyBattleDataResponse, rewardFlag bool, battlePair obj.DailyBattleDataPair) DailyBattleStatusResponse {
    return DailyBattleStatusResponse{
        dbdr,
        rewardFlag,
        battlePair,
    }
}

func DefaultDailyBattleStatus(base responseobjs.BaseInfo) DailyBattleStatusResponse {
    dbdr := DefaultDailyBattleData(base)
    rewardFlag := false
    battlePair := obj.DefaultDailyBattleDataPair()
    return DailyBattleStatus(
        dbdr,
        rewardFlag,
        battlePair,
    )
}
