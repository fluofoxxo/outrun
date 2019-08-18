package muxhandlers

import (
    "github.com/fluofoxxo/outrun/emess"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/responses"
    "github.com/fluofoxxo/outrun/status"
)

func GetDailyChallengeData(helper *helper.Helper) {
    // no player, agnostic
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.DailyChallengeData(baseInfo)
    err := helper.SendResponse(response)
    if err != nil {
        helper.InternalErr("Error sending response", err)
    }
}
