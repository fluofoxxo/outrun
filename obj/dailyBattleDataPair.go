package obj

import "github.com/jinzhu/now"

type DailyBattleDataPair struct {
    StartTime       int64           `json:"rewardStartTime"`
    EndTime         int64           `json:"rewardEndTime"`
    BattleData      DailyBattleData `json:"rewardBattleData"`
    RivalBattleData DailyBattleData `json:"rewardRivalBattleData"`
}

func DefaultDailyBattleDataPair() DailyBattleDataPair {
    startTime := now.BeginningOfDay().Unix()
    endTime := now.EndOfDay().Unix()
    battleData := DefaultDailyBattleData()
    rivalBattleData := DefaultDailyBattleData()
    return DailyBattleDataPair{
        startTime,
        endTime,
        battleData,
        rivalBattleData,
    }
}
