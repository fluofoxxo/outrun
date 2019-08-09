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

func GetTickerHandler(w http.ResponseWriter, r *http.Request) {

}

func GetInformationHandler(w http.ResponseWriter, r *http.Request) {
    // player agnostic
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    resp := responses.DefaultInformationResponse(baseInfo)
    respJ, err := responses.ToJSON(resp)
    if err != nil {
        log.Println("[ERR] (GetInformationHandler) Error in JSON marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[OUT] (GetInformationHandler) All OK")
    helper.Respond([]byte(respJ), w)
}

func GetVariousParameterHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: responds with a default value for now. fix this to tie to user
    baseInfo := responses.NewBaseInfo(consts.EM_OK, 0, 0)
    resp := responses.NewVariousParameterResponse(baseInfo)
    respJ, err := responses.ToJSON(resp)
    if err != nil {
        log.Println("[ERR] (GetVariousParameterHandler) Error in JSON marshalling: " + err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte("Internal server error"))
        return
    }
    log.Println("[OUT] (GetVariousParameterHandler) All OK")
    helper.Respond([]byte(respJ), w)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    recv := cryption.GetReceivedMessage(r)

    var request requests.NotPlayingBase // try getting the login base
    err := json.Unmarshal(recv, &request)
    if err != nil {
        log.Println("[ERR] (LoginHandler) HTTP 400 - JSON could not be formed: " + err.Error())
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid request"))
        return
    }

    rUserID := request.LineAuth.UserID
    rPassword := request.LineAuth.Password
    log.Println("userID=" + rUserID + ", password=" + rPassword)

    baseInfo := responses.NewBaseInfo("", 0, 0) // TODO: THE STATUS CODE NEEDS TO BE FILLED OUT BEFORE RESPONDING. THIS IS ONLY DONE ONCE
    logme := " (userID=" + rUserID + ", password=" + rPassword + ")"
    // check if it's a new account
    if rUserID == "0" && rPassword == "" { // gate LoginAlpha
        // new account
        uid, _, password, key := db.NewAccount()
        baseInfo.StatusCode = consts.SC_INVALID_PASSWORD
        baseInfo.SetErrorMessage(consts.EM_REQUEST_PASSWORD)
        respS := responses.NewLoginStage1FirstResponse(baseInfo, uid, password, key)
        respJBytes, err := json.Marshal(respS)
        if err != nil {
            log.Println("[ERR] (LoginHandler) Error marshalling user data: " + err.Error())
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte("Internal server error"))
            return
        }
        //w.Write(respJBytes)
        log.Println("[OUT] (LoginHandler) Completed Gate LoginAlpha" + logme)
        helper.Respond(respJBytes, w)

    } else if rUserID == "0" && rPassword != "" { // gate LoginBravo
        // invalid request
        log.Println("[ERR] (LoginHandler) HTTP 400 - Invalid request")
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid request")) // TODO: replace with proper status code
        return

    } else if rUserID != "0" && rPassword == "" { // gate LoginCharlie
        // game is looking to log in
        // for now, let's pretend like it worked no matter what
        // WE NEED TO TEST THIS!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
        baseInfo.StatusCode = consts.SC_INVALID_PASSWORD
        baseInfo.SetErrorMessage(consts.EM_REQUEST_PASSWORD)
        log.Println("[OUT] (LoginHandler) Entering gate LoginCharlie")
        player, err := db.GetPlayerByUID(rUserID)
        if err != nil {
            log.Println("[ERR] (LoginHandler)" + logme + " Error fetching user data from players/" + rUserID + ": " + err.Error())
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte("Internal server error")) // TODO: replace with proper status code
            return
        }
        response := responses.NewLoginStage1Response(baseInfo, player.Key)
        responseJ, err := json.Marshal(response)
        if err != nil {
            log.Println("[ERR] (LoginHandler)" + logme + " Error marshalling: " + err.Error())
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte("Internal server error")) // TODO: replace with proper status code
            return
        }
        //w.Write(responseJ)
        log.Println("[OUT] (LoginHandler) Completed gate LoginCharlie" + logme)
        helper.Respond(responseJ, w)

    } else if rUserID != "0" && rPassword != "" { // gate LoginDelta
        // game is trying to log in using given key
        // for now, let's pretend like it worked no matter what
        baseInfo.SetErrorMessage(consts.EM_OK)
        sid, err := db.AssignSessionID(rUserID)
        if err != nil {
            log.Println("[ERR] (LoginHandler) (LoginDelta, " + logme + ") Error assigning SID: " + err.Error())
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte("Internal server error"))
            return
        }
        //response := responses.NewLoginSuccessResponse(baseInfo, sid, "DUMMY USERNAME")
        response := responses.NewLoginSuccessResponse(baseInfo, sid, "")
        responseJ, err := json.Marshal(response)
        if err != nil {
            log.Println("[ERR] (LoginHandler)" + logme + " Error marshalling: " + err.Error())
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte("Internal server error"))
            return
        }
        log.Println("[OUT] (LoginHandler) Completed gate LoginDelta" + logme)
        //w.Write([]byte("Gnight gurl. I'll see you tomorrow."))
        helper.Respond(responseJ, w)
    }
    //w.Write([]byte("SOPTOPTSPOST"))
}
