package responses

import (
    "github.com/fluofoxxo/outrun/obj"
    "github.com/fluofoxxo/outrun/responses/responseobjs"
)

type RedStarExchangeListResponse struct {
    BaseResponse
    ItemList      []obj.RedStarItem `json:"itemList"`
    TotalItems    int64             `json:"totalItems"`
    MonthPurchase int64             `json:"monthPurchase"`
    Birthday      string            `json:"birthday"`
}

func RedStarExchangeList(base responseobjs.BaseInfo, itemList []obj.RedStarItem, monthPurchase int64, birthday string) RedStarExchangeListResponse {
    baseResponse := NewBaseResponse(base)
    totalItems := int64(len(itemList))
    return RedStarExchangeListResponse{
        baseResponse,
        itemList,
        totalItems,
        monthPurchase,
        birthday,
    }
}

func DefaultRedStarExchangeList(base responseobjs.BaseInfo) RedStarExchangeListResponse {
    itemList := []obj.RedStarItem{}
    monthPurchase := int64(0)
    birthday := "1900-1-1"
    return RedStarExchangeList(base, itemList, monthPurchase, birthday)
}
