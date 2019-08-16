package muxhandlers

import (
    "github.com/fluofoxxo/outrun/emess"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/responses"
    "github.com/fluofoxxo/outrun/status"
)

func SendApollo(helper *helper.Helper) {
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.NewBaseResponse(baseInfo)
    err := helper.SendResponse(response)
    if err != nil {
        helper.InternalErr("Error sending response", err)
    }
}
