package handlers

import (
    "log"
    "net/http"

    "github.com/fluofoxxo/outrun/consts"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/objects"
    "github.com/fluofoxxo/outrun/responses"
)

func GetCampaignListHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: discover what this does for the game
    // for now, return a default response (empty).
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    response := responses.NewCampaignListResponse(baseInfo, []objects.Campaign{})
    responseJ, err := responses.ToJSON(response)
    if err != nil {
        log.Println("[ERR] (GetCampaignListHandler) Error marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
    }
    log.Println("[OUT] (GetCampaignListHandler) All OK")
    helper.Respond([]byte(responseJ), w)
}

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
