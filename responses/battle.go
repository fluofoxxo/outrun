package responses

import (
    "time"

    "github.com/fluofoxxo/outrun/obj"
    "github.com/fluofoxxo/outrun/responses/responseobjs"
)

type DailyBattleDataResponse struct {
    BaseResponse
    EndTime      int64               `json:"endTime"`
    BattleStatus obj.DailyBattleData `json:"battleData"`
}

func DailyBattleData(base responseobjs.BaseInfo, endTime int64, battleData obj.DailyBattleData) DailyBattleDataResponse {
    baseResponse := NewBaseResponse(base)
    return DailyBattleDataResponse{
        baseResponse,
        endTime,
        battleData,
    }
}

func DefaultDailyBattleData(base responseobjs.BaseInfo) DailyBattleDataResponse {
    battleData := obj.DefaultDailyBattleData()
    return DailyBattleData(
        base,
        time.Now().Unix()+80000, // ~22 hours from now
        battleData,
    )
}

type DailyBattleStatusResponse struct {
    BaseResponse
    EndTime                 int64                 `json:"endTime"`
    BattleStatus            obj.DailyBattleStatus `json:"battleStatus"`
    RewardFlag              bool                  `json:"rewardFlag"`
    obj.DailyBattleDataPair                       // embedded in root. Can be missing if rewardFlag = 0
}

func DailyBattleStatus(base responseobjs.BaseInfo, endTime int64, battleStatus obj.DailyBattleStatus, rewardFlag bool, battlePair obj.DailyBattleDataPair) DailyBattleStatusResponse {
    baseResponse := NewBaseResponse(base)
    return DailyBattleStatusResponse{
        baseResponse,
        endTime,
        battleStatus,
        rewardFlag,
        battlePair,
    }
}

func DefaultDailyBattleStatus(base responseobjs.BaseInfo) DailyBattleStatusResponse {
    endTime := time.Now().Unix() + 700
    battleStatus := obj.DefaultDailyBattleStatus()
    rewardFlag := false
    battlePair := obj.DefaultDailyBattleDataPair()
    return DailyBattleStatus(
        base,
        endTime,
        battleStatus,
        rewardFlag,
        battlePair,
    )
}
