package handlers

import (
    "encoding/json"
    "log"
    "net/http"

    "fluofoxxo.pw/subsonic2/consts"
    "github.com/fluofoxxo/outrun/cryption"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/requests"
    "github.com/fluofoxxo/outrun/responses"
)

func GetMileageDataHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: implement this for each player, as this includes
    // unique data such as episode, chapter, max score, etc.
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    response := responses.DefaultMileageDataResponse(baseInfo)
    responseJ, err := responses.ToJSON(response)
    if err != nil {
        log.Println("[ERR] (GetMileageDataHandler) Error marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[OUT] (GetMileageDataHandler) All OK")
    helper.Respond([]byte(responseJ), w)
}

func GetMileageRewardHandler(w http.ResponseWriter, r *http.Request) {
    recv := cryption.GetReceivedMessage(r)

    var request requests.MileageRewardRequest
    err := json.Unmarshal(recv, &request)
    if err != nil {
        log.Println("[ERR] (GetMileageRewardHandler) Error unmarshalling: " + err.Error())
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid request"))
        return
    }

    // TODO: this is player agnostic for now!
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    response := responses.DefaultMileageRewardResponse(baseInfo)
    responseJ, err := responses.ToJSON(response)
    if err != nil {
        log.Println("[ERR] (GetFreeItemListHandler) Error making response: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[OUT] (GetMileageRewardHandler) All OK")
    helper.Respond([]byte(responseJ), w)
}
