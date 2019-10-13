package netobj

import (
    "github.com/fluofoxxo/outrun/logic/roulette"
)

type ChaoRouletteGroup struct {
    ChaoWheelOptions ChaoWheelOptions `json:"ORN_lastChaoWheelOptions"` // actual wheel options for this wheel
    WheelChao        []string         `json:"ORN_wheelChao"`            // what Chao/characters are in this wheel
    ChaoRouletteInfo RouletteInfo     `json:"ORN_chaoRouletteInfo"`     // may not be needed
}

func DefaultChaoRouletteGroup(playerState PlayerState, exclusions []string) ChaoRouletteGroup {
    chaoWheelOptions := DefaultChaoWheelOptions(playerState)
    wheelChao, err := roulette.GetRandomChaoRouletteItems(chaoWheelOptions.Rarity, exclusions) // populate based on given rarities
    if err != nil {
        panic(err) // TODO: Find a better way to handle error. Hard to manage since the player creators don't already output errors
    }
    chaoRouletteInfo := DefaultRouletteInfo()
    return ChaoRouletteGroup{
        chaoWheelOptions,
        wheelChao,
        chaoRouletteInfo,
    }
}
