package muxhandlers

import (
    "encoding/json"
    "strconv"

    "github.com/fluofoxxo/outrun/emess"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/obj"
    "github.com/fluofoxxo/outrun/obj/constobjs"
    "github.com/fluofoxxo/outrun/requests"
    "github.com/fluofoxxo/outrun/responses"
    "github.com/fluofoxxo/outrun/status"
)

func GetRedStarExchangeList(helper *helper.Helper) {
    recv := helper.GetGameRequest()
    var request requests.RedStarExchangeListRequest
    err := json.Unmarshal(recv, &request)
    if err != nil {
        helper.Err("Error unmarshalling", err)
        return
    }

    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    var response responses.RedStarExchangeListResponse
    var redStarItems []obj.RedStarItem
    helper.Out("Recv ItemType " + strconv.Itoa(int(request.ItemType)))
    if request.ItemType == 0 {
        redStarItems = []obj.RedStarItem{}
    } else if request.ItemType == 1 {
        redStarItems = constobjs.RedStarItemsType1
    } else if request.ItemType == 2 {
        redStarItems = constobjs.RedStarItemsType2
    } else if request.ItemType == 4 {
        redStarItems = constobjs.RedStarItemsType4
    } else {
        helper.Respond([]byte("Invalid request"))
        return
    }

    response = responses.RedStarExchangeList(baseInfo, redStarItems, 0, "1900-1-1")
    err = helper.SendResponse(response)
    if err != nil {
        helper.InternalErr("Error sending response", err)
    }
}
