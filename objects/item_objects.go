package objects

import (
    "github.com/fluofoxxo/outrun/playerdata"
)

type FreeItem struct {
    pdata.ItemInfo
    MaxItem int64 `json:"maxItem,string"`
}

func NewFreeItem(itemId string, numItem, maxItem int64) FreeItem {
    ii := pdata.MakeItemInfo(itemId, numItem)
    fi := FreeItem{ii, maxItem}
    return fi
}
