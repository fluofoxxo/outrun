package muxhandlers

import (
    "encoding/json"

    "github.com/fluofoxxo/outrun/emess"
    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/requests"
    "github.com/fluofoxxo/outrun/responses"
    "github.com/fluofoxxo/outrun/status"
)

func GetWeeklyLeaderboardOptions(helper *helper.Helper) {
    recv := helper.GetGameRequest()
    var request requests.LeaderboardRequest
    err := json.Unmarshal(recv, &request)
    if err != nil {
        helper.Err("Error unmarshalling", err)
        return
    }
    mode := request.Mode
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.DefaultWeeklyLeaderboardOptions(baseInfo, mode)
    err = helper.SendResponse(response)
    if err != nil {
        helper.InternalErr("Error sending response", err)
    }
}
