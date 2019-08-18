package obj

import "github.com/fluofoxxo/outrun/enums"

type ChaoWheelOptions struct {
    Rarity               []int64    `json:"rarity"`
    ItemWeight           []int64    `json:"itemWeight"`
    CampaignList         []Campaign `json:"campaignList"`
    SpinCost             int64      `json:"spinCost"`
    ChaoRouletteType     int64      `json:"chaoRouletteType"` // value from enums.ChaoWheelType*
    NumSpecialEgg        int64      `json:"numSpecialEgg"`
    RouletteAvailable    int64      `json:"rouletteAvailable"`    // flag
    NumChaoRouletteToken int64      `json:"numChaoRouletteToken"` // number of premium roulette tickets
    NumChaoRoulette      int64      `json:"numChaoRoulette"`      // == 0 --> chaoWheelOptions.IsTutorial
}

func NewChaoWheelOptions(rarity, itemWeight []int64, campaignList []Campaign, spinCost, chaoRouletteType, numSpecialEgg, rouletteAvailable, numChaoRouletteToken, numChaoRoulette int64) ChaoWheelOptions {
    return ChaoWheelOptions{
        rarity,
        itemWeight,
        campaignList,
        spinCost,
        chaoRouletteType,
        numSpecialEgg,
        rouletteAvailable,
        numChaoRouletteToken,
        numChaoRoulette,
    }
}

func DefaultChaoWheelOptions() ChaoWheelOptions {
    rarity := []int64{2, 1, 100, 1, 2, 1, 100, 1}
    itemWeight := []int64{6, 17, 5, 17, 16, 17, 5, 17}
    campaignList := []Campaign{}
    spinCost := int64(1234)
    chaoRouletteType := enums.ChaoWheelTypeSpecial
    numSpecialEgg := int64(183)
    rouletteAvailable := int64(1)
    numChaoRouletteToken := int64(47)
    numChaoRoulette := int64(1)
    return NewChaoWheelOptions(rarity, itemWeight, campaignList, spinCost, chaoRouletteType, numSpecialEgg, rouletteAvailable, numChaoRouletteToken, numChaoRoulette)
}
