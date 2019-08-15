package muxhandlers

import (
	"encoding/json"

	"github.com/fluofoxxo/outrun/db"
	"github.com/fluofoxxo/outrun/emess"
	"github.com/fluofoxxo/outrun/helper"
	"github.com/fluofoxxo/outrun/requests"
	"github.com/fluofoxxo/outrun/responses"
	"github.com/fluofoxxo/outrun/status"
)

func Login(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.LoginRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	uid := request.LineAuth.UserID
	password := request.LineAuth.Password

	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	if uid == "0" && password == "" {
		// LoginAlpha
		newPlayer := db.NewAccount()
	}
}
