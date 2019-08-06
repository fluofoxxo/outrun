package objects

import (
    "github.com/fluofoxxo/outrun/playerdata"
)

type WheelOptions struct {
    Items                []string         `json:"items"`
    Item                 []int64          `json:"item"`
    ItemWeight           []int64          `json:"itemWeight"`
    ItemWon              int64            `json:"itemWon"`
    NextFreeSpin         int64            `json:"nextFreeSpin"` // midnight (start of next day)
    SpinCost             int64            `json:"spinCost"`
    RouletteRank         int64            `json:"rouletteRank"`
    NumRouletteToken     int64            `json:"numRouletteToken"`
    NumJackpotRing       int64            `json:"numJackpotRing"`
    NumRemainingRoulette int64            `json:"numRemainingRoulette`
    ItemList             []pdata.ItemInfo `json:"itemList"`
}
