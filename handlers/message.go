package handlers

import (
    "log"
    "net/http"

    "github.com/fluofoxxo/outrun/consts"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/responses"
)

func GetMessageListHandler(w http.ResponseWriter, r *http.Request) {
    // this is independent of the player, so we do not need to parse the session ID
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    response := responses.NewMessageListResponse(baseInfo)
    responseJ, err := responses.ToJSON(response)
    if err != nil {
        log.Println("[ERR] (GetMessageListHandler) Error marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[OUT] (GetMessageListHandler) All OK")
    helper.Respond([]byte(responseJ), w)
}
