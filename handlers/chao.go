package handlers

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/fluofoxxo/outrun/cryption"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/responses"
    "github.com/fluofoxxo/outrun/consts"
    "github.com/fluofoxxo/outrun/db"
    "github.com/fluofoxxo/outrun/requests"
)

func GetPrizeChaoWheelSpinHandler(w http.ResponseWriter, r *http.Request) {
    // this appears to be player agnostic.
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    response := responses.DefaultPrizeChaoWheelResponse(baseInfo)
    responseJ, err := responses.ToJSON(response)
    if err != nil {
        log.Println("[ERR] (GetPrizeChaoWheelSpinHandler) Error marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[OUT] (GetPrizeChaoWheelSpinHandler) All OK")
    helper.Respond([]byte(responseJ), w)
}

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

func EquipChaoHandler(w http.ResponseWriter, r *http.Request) {
    recv := cryption.GetReceivedMessage(r)

    var request requests.EquipChaoRequest
    err := json.Unmarshal(recv, &request)
    if err != nil {
        log.Println("[ERR] (EquipChaoHandler) Error unmarshalling: " + err.Error())
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid request"))
        return
    }

    player, err := db.SessionIDToPlayer(request.SessionID)
    if err != nil {
        log.Println("[ERR] (EquipChaoHandler) Error getting player: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }

    if request.MainChaoID != "-1" {
        player.MainChaoID = request.MainChaoID
    }
    if request.SubChaoID != "-1" {
        player.SubChaoID = request.SubChaoID
    }

    err = db.SavePlayer(player)
    if err != nil {
        log.Println("[ERR] (EquipChaoHandler) Error saving player: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }

    baseInfo := responses.NewBaseInfo(consts.EM_OK, 1, 0)
    response := responses.DefaultEquipChaoResponse(baseInfo, player)
    respJ, err := responses.ToJSON(response)
    if err != nil {
        log.Println("[ERR] (EquipChaoHandler) Error marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[OUT] (EquipChaoHandler) No issues")
    helper.Respond([]byte(respJ), w)
}
