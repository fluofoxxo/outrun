package handlers

import (
    "log"
    "net/http"

    "github.com/fluofoxxo/outrun/consts"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/objects"
    "github.com/fluofoxxo/outrun/responses"
)

func GetDailyChalDataHandler(w http.ResponseWriter, r *http.Request) {
    // this is independent of the player, so we do not need to parse the session ID
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    response := responses.NewDailyChalDataResponse(baseInfo)
    responseJ, err := responses.ToJSON(response)
    if err != nil {
        log.Println("[ERR] (GetDailyChalDataHandler) Error marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[OUT] (GetDailyChalDataHandler) All OK")
    helper.Respond([]byte(responseJ), w)
}

func GetCostListHandler(w http.ResponseWriter, r *http.Request) {
    // independent of player
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    response := responses.NewConsumedCostListResponse(baseInfo, objects.DEFAULT_CONSUMEDCOSTLIST)
    responseJ, err := responses.ToJSON(response)
    if err != nil {
        log.Println("[ERR] (GetCostListHandler) Error marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[OUT] (GetCostListHandler) All OK")
    helper.Respond([]byte(responseJ), w)
}
