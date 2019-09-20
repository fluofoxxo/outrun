package muxhandlers

import (
    "github.com/fluofoxxo/outrun/emess"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/responses"
    "github.com/fluofoxxo/outrun/status"
)

func GetDailyBattleData(helper *helper.Helper) {
    // TODO: Right now, send agnostic data. In reality, this definitely should be player based!
    /*
       player, err := helper.GetCallingPlayer()
       if err != nil {
           helper.InternalErr("error getting calling player", err)
           return
       }
    */
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.DefaultDailyBattleData(baseInfo)
    err := helper.SendInsecureResponse(response)
    if err != nil {
        helper.InternalErr("error sending response", err)
    }
}

func UpdateDailyBattleStatus(helper *helper.Helper) {
    // TODO: what is this function meant to do? Why is it referred to as 'update?'
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.DefaultDailyBattleStatus(baseInfo)
    err := helper.SendInsecureResponse(response)
    if err != nil {
        helper.InternalErr("error sending reponse", err)
    }
}
