package muxhandlers

import (
    "github.com/fluofoxxo/outrun/emess"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/responses"
    "github.com/fluofoxxo/outrun/status"
)

func GetPlayerState(helper *helper.Helper) {
    player, err := helper.GetCallingPlayer()
    if err != nil {
        helper.InternalErr("Error getting calling player", err)
        return
    }
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.PlayerState(baseInfo, player.PlayerState)
    helper.SendResponse(response)
}

func GetCharacterState(helper *helper.Helper) {
    player, err := helper.GetCallingPlayer()
    if err != nil {
        helper.InternalErr("Error getting calling player", err)
        return
    }
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.CharacterState(baseInfo, player.CharacterState)
    helper.SendResponse(response)
}

func GetChaoState(helper *helper.Helper) {
    player, err := helper.GetCallingPlayer()
    if err != nil {
        helper.InternalErr("Error getting calling player", err)
        return
    }
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.ChaoState(baseInfo, player.ChaoState)
    helper.SendResponse(response)
}
