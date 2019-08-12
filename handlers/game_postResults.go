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

func PostGameResultsHandler(w http.ResponseWriter, r *http.Request) {
	recv := cryption.GetReceivedMessage(r)

	var request requests.PostGameResultsRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		log.Println("[ERR] (PostGameResultsHandler) Error unmarshalling: " + err.Error())
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
		log.Println("[ERR] (PostGameResultsHandler) Error getting player: " + err.Error())
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
		log.Println("[ERR] (PostGameResultsHandler) Error saving player: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
	response := responses.DefaultPostGameResultsResponse(baseInfo, player)
	/*
		if err != nil {
			log.Println("[ERR] (PostGameResultsHandler) Error making response: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
			return
		}
	*/
	respJ, err := responses.ToJSON(response)
	if err != nil {
		log.Println("[ERR] (PostGameResultsHandler) Error marshalling: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	log.Println("[OUT] (PostGameResultsHandler) All OK")
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
