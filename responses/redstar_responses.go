package responses

import (
    "github.com/fluofoxxo/outrun/objects"
)

type RedStarExchangeListNoCampaignResponse struct {
    BaseResponse
    ItemList      []objects.RedStarItemSansCampaign `json:"itemList"`
    TotalItems    int64                             `json:"totalItems"` // equal to length of ItemList
    MonthPurchase int64                             `json:"monthPurchase"`
    Birthday      string                            `json:"birthday"`
}

func MakeRedStarExchangeListNoCampaignResponse(base BaseInfo, itemList []objects.RedStarItemSansCampaign, birthday string) RedStarExchangeListNoCampaignResponse {
    br := NewBaseResponse(base)
    //il := objects.RedStarItemSansCampaign{} // TODO: what does this do?
    rselr := RedStarExchangeListNoCampaignResponse{
        br,
        itemList,
        int64(len(itemList)),
        0, // TODO: find out what this does
        birthday,
    }
    return rselr
}

func NewRedStarExchangeListNoCampaignResponse(base BaseInfo, itemList []objects.RedStarItemSansCampaign) RedStarExchangeListNoCampaignResponse {
    return MakeRedStarExchangeListNoCampaignResponse(
        base,
        itemList,
        "1900-1-1", // TODO: dummy birthday. Replace with real birthday.
    )
}

type RedStarExchangeListResponse struct {
    BaseResponse
    ItemList      []objects.RedStarItem `json:"itemList"`
    TotalItems    int64                 `json:"totalItems"` // equal to length of ItemList
    MonthPurchase int64                 `json:"monthPurchase"`
    Birthday      string                `json:"birthday"`
}

func MakeRedStarExchangeListResponse(base BaseInfo, itemList []objects.RedStarItem, birthday string) RedStarExchangeListResponse {
    br := NewBaseResponse(base)
    rselr := RedStarExchangeListResponse{
        br,
        itemList,
        int64(len(itemList)),
        0,
        birthday,
    }
    return rselr
}

func NewRedStarExchangeListResponse(base BaseInfo, itemList []objects.RedStarItem) RedStarExchangeListResponse {
    return MakeRedStarExchangeListResponse(
        base,
        itemList,
        "1900-1-1", // TODO: dummy birthday. Replace with real birthday.
    )
}
