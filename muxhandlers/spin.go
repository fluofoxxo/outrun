package muxhandlers

import (
    "encoding/json"
    "strconv"

    "github.com/fluofoxxo/outrun/config"
    "github.com/fluofoxxo/outrun/db"
    "github.com/fluofoxxo/outrun/emess"
    "github.com/fluofoxxo/outrun/enums"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/netobj"
    "github.com/fluofoxxo/outrun/requests"
    "github.com/fluofoxxo/outrun/responses"
    "github.com/fluofoxxo/outrun/status"
)

func GetWheelOptions(helper *helper.Helper) {
    player, err := helper.GetCallingPlayer()
    if err != nil {
        helper.InternalErr("Error getting calling player", err)
        return
    }
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    //player.LastWheelOptions = netobj.DefaultWheelOptions(player.PlayerState) // generate new wheel for 'reroll' mechanic
    response := responses.WheelOptions(baseInfo, player.LastWheelOptions)
    err = helper.SendResponse(response)
    if err != nil {
        helper.InternalErr("Error sending response", err)
    }
}

func CommitWheelSpin(helper *helper.Helper) {
    // TODO: Should this be player agnostic? Make executive decision another time
    recv := helper.GetGameRequest()
    var request requests.CommitWheelSpinRequest
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
    if config.CFile.DebugPrints {
        helper.Out(strconv.Itoa(int(request.Count)))
    }

    wonItem := player.LastWheelOptions.Items[player.LastWheelOptions.ItemWon]
    itemExists := player.IndexOfItem(wonItem) != -1
    if itemExists {
        amountOfItemWon := player.LastWheelOptions.Item[player.LastWheelOptions.ItemWon]
        if config.CFile.DebugPrints {
            helper.Out(wonItem)
            helper.Out(strconv.Itoa(int(amountOfItemWon)))
        }
        itemIndex := player.IndexOfItem(wonItem)
        if config.CFile.DebugPrints {
            helper.Out(strconv.Itoa(int(player.PlayerState.Items[itemIndex].Amount)))
        }
        player.PlayerState.Items[itemIndex].Amount += amountOfItemWon
        if config.CFile.DebugPrints {
            helper.Out(strconv.Itoa(int(player.PlayerState.Items[itemIndex].Amount)))
        }
    } else {
        if wonItem == strconv.Itoa(enums.IDTypeItemRouletteWin) {
            // Jackpot
            player.PlayerState.NumRings += player.LastWheelOptions.NumJackpotRing
        } else if wonItem == strconv.Itoa(enums.IDTypeRedRing) {
            // Red rings
            player.PlayerState.NumRedRings += player.LastWheelOptions.Item[player.LastWheelOptions.ItemWon]
        } else {
            helper.Warn("item '" + wonItem + "' not found")
        }
    }

    // generate new wheel
    player.LastWheelOptions = netobj.DefaultWheelOptions(player.PlayerState)
    player.LastWheelOptions.NumRouletteToken = player.PlayerState.NumRouletteTicket

    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.WheelSpin(baseInfo, player.PlayerState, player.CharacterState, player.ChaoState, player.LastWheelOptions)

    err = db.SavePlayer(player)
    if err != nil {
        helper.InternalErr("Error saving player", err)
        return
    }

    err = helper.SendResponse(response)
    if err != nil {
        helper.InternalErr("Error sending response", err)
    }
}
