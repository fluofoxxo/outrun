package handlers

import (
    "log"
    "net/http"

    "github.com/fluofoxxo/outrun/consts"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/responses"
)

func GetChaoWheelOptionsHandler(w http.ResponseWriter, r *http.Request) {
    // this appears to be player agnostic.
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    response := responses.DefaultChaoWheelOptionsResponse(baseInfo)
    responseJ, err := responses.ToJSON(response)
    if err != nil {
        log.Println("[ERR] (GetChaoWheelOptionsHandler) Error marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[OUT] (GetChaoWheelOptionsHandler) All OK")
    helper.Respond([]byte(responseJ), w)
}
