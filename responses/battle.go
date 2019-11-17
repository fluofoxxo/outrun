package responses

import (
	"github.com/fluofoxxo/outrun/obj"
	"github.com/fluofoxxo/outrun/responses/responseobjs"
)

type UpdateDailyBattleStatusResponse struct {
	BaseResponse
	EndTime      int64            `json:"endTime"`
	BattleStatus obj.BattleStatus `json:"battleStatus"`
	RewardFlag   bool             `json:"rewardFlag"` // TODO: allow not false after testing
}

func UpdateDailyBattleStatus(base responseobjs.BaseInfo, endTime int64, battleStatus obj.BattleStatus, rewardFlag bool) UpdateDailyBattleStatusResponse {
	baseResponse := NewBaseResponse(base)
	return UpdateDailyBattleStatusResponse{
		baseResponse,
		endTime,
		battleStatus,
		rewardFlag,
	}
}
