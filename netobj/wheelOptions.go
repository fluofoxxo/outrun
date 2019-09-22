package netobj

import (
    "strconv"

    "github.com/fluofoxxo/outrun/enums"
    "github.com/fluofoxxo/outrun/obj"
    "github.com/jinzhu/now"
)

type WheelOptions struct {
    Items                []string   `json:"items"`
    Item                 []int64    `json:"item"`
    ItemWeight           []int64    `json:"itemWeight"`
    ItemWon              int64      `json:"itemWon"`
    NextFreeSpin         int64      `json:"nextFreeSpin"` // midnight (start of next day)
    SpinCost             int64      `json:"spinCost"`
    RouletteRank         int64      `json:"rouletteRank"`
    NumRouletteToken     int64      `json:"numRouletteToken"`
    NumJackpotRing       int64      `json:"numJackpotRing"`
    NumRemainingRoulette int64      `json:"numRemainingRoulette"`
    ItemList             []obj.Item `json:"itemList"`
}

func DefaultWheelOptions() WheelOptions {
    // TODO: Modifying this seems like a good way of figuring out what the game thinks each ID means in terms of items.
    // const the below
    //items := []string{"200000", "120000", "120001", "120002", "200000", "900000", "120003", "120004"}
    items := []string{
        strconv.Itoa(enums.IDTypeItemRouletteWin),
        strconv.Itoa(enums.IDTypeItemRouletteWin),
        strconv.Itoa(enums.IDTypeItemRouletteWin),
        strconv.Itoa(enums.IDTypeItemRouletteWin),
        strconv.Itoa(enums.IDTypeItemRouletteWin),
        strconv.Itoa(enums.IDTypeItemRouletteWin),
        strconv.Itoa(enums.IDTypeItemRouletteWin),
        strconv.Itoa(int(enums.ItemIDInvincible)),
    }
    item := []int64{1, 2, 2, 2, 1, 3, 2, 2}
    itemWeight := []int64{1250, 1250, 1250, 1250, 1250, 1250, 1250, 1250}
    itemWon := int64(0)
    nextFreeSpin := now.EndOfDay().Unix() + 1 // midnight
    spinCost := int64(87)
    rouletteRank := int64(enums.WheelRankSuper)
    numRouletteToken := int64(6)
    numJackpotRing := int64(61213)
    numRemainingRoulette := int64(11)
    itemList := []obj.Item{}
    out := WheelOptions{
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
    return out
}
