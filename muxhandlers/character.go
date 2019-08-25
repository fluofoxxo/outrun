package muxhandlers

import (
	"encoding/json"

	"github.com/fluofoxxo/outrun/db"
	"github.com/fluofoxxo/outrun/emess"
	"github.com/fluofoxxo/outrun/helper"
	"github.com/fluofoxxo/outrun/requests"
	"github.com/fluofoxxo/outrun/responses"
	"github.com/fluofoxxo/outrun/status"
)

func ChangeCharacter(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.ChangeCharacterRequest
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

	mainCharaID := request.MainCharaID
	subCharaID := request.SubCharaID
	if mainCharaID != "-1" {
		player.PlayerState.MainCharaID = mainCharaID
	}
	if subCharaID != "-1" {
		player.PlayerState.SubCharaID = subCharaID
	}
	db.SavePlayer(player)

	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.ChangeCharacter(baseInfo, player.PlayerState)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}
