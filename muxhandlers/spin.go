package muxhandlers

import (
    "github.com/fluofoxxo/outrun/emess"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/netobj"
    "github.com/fluofoxxo/outrun/responses"
    "github.com/fluofoxxo/outrun/status"
)

func GetWheelOptions(helper *helper.Helper) {
    // no player, as this is (but perhaps should not be) player agnostic
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    wheelOptions := netobj.DefaultWheelOptions()
    response := responses.WheelOptions(baseInfo, wheelOptions)
    helper.SendResponse(response)
}
