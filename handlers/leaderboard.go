package handlers

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/fluofoxxo/outrun/consts"
    "github.com/fluofoxxo/outrun/cryption"
    "github.com/fluofoxxo/outrun/db"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/requests"
    "github.com/fluofoxxo/outrun/responses"
)

func GetLeagueData(w http.ResponseWriter, r *http.Request) {
    // player agnostic
    recv := cryption.GetReceivedMessage(r)
    var request requests.LeaderboardRequest
    err := json.Unmarshal(recv, &request)
    if err != nil {
        log.Println("[ERR] (GetLeagueData) Error in JSON unmarshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    mode := request.Mode // mode is first 0, then 1 (but appears to return the same response)
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    resp := responses.DefaultLeagueDataResponse(baseInfo, mode)
    respJ, err := responses.ToJSON(resp)
    if err != nil {
        log.Println("[ERR] (GetLeagueData) Error in JSON marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[OUT] (GetLeagueData) All OK")
    helper.Respond([]byte(respJ), w)
}

func GetWeeklyLeaderboardEntriesHandler(w http.ResponseWriter, r *http.Request) {
    // player agnostic
    recv := cryption.GetReceivedMessage(r)
    var request requests.LeaderboardEntriesRequest
    err := json.Unmarshal(recv, &request)
    if err != nil {
        log.Println("[ERR] (GetWeeklyLeaderboardEntriesHandler) Error in JSON unmarshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    mode := request.Mode // mode is first 0, then 1 (but appears to return the same response)
    sid := request.SessionID
    uid, err := db.SessionIDToUID(sid)
    if err != nil {
        log.Println("[ERR] (GetWeeklyLeaderboardEntriesHandler) Error in fetching SID: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    resp := responses.DefaultLeaderboardEntriesResponse(baseInfo, uid, mode)
    respJ, err := responses.ToJSON(resp)
    if err != nil {
        log.Println("[ERR] (GetWeeklyLeaderboardEntriesHandler) Error in JSON marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[OUT] (GetWeeklyLeaderboardOptionsHandler) All OK")
    helper.Respond([]byte(respJ), w)
}

func GetWeeklyLeaderboardOptionsHandler(w http.ResponseWriter, r *http.Request) {
    // player agnostic
    recv := cryption.GetReceivedMessage(r)
    var request requests.LeaderboardRequest
    err := json.Unmarshal(recv, &request)
    if err != nil {
        log.Println("[ERR] (GetWeeklyLeaderboardOptionsHandler) Error in JSON unmarshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    mode := request.Mode // mode is first 0, then 1 (but appears to return the same response)
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    resp := responses.DefaultLeaderboardResponse(baseInfo, mode)
    respJ, err := responses.ToJSON(resp)
    if err != nil {
        log.Println("[ERR] (GetWeeklyLeaderboardOptionsHandler) Error in JSON marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[OUT] (GetWeeklyLeaderboardOptionsHandler) All OK")
    helper.Respond([]byte(respJ), w)
}
