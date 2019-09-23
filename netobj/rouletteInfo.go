package netobj

import "github.com/jinzhu/now"

type RouletteInfo struct {
    RoulettePeriodEnd     int64 `json:"ORN_roulettePeriodEnd"`
    RouletteCountInPeriod int64 `json:"ORN_rouletteCountInPeriod"`
    GotJackpotThisPeriod  bool  `json:"ORN_gotJackpotThisPeriod"`
}

func DefaultRouletteInfo() RouletteInfo {
    roulettePeriodEnd := now.EndOfDay().Unix()
    rouletteCountInPeriod := int64(0)
    gotJackpotThisPeriod := false
    return RouletteInfo{
        roulettePeriodEnd,
        rouletteCountInPeriod,
        gotJackpotThisPeriod,
    }
}
