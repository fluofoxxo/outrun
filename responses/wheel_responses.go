package responses

import (
    "github.com/fluofoxxo/outrun/consts"
    "github.com/fluofoxxo/outrun/objects"
    "github.com/fluofoxxo/outrun/playerdata"
    "github.com/jinzhu/now"
)

type WheelOptionsResponse struct {
    BaseResponse
    objects.WheelOptions `json:"wheelOptions"`
}

func MakeWheelOptionsResponse(base BaseInfo, items []string, item, itemWeight []int64, itemWon, nextFreeSpin, spinCost, rouletteRank, numRouletteToken, numJackpotRing, numRemainingRoulette int64, itemList []pdata.ItemInfo) WheelOptionsResponse {
    br := NewBaseResponse(base)
    wo := objects.WheelOptions{
        items,
        item,
        itemWeight,
        itemWon,
        nextFreeSpin,
        spinCost,
        rouletteRank,
        numRouletteToken,
        numJackpotRing,
        numRemainingRoulette,
        itemList,
    }
    wor := WheelOptionsResponse{
        br,
        wo,
    }
    return wor
}

var WHEEL_ITEMLIST = []pdata.ItemInfo{} // could not be included in consts due to cyclical dependencies

func NewWheelOptionsResponse(base BaseInfo) WheelOptionsResponse {
    nextFreeSpin := now.EndOfDay().Unix() + 1 // end of day plus one second (midnight)
    wor := MakeWheelOptionsResponse(
        base,
        consts.WHEEL_ITEMS,
        consts.WHEEL_ITEM,
        consts.WHEEL_ITEMWEIGHT,
        consts.WHEEL_ITEMWON,
        nextFreeSpin,
        consts.WHEEL_SPINCOST,
        consts.WHEEL_ROULETTERANK,
        consts.WHEEL_NUMROULETTETOKEN,
        consts.WHEEL_NUMJACKPOTRING,
        consts.WHEEL_NUMREMAININGROULETTE,
        WHEEL_ITEMLIST,
    )
    return wor
}
