package objects

type RedStarItem struct {
    StoreItemID string   `json:"storeItemId"`
    ItemID      string   `json:"itemId"`
    NumItem     int64    `json:"numItem"`
    Price       int64    `json:"price"`
    PriceDisp   string   `json:"priceDisp"`
    ProductID   string   `json:"productId"`
    Campaign    Campaign `json:"campaign"`
}

func MakeRedStarItem(storeItemId, itemId string, numItem, price int64, priceDisp, productId string, campaign Campaign) RedStarItem {
    return RedStarItem{
        storeItemId,
        itemId,
        numItem,
        price,
        priceDisp,
        productId,
        campaign,
    }
}

type RedStarItemSansCampaign struct {
    StoreItemID string `json:"storeItemId"`
    ItemID      string `json:"itemId"`
    NumItem     int64  `json:"numItem"`
    Price       int64  `json:"price"`
    PriceDisp   string `json:"priceDisp"`
    ProductID   string `json:"productId"`
    //Campaign    string `json:"campaign"`  // may need to be a null string
}

func MakeRedStarItemSansCampaign(storeItemId, itemId string, numItem, price int64, priceDisp, productId string) RedStarItemSansCampaign {
    return RedStarItemSansCampaign{
        storeItemId,
        itemId,
        numItem,
        price,
        priceDisp,
        productId,
    }
}
