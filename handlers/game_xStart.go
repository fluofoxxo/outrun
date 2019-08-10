package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"fluofoxxo.pw/subsonic2/consts"
	"github.com/fluofoxxo/outrun/cryption"
	"github.com/fluofoxxo/outrun/db"
	"github.com/fluofoxxo/outrun/helper"
	"github.com/fluofoxxo/outrun/requests"
	"github.com/fluofoxxo/outrun/responses"
)

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

func ActStartHandler(w http.ResponseWriter, r *http.Request) {
	recv := cryption.GetReceivedMessage(r)

	var request requests.ActStartRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		log.Println("[ERR] (ActStartHandler) Error unmarshalling: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request"))
		return
	}
	sid := request.SessionID
	player, err := db.SessionIDToPlayer(sid)
	if err != nil {
		log.Println("[ERR] (ActStartHandler) Error getting player: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	}

	baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
	response := responses.DefaultActStartResponse(baseInfo, player.PlayerState)
	responseJ, err := responses.ToJSON(response)
	if err != nil {
		log.Println("[ERR] (ActStartHandler) Error marshalling: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
	}
	log.Println("[OUT] (ActStartHandler) All OK")
	helper.Respond([]byte(responseJ), w)
}
