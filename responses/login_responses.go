package responses

import (
    "github.com/fluofoxxo/outrun/playerdata"
    "time"
)

/*
Login Stage 1 notes:
  - Status code is -10104 (consts.SC_INVALID_PASSWORD)
  - Used to get a key in order to verify the password's integrity
  - Base deviations:
    - key (pulled from user)
  - When executed for the first time (userId == "0")...
    - Also respond with:
      - userId (new)
      - password (new)
      - key (new)
      - countryId = "1"
      - countryCode = "US"
*/

type LoginStage1Response struct {
    BaseResponse
    Key string `json:"key"`
}

func NewLoginStage1Response(base BaseInfo, key string) LoginStage1Response {
    br := NewBaseResponse(base)
    ls1r := LoginStage1Response{
        br,
        key,
    }
    return ls1r
}

type LoginStage1FirstResponse struct {
    BaseResponse
    UserID      string `json:"userId"`
    Password    string `json:"password"`
    Key         string `json:"key"`
    CountryID   int64  `json:"countryId,string"`
    CountryCode string `json:"countryCode"`
}

func NewLoginStage1FirstResponse(base BaseInfo, uid, password, key string) LoginStage1FirstResponse {
    br := NewBaseResponse(base)
    ls1fr := LoginStage1FirstResponse{
        br,
        uid,
        password,
        key,
        1,
        "US",
    }
    return ls1fr
}

// NOTE: sessionTimeLimit seems to be server_time + 1 hr (3600)

type LoginSuccessResponse struct {
    BaseResponse
    SessionID            string         `json:"sessionId"`
    SessionTimeLimit     int64          `json:"sessionTimeLimit"`
    EnergyRecoveryTime   int64          `json:"energyRecveryTime,string"` // yes, that misspelling really is in the json
    EnergyRecoveryMax    int64          `json:"energyRecoveryMax,string"`
    InviteBasicIncentive pdata.ItemInfo `json:"inviteBasicIncentiv"`
    Username             string         `json:"userName"`
}

func NewLoginSuccessResponse(base BaseInfo, sid, username string) LoginSuccessResponse {
    br := NewBaseResponse(base)
    incentive := pdata.MakeItemInfo("900000", 10)
    lsr := LoginSuccessResponse{
        br,
        sid,
        time.Now().UTC().Unix() + 3600,
        900,
        5,
        incentive,
        username,
    }
    return lsr
}
