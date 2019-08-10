package handlers

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/fluofoxxo/outrun/consts"
    "github.com/fluofoxxo/outrun/cryption"
    "github.com/fluofoxxo/outrun/db"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/objects"
    "github.com/fluofoxxo/outrun/requests"
    "github.com/fluofoxxo/outrun/responses"
)

func GetFreeItemListHandler(w http.ResponseWriter, r *http.Request) {
    recv := cryption.GetReceivedMessage(r)

    var request requests.BasicRequest
    err := json.Unmarshal(recv, &request)
    if err != nil {
        log.Println("[ERR] (GetFreeItemListHandler) Error unmarshalling: " + err.Error())
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid request"))
        return
    }

    // TODO: this is player agnostic for now!
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    response := responses.DefaultFreeItemListResponse(baseInfo)
    if err != nil {
        log.Println("[ERR] (GetFreeItemListHandler) Error making response: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    respJ, err := responses.ToJSON(response)
    if err != nil {
        log.Println("[ERR] (GetFreeItemListHandler) Error marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[OUT] (GetFreeItemListHandler) All OK")
    helper.Respond([]byte(respJ), w)
}

func QuickPostGameResultsHandler(w http.ResponseWriter, r *http.Request) {
    recv := cryption.GetReceivedMessage(r)

    var request requests.QuickPostGameResultsRequest
    err := json.Unmarshal(recv, &request)
    if err != nil {
        log.Println("[ERR] (QuickPostGameResultsHandler) Error unmarshalling: " + err.Error())
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid request"))
        return
    }

    sid := request.SessionID
    score := request.Score
    rings := request.NumRings
    distance := request.Distance
    animals := request.NumAnimals
    player, err := db.SessionIDToPlayer(sid)
    if err != nil {
        log.Println("[ERR] (QuickPostGameResultsHandler) Error getting player: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }

    player.PlayerState.NumRings += rings
    player.PlayerState.Energy -= 1
    player.PlayerState.TotalDistance += distance
    player.PlayerState.NumAnimals += animals
    if player.PlayerState.QuickTotalHighScore < score {
        player.PlayerState.QuickTotalHighScore = score
    }
    if player.PlayerState.MaximumDistance < distance {
        player.PlayerState.MaximumDistance = distance
    }

    err = db.SavePlayer(player)
    if err != nil {
        log.Println("[ERR] (QuickPostGameResultsHandler) Error saving player: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }

    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    response, err := responses.DefaultQuickPostGameResultsResponse(baseInfo, player)
    if err != nil {
        log.Println("[ERR] (QuickPostGameResultsHandler) Error making response: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    respJ, err := responses.ToJSON(response)
    if err != nil {
        log.Println("[ERR] (QuickPostGameResultsHandler) Error marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[OUT] (QuickPostGameResultsHandler) All OK")
    helper.Respond([]byte(respJ), w)
}

func QuickActStartHandler(w http.ResponseWriter, r *http.Request) {
    recv := cryption.GetReceivedMessage(r)

    var request requests.QuickActStartRequest
    err := json.Unmarshal(recv, &request)
    if err != nil {
        log.Println("[ERR] (QuickActStartHandler) Error unmarshalling: " + err.Error())
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid request"))
        return
    }
    sid := request.SessionID
    player, err := db.SessionIDToPlayer(sid)
    if err != nil {
        log.Println("[ERR] (QuickActStartHandler) Error getting player: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
    }

    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    response := responses.DefaultQuickActStartResponse(baseInfo, player.PlayerState)
    responseJ, err := responses.ToJSON(response)
    if err != nil {
        log.Println("[ERR] (QuickActStartHandler) Error marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
    }
    log.Println("[OUT] (QuickActStartHandler) All OK")
    helper.Respond([]byte(responseJ), w)
}

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
