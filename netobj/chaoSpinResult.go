package netobj

import "github.com/fluofoxxo/outrun/obj"

type ChaoSpinResult struct {
    WonPrize ChaoSpinPrize `json:"getChao"`  // chao or character
    ItemList []obj.Item    `json:"itemList"` // TODO: what does this do?
    ItemWon  int64         `json:"itemWon"`  // probably index of item in ItemList
}

func DefaultChaoSpinResultNoItems(wonPrize ChaoSpinPrize) ChaoSpinResult {
    return ChaoSpinResult{
        wonPrize,
        []obj.Item{},
        -1, // TODO: if something breaks, it might be because of this
    }
}
