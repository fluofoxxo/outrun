package objects

import (
    "github.com/fluofoxxo/outrun/playerdata"
)

type Incentive struct {
    pdata.ItemInfo
    NumIncentiveCont int64 `json:"numIncentiveCont"`
}

func MakeIncentive(itemId string, numItem, numIncentiveCont int64) Incentive {
    ii := pdata.MakeItemInfo(itemId, numItem)
    i := Incentive{
        ii,
        numIncentiveCont,
    }
    return i
}
