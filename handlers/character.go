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

func ChangeCharacterHandler(w http.ResponseWriter, r *http.Request) {
	recv := cryption.GetReceivedMessage(r)

	var request requests.ChangeCharacterRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		log.Println("[ERR] (ChangeCharacterHandler) Error unmarshalling: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request"))
		return
	}

	player, err := db.SessionIDToPlayer(request.SessionID)
	if err != nil {
		log.Println("[ERR] (ChangeCharacterHandler) Error getting player: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	if request.MainCharacterID != "-1" {
		player.MainCharaID = request.MainCharacterID
	}
	if request.SubCharacterID != "-1" {
		player.SubCharaID = request.SubCharacterID
	}

	err = db.SavePlayer(player)
	if err != nil {
		log.Println("[ERR] (ChangeCharacterHandler) Error saving player: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	baseInfo := responses.NewBaseInfo(consts.EM_OK, 1, 0)
	response := responses.DefaultChangeCharacterResponse(baseInfo, player)
	respJ, err := responses.ToJSON(response)
	if err != nil {
		log.Println("[ERR] (ChangeCharacterHandler) Error marshalling: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	log.Println("[OUT] (ChangeCharacterHandler) No issues")
	helper.Respond([]byte(respJ), w)
}
