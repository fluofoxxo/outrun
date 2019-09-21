package muxhandlers

import (
    "github.com/fluofoxxo/outrun/emess"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/responses"
    "github.com/fluofoxxo/outrun/status"
)

func GetItemStockNum(helper *helper.Helper) {
    // TODO: Flesh out properly! The game responds with
    // [IDRouletteTicketPremium, IDRouletteTicketItem, IDSpecialEgg]
    // for item IDs, along with an event ID, likely for event characters.
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.DefaultItemStockNum(baseInfo)
    err := helper.SendResponse(response)
    if err != nil {
        helper.InternalErr("Error sending response", err)
    }
}
