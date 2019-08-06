package handlers

import (
    "log"
    "net/http"

    "github.com/fluofoxxo/outrun/consts"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/objects"
    "github.com/fluofoxxo/outrun/responses"
)

func GetEventListHandler(w http.ResponseWriter, r *http.Request) {
    // for right now, this is agnostic. ideally, the current event should
    // either be global or be tied to the player's choice.
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    events := []objects.Event{}
    response := responses.NewEventListResponse(baseInfo, events)
    responseJ, err := responses.ToJSON(response)
    if err != nil {
        log.Println("[ERR] (GetEventListHandler) Error marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[OUT] (GetEventListHandler) All OK")
    helper.Respond([]byte(responseJ), w)
}
