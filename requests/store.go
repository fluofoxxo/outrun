package requests

type RedStarExchangeListRequest struct {
    Base
    ItemType int64 `json:"itemType,string"`
}
