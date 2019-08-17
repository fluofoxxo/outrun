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

func GetPlayerState(helper *helper.Helper) {
    recv := helper.GetGameRequest()
    var request requests.Base
    err := json.Unmarshal(recv, &request)
    if err != nil {
        helper.Err("Error unmarshalling", err)
        return
    }
    sid := request.SessionID
    player, err := db.GetPlayerBySessionID(sid)
    if err != nil {
        helper.InternalErr("Error getting player by session ID", err)
        return
    }
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.PlayerState(baseInfo, player.PlayerState)
    helper.SendResponse(response)
}

func GetCharacterState(helper *helper.Helper) {
    recv := helper.GetGameRequest()
    var request requests.Base
    err := json.Unmarshal(recv, &request)
    if err != nil {
        helper.Err("Error unmarshalling", err)
        return
    }
    sid := request.SessionID
    player, err := db.GetPlayerBySessionID(sid)
    if err != nil {
        helper.InternalErr("Error getting player by session ID", err)
        return
    }
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.CharacterState(baseInfo, player.CharacterState)
    helper.SendResponse(response)
}

func GetChaoState(helper *helper.Helper) {
    recv := helper.GetGameRequest()
    var request requests.Base
    err := json.Unmarshal(recv, &request)
    if err != nil {
        helper.Err("Error unmarshalling", err)
        return
    }
    sid := request.SessionID
    player, err := db.GetPlayerBySessionID(sid)
    if err != nil {
        helper.InternalErr("Error getting player by session ID", err)
        return
    }
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.ChaoState(baseInfo, player.ChaoState)
    helper.SendResponse(response)
}
