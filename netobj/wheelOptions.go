package netobj

import (
    "math/rand"

    "github.com/fluofoxxo/outrun/consts"
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

func DefaultWheelOptions(numRouletteTicket, rouletteCountInPeriod int64) WheelOptions {
    // TODO: Modifying this seems like a good way of figuring out what the game thinks each ID means in terms of items.
    // const the below
    // NOTE: Free spins occur when numRemainingRoulette > numRouletteToken
    //items := []string{"200000", "120000", "120001", "120002", "200000", "900000", "120003", "120004"}
    items := []string{}
    item := []int64{}
    for _ = range make([]byte, 8) { // loop 8 times
        randomItem := consts.RandomItemListNormalWheel[rand.Intn(len(consts.RandomItemListNormalWheel))]
        randomItemAmount := consts.NormalWheelItemAmountRange[randomItem].GetRandom()
        items = append(items, randomItem)
        item = append(item, randomItemAmount)
    }
    /*
       item := []int64{1, 2, 2, 2, 1, 3, 2, 2}
    */
    itemWeight := []int64{1250, 1250, 1250, 1250, 1250, 1250, 1250, 1250}
    //itemWon := int64(0)
    itemWon := int64(rand.Intn(len(items)))
    nextFreeSpin := now.EndOfDay().Unix() + 1 // midnight
    spinCost := int64(87)
    rouletteRank := int64(enums.WheelRankSuper) // TODO: change this
    //numRouletteToken := playerState.NumRouletteTicket
    numRouletteToken := numRouletteTicket
    numJackpotRing := int64(consts.RouletteJackpotRings)
    // TODO: get rid of logic here!
    numRemainingRoulette := numRouletteTicket + consts.RouletteFreeSpins - rouletteCountInPeriod
    if numRemainingRoulette < numRouletteToken {
        numRemainingRoulette = numRouletteToken
    }
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
