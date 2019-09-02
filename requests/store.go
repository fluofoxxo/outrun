package requests

type RedStarExchangeListRequest struct {
    Base
    ItemType int64 `json:"itemType,string"`
}

type RedStarExchange struct {
    Base
    ItemID string `json:"itemId"`
}
