package responses

import (
    "github.com/fluofoxxo/outrun/consts"
    "github.com/fluofoxxo/outrun/objects"
)

var CHAL_INCENTIVELIST = getBasicIncentives() // TODO: putting in consts would be cyclical ; THIS SHOULD BE CHANGED IF WE NEED TO CHANGE THE DAILY CHALLENGE

type DailyChalDataResponse struct {
    BaseResponse
    IncentiveList     []objects.Incentive `json:"incentiveList"`
    IncentiveListCont int64               `json:"incentiveListCont"` // does cont mean count?
    NumDailyChalCont  int64               `json:"numDailyChalCont"`
    NumDailyChalDay   int64               `json:"numDailyChalDay"`
    MaxDailyChalDay   int64               `json:"maxDailyChalDay"`
    ChalEndTime       int64               `json:"chalEndTime"`
}

func NewDailyChalDataResponse(base BaseInfo) DailyChalDataResponse {
    // daily challenges are independent of the player, so constants are probably appropriate.
    chalEndTime := int64(1470322800) // 08/04/2016 @ 3:00pm (UTC)
    br := NewBaseResponse(base)
    dcdr := DailyChalDataResponse{
        br,
        CHAL_INCENTIVELIST,
        consts.CHAL_DAILY_INCENTIVELISTCONT,
        consts.CHAL_DAILY_NUMDAILYCHALCONT,
        consts.CHAL_DAILY_NUMDAILYCHALDAY,
        consts.CHAL_DAILY_MAXDAILYCHALDAY,
        chalEndTime,
    }
    return dcdr
}
