package responses

import (
    "strconv"

    "github.com/fluofoxxo/outrun/enums"
    "github.com/fluofoxxo/outrun/obj"
    "github.com/fluofoxxo/outrun/responses/responseobjs"
)

type DailyChallengeDataResponse struct {
    BaseResponse
    IncentiveList          []obj.Incentive `json:"incentiveList"`
    IncentiveListCount     int64           `json:"incentiveListCont"`
    NumDailyChallengeCount int64           `json:"numDilayChalCont"`
    NumDailyChallengeDay   int64           `json:"numDailyChalDay"`
    MaxDailyChallengeDay   int64           `json:"maxDailyChalDay"`
    EndTime                int64           `json:"chalEndTime"`
}

func DailyChallengeData(base responseobjs.BaseInfo) DailyChallengeDataResponse {
    ilSrc := []int64{enums.ItemIDMagnet, enums.ItemIDMagnet, enums.ItemIDMagnet, enums.ItemIDMagnet, enums.ItemIDMagnet, enums.ItemIDMagnet, enums.ItemIDMagnet} // must be length of seven!
    incentiveList := []obj.Incentive{}
    for amountSrc, id := range ilSrc {
        item := obj.NewItem(strconv.Itoa(int(id)), int64(amountSrc+1))
        incentive := obj.NewIncentive(
            item,
            int64(amountSrc+1),
        )
        incentiveList = append(incentiveList, incentive)
    }
    incentiveListCount := int64(len(incentiveList))
    numDailyChallengeCount := int64(0)
    numDailyChallengeDay := int64(2)
    maxDailyChallengeDay := int64(10) // is this how many you can get a day? In that case, doesn't 10 make no sense?
    endTime := int64(1470322800)      // 08/04/2016 @ 3:00PM (UTC)
    baseResponse := NewBaseResponse(base)
    return DailyChallengeDataResponse{
        baseResponse,
        incentiveList,
        incentiveListCount,
        numDailyChallengeCount,
        numDailyChallengeDay,
        maxDailyChallengeDay,
        endTime,
    }
}
