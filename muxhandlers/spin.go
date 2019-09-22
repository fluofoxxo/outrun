package muxhandlers

import (
    "encoding/json"
    "strconv"

    "github.com/fluofoxxo/outrun/config"
    "github.com/fluofoxxo/outrun/emess"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/netobj"
    "github.com/fluofoxxo/outrun/requests"
    "github.com/fluofoxxo/outrun/responses"
    "github.com/fluofoxxo/outrun/status"
)

func GetWheelOptions(helper *helper.Helper) {
    // no player, as this is (but perhaps should not be) player agnostic
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    wheelOptions := netobj.DefaultWheelOptions()
    response := responses.WheelOptions(baseInfo, wheelOptions)
    err := helper.SendResponse(response)
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
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.DefaultWheelSpin(baseInfo, player)
    err = helper.SendResponse(response)
    if err != nil {
        helper.InternalErr("Error sending response", err)
    }
}
