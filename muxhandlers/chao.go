package muxhandlers

import (
	"encoding/json"

	"github.com/fluofoxxo/outrun/config"
	"github.com/fluofoxxo/outrun/db"
	"github.com/fluofoxxo/outrun/emess"
	"github.com/fluofoxxo/outrun/helper"
	"github.com/fluofoxxo/outrun/requests"
	"github.com/fluofoxxo/outrun/responses"
	"github.com/fluofoxxo/outrun/status"
)

func GetChaoWheelOptions(helper *helper.Helper) {
	// no player, agnostic
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultChaoWheelOptions(baseInfo)
	err := helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetPrizeChaoWheelSpin(helper *helper.Helper) {
	// agnostic
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultPrizeChaoWheel(baseInfo)
	err := helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func EquipChao(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.EquipChaoRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}

	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}

	mainChaoID := request.MainChaoID
	subChaoID := request.SubChaoID
	if mainChaoID != "-1" {
		player.PlayerState.MainChaoID = mainChaoID
	}
	if subChaoID != "-1" {
		player.PlayerState.SubChaoID = subChaoID
	}
	if config.CFile.DebugPrints {
		helper.Out("Main Chao: " + mainChaoID)
		helper.Out("Sub Chao: " + subChaoID)
	}
	db.SavePlayer(player)

	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.EquipChao(baseInfo, player.PlayerState)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func CommitChaoWheelSpin(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultChaoWheelSpin(baseInfo, player)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}
