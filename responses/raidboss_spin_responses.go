package responses

import (
    "github.com/fluofoxxo/outrun/playerdata"
)

type ItemStockNumResponse struct {
    BaseResponse
    ItemStockList []pdata.ItemInfo `json:"itemStockList"`
}

func NewItemStockNumResponse(base BaseInfo, itemStockList []pdata.ItemInfo) ItemStockNumResponse {
    br := NewBaseResponse(base)
    return ItemStockNumResponse{
        br,
        itemStockList,
    }
}

func DefaultItemStockNumResponse(base BaseInfo) ItemStockNumResponse {
    itemStockList := []pdata.ItemInfo{
        pdata.MakeItemInfo("230000", 11),
        pdata.MakeItemInfo("240000", 12),
        pdata.MakeItemInfo("220000", 13),
    }
    return NewItemStockNumResponse(base, itemStockList)
}
