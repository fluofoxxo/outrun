package handlers

import (
    "log"
    "net/http"

    "github.com/fluofoxxo/outrun/consts"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/responses"
)

func GetItemStockNumHandler(w http.ResponseWriter, r *http.Request) {
    // for right now, this is agnostic. ideally, the current event should
    // either be global or be tied to the player's choice.
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    response := responses.DefaultItemStockNumResponse(baseInfo)
    respJ, err := responses.ToJSON(response)
    if err != nil {
        log.Println("[ERR] (GetItemStockNumHandler) Error marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[OUT] (GetItemStockNumHandler) All OK")
    helper.Respond([]byte(respJ), w)
}
