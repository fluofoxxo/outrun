package muxhandlers

import (
	"encoding/json"
	"strconv"

	"github.com/fluofoxxo/outrun/consts"
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

func UpgradeCharacter(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.UpgradeCharacterRequest
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

	charaID := request.CharacterID
	abilityID := request.AbilityID
	abilityIDFromStr, err := strconv.Atoi(abilityID)
	if err != nil {
		helper.Err("Error in strconv.Atoi", err)
		return
	}

	sum := func(in []int64) int64 {
		v := int64(0)
		for _, val := range in {
			v += val
		}
		return v
	}

	sendStatus := status.OK
	abilityIndex := abilityIDFromStr - 120000 // minus enums.UpgradeAbilityInvincibility
	index := player.IndexOfChara(charaID)
	abilitySum := sum(player.CharacterState[index].AbilityLevel)
	amountNeedToBePaid := player.CharacterState[index].Cost - player.CharacterState[index].Exp
	if player.PlayerState.NumRings-amountNeedToBePaid < 0 {
		sendStatus = status.NotEnoughRings
	} else {
		if abilitySum < 100 {
			player.CharacterState[index].AbilityLevel[abilityIndex]++
			player.CharacterState[index].Exp -= consts.CharacterUpgradeIncrease // TODO: subtracting exp actually increases cost... Probably not what we need to do.
			player.PlayerState.NumRings -= amountNeedToBePaid
			db.SavePlayer(player)
		} else {
			sendStatus = status.CharacterLevelLimit
		}
	}

	baseInfo := helper.BaseInfo(emess.OK, int64(sendStatus))
	response := responses.DefaultUpgradeCharacter(baseInfo, player)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}
