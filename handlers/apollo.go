package handlers

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/fluofoxxo/outrun/consts"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/responses"
)

func SendApolloHandler(w http.ResponseWriter, r *http.Request) {
    // don't log anything, just repond normally
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 1, 0)
    br := responses.NewBaseResponse(baseInfo)
    log.Println("[OUT] (SendApolloHandler) Responding with base response")
    jb, err := json.Marshal(br)
    if err != nil {
        log.Println("[ERR] (SendApolloHandler) Error marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[ERR] (SendApolloHandler) No issues")
    //w.Write(jb)
    helper.Respond(jb, w)
}
