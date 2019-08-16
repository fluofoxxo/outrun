package responses

import (
    "time"

    "github.com/fluofoxxo/outrun/netobj"
    "github.com/fluofoxxo/outrun/obj"
    "github.com/fluofoxxo/outrun/responses/responseobjs"
)

type LoginCheckKeyResponse struct {
    BaseResponse
    Key string `json:"key"`
}

func LoginCheckKey(base responseobjs.BaseInfo, key string) LoginCheckKeyResponse {
    baseResponse := NewBaseResponse(base)
    out := LoginCheckKeyResponse{
        baseResponse,
        key,
    }
    return out
}

type LoginRegisterResponse struct {
    BaseResponse
    UserID      string `json:"userId"`
    Password    string `json:"password"`
    Key         string `json:"key"`
    CountryID   int64  `json:"countryId,string"`
    CountryCode string `json:"countryCode"`
}

func LoginRegister(base responseobjs.BaseInfo, uid, password, key string) LoginRegisterResponse {
    // TODO: fetch correct country code and whatnot
    baseResponse := NewBaseResponse(base)
    out := LoginRegisterResponse{
        baseResponse,
        uid,
        password,
        key,
        1,
        "US",
    }
    return out
}

type LoginSuccessResponse struct {
    BaseResponse
    Username             string   `json:"userName"`
    SessionID            string   `json:"sessionId"`
    SessionTimeLimit     int64    `json:"sessionTimeLimit"`
    EnergyRecoveryTime   int64    `json:"energyRecveryTime,string"` // misspelling is _actually_ in the game!
    EnergyRecoveryMax    int64    `json:"energyRecoveryMax,string"` // seconds until energy recovers
    InviteBasicIncentive obj.Item `json:"inviteBasicIncentiv"`
}

func LoginSuccess(base responseobjs.BaseInfo, sid, username string) LoginSuccessResponse {
    baseResponse := NewBaseResponse(base)
    out := LoginSuccessResponse{
        baseResponse,
        username,
        sid,
        time.Now().Unix() + 3600, // hour from now  // TODO: does this need to be UTC?
        360,                      // 6 minutes from now, regen energy
        17171,
        obj.NewItem("900000", 13),
    }
    return out
}

type VariousParameterResponse struct {
    BaseResponse
    netobj.PlayerVarious
}

func VariousParameter(base responseobjs.BaseInfo, player netobj.Player) VariousParameterResponse {
    baseResponse := NewBaseResponse(base)
    out := VariousParameterResponse{
        baseResponse,
        player.PlayerVarious,
    }
    return out
}
