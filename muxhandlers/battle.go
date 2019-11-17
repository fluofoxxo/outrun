package muxhandlers

import (
	"encoding/json"
	"time"

	"github.com/fluofoxxo/outrun/emess"
	"github.com/fluofoxxo/outrun/helper"
	"github.com/fluofoxxo/outrun/obj"
	"github.com/fluofoxxo/outrun/requests"
	"github.com/fluofoxxo/outrun/responses"
	"github.com/fluofoxxo/outrun/status"
)

func UpdateDailyBattleStatus(helper *helper.Helper) {
	data := helper.GetGameRequest()
	var request requests.Base
	err := json.Unmarshal(data, &request)
	if err != nil {
		helper.InternalErr("Error unmarshalling", err)
		return
	}
	endTime := time.Now().UTC().Unix() + 180 // three minutes from now, for testing
	rewardFlag := false
	battleStatus := obj.DefaultBattleStatus()
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.UpdateDailyBattleStatus(baseInfo, endTime, battleStatus, rewardFlag)
	err = helper.SendInsecureResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
}
