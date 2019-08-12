package objects

import (
    "time"

    "github.com/fluofoxxo/outrun/consts"
    "github.com/fluofoxxo/outrun/playerdata"
)

type WheelOptions struct {
    Items                []string         `json:"items"` // array of string(ints) representing item IDs
    Item                 []int64          `json:"item"`  // Array of item quantities!!
    ItemWeight           []int64          `json:"itemWeight"`
    ItemWon              int64            `json:"itemWon"`      // index of item from Items
    NextFreeSpin         int64            `json:"nextFreeSpin"` // midnight (start of next day)
    SpinCost             int64            `json:"spinCost"`
    RouletteRank         int64            `json:"rouletteRank"`
    NumRouletteToken     int64            `json:"numRouletteToken"`
    NumJackpotRing       int64            `json:"numJackpotRing"`
    NumRemainingRoulette int64            `json:"numRemainingRoulette`
    ItemList             []pdata.ItemInfo `json:"itemList"`
}

func DefaultWheelOptions() WheelOptions {
    items := []string{"200000", "120000", "120001", "120002", "120003", "120004", "120005", "120006"} // item IDs
    item := []int64{0, 0, 0, 0, 0, 0, 0, 0}                                                           // item counts
    itemWeight := []int64{1, 1, 1, 1, 1, 1, 1, 1}                                                     // weight on the wheel
    itemWon := int64(5)                                                                               // index of item won
    nextFreeSpin := time.Now().UTC().Unix() + 60                                                      // one minute from now
    spinCost := int64(100)                                                                            // cost of next spin?
    rouletteRank := consts.WHEEL_RANK_SUPER
    numRouletteToken := int64(1)     // TODO: does this need to be a player-based variable for how many tokens they have?
    numJackpotRing := int64(10000)   // TODO: is this how many rings are in the jackpot? Should be easy to find out!
    numRemainingRoulette := int64(3) // unknown, possibly free spins left?
    itemList := pdata.DEFAULT_ITEMS  // TODO: maybe don't set this as default?
    return WheelOptions{
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
}
