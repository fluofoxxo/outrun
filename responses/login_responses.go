package responses

import (
    "encoding/json"
    "time"

    "github.com/fluofoxxo/outrun/consts"
    "github.com/fluofoxxo/outrun/objects"
    "github.com/fluofoxxo/outrun/playerdata"
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

type TickerResponse struct {
    BaseResponse
    TickerList []objects.Ticker `json:"tickerList"`
}

func NewTickerResponse(base BaseInfo, tickerList []objects.Ticker) TickerResponse {
    br := NewBaseResponse(base)
    tr := TickerResponse{
        br,
        tickerList,
    }
    return tr
}

func DefaultTickerResponse(base BaseInfo) TickerResponse {
    return NewTickerResponse(
        base,
        []objects.Ticker{},
    )
}

type LoginBonusResponse struct {
    BaseResponse
    LoginBonusStatus          objects.LoginBonusStatus   `json:"loginBonusStatus"`
    StartTime                 int64                      `json:"startTime"`
    EndTime                   int64                      `json:"endTime"`
    LoginBonusRewardList      []objects.LoginBonusReward `json:"loginBonusRewardList"`
    FirstLoginBonusRewardList []objects.LoginBonusReward `json:"firstLoginBonusRewardList"`
    RewardID                  int64                      `json:"rewardId"`
    RewardDays                int64                      `json:"rewardDays"`
    FirstRewardDays           int64                      `json:"firstRewardDays"`
}

func NewLoginBonusResponse(base BaseInfo, lbs objects.LoginBonusStatus, st, et int64, lbrl, flbrl []objects.LoginBonusReward, rid, rd, frd int64) LoginBonusResponse {
    br := NewBaseResponse(base)
    lbr := LoginBonusResponse{
        br,
        lbs,
        st,
        et,
        lbrl,
        flbrl,
        rid,
        rd,
        frd,
    }
    return lbr
}

func DefaultLoginBonusResponse(base BaseInfo) LoginBonusResponse {
    // TODO: change these constants and find out what they do!!
    lbs := objects.LoginBonusStatus{2, 2, 1465830000}
    var loginBonusRewardList []objects.LoginBonusReward
    var firstLoginBonusRewardList []objects.LoginBonusReward
    err := json.Unmarshal([]byte(consts.JSON_DEFAULT_LOGINBONUS_LOGINBONUSREWARDLIST), &loginBonusRewardList)
    if err != nil {
        panic(err)
    }
    err = json.Unmarshal([]byte(consts.JSON_DEFAULT_LOGINBONUS_FIRSTLOGINBONUSREWARDLIST), &firstLoginBonusRewardList)
    if err != nil {
        panic(err)
    }
    return NewLoginBonusResponse(
        base,
        lbs,
        1465743600,
        1466348400,
        loginBonusRewardList,
        firstLoginBonusRewardList,
        -1,
        -1,
        -1,
    )
}
