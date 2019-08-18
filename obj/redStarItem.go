package obj

type RedStarItem struct {
    StoreItemID string    `json:"storeItemId"`
    ItemID      string    `json:"itemId"`
    PriceDisp   string    `json:"priceDisp"`
    ProductID   string    `json:"productId"`
    NumItem     int64     `json:"numItem"`
    Price       int64     `json:"price"`
    Campaign    *Campaign `json:"campaign"` // use nil to remove campaign. But does this make the game freak out...?
}

func NewRedStarItem(storeItemId, itemId, priceDisp, productID string, numItem, price int64, campaign *Campaign) RedStarItem {
    return RedStarItem{
        storeItemId,
        itemId,
        priceDisp,
        productID,
        numItem,
        price,
        campaign,
    }
}
