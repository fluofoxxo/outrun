package muxhandlers

import (
    "github.com/fluofoxxo/outrun/emess"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/responses"
    "github.com/fluofoxxo/outrun/status"
)

func GetFacebookIncentive(helper *helper.Helper) {
    // We respond with no presents for now.
    player, err := helper.GetCallingPlayer()
    if err != nil {
        helper.InternalErr("Error getting calling player", err)
        return
    }
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.DefaultFacebookIncentive(baseInfo, player)
    err = helper.SendResponse(response)
    if err != nil {
        helper.InternalErr("Error sending response", err)
    }
}
