package handlers

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/fluofoxxo/outrun/consts"
    "github.com/fluofoxxo/outrun/cryption"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/objects"
    "github.com/fluofoxxo/outrun/requests"
    "github.com/fluofoxxo/outrun/responses"
)

func GetRedStarExchangeListHandler(w http.ResponseWriter, r *http.Request) {
    recv := cryption.GetReceivedMessage(r)

    var request requests.RedStarExchangeListRequest
    err := json.Unmarshal(recv, &request)
    if err != nil {
        log.Println("[ERR] (GetRedStarExchangeListHandler) HTTP 400 - JSON could not be formed: " + err.Error())
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid request"))
        return
    }
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    if request.ItemType == 0 {
        response := responses.NewRedStarExchangeListResponse(baseInfo, []objects.RedStarItem{})
        responseJ, err := responses.ToJSON(response)
        if err != nil {
            log.Println("[ERR] (GetRedStarExchangeListHandler) IT=0 - Error marshalling: " + err.Error())
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte("Internal server error"))
            return
        }
        log.Println("[OUT] (GetRedStarExchangeListHandler) IT=0 - All OK")
        helper.Respond([]byte(responseJ), w)
    } else if request.ItemType == 1 {
        response := responses.MakeRedStarExchangeListNoCampaignResponse(
            baseInfo,
            objects.DEFAULT_REDSTARITEMS_ITEMTYPE_1,
            "1900-1-1", // TODO: replace this with actual birthday
        )
        responseJ, err := responses.ToJSON(response)
        if err != nil {
            log.Println("[ERR] (GetRedStarExchangeListHandler) IT=1 - Error marshalling: " + err.Error())
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte("Internal server error"))
            return
        }
        log.Println("[OUT] (GetRedStarExchangeListHandler) IT=1 - All OK")
        helper.Respond([]byte(responseJ), w)
    } else if request.ItemType == 2 {
        response := responses.MakeRedStarExchangeListResponse(
            baseInfo,
            objects.DEFAULT_REDSTARITEMS_ITEMTYPE_2,
            "1900-1-1", // TODO: replace with actual birthday
        )
        responseJ, err := responses.ToJSON(response)
        if err != nil {
            log.Println("[ERR] (GetRedStarExchangeListHandler) IT=2 - Error marshalling: " + err.Error())
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte("Internal server error"))
            return
        }
        log.Println("[OUT] (GetRedStarExchangeListHandler) IT=2 - All OK")
        helper.Respond([]byte(responseJ), w)
    } else if request.ItemType == 4 {
        response := responses.MakeRedStarExchangeListNoCampaignResponse(
            baseInfo,
            objects.DEFAULT_REDSTARITEMS_ITEMTYPE_4,
            "1900-1-1", // TODO: replace this with actual birthday
        )
        responseJ, err := responses.ToJSON(response)
        if err != nil {
            log.Println("[ERR] (GetRedStarExchangeListHandler) IT=4 - Error marshalling: " + err.Error())
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte("Internal server error"))
            return
        }
        log.Println("[OUT] (GetRedStarExchangeListHandler) IT=4 - All OK")
        helper.Respond([]byte(responseJ), w)
    } else {
        log.Println("[OUT] Improper or no ItemType. Malformed request from client.")
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid request"))
    }
}
