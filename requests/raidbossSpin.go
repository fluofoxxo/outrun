package requests

type ItemStockNumRequest struct {
    Base
    EventID int64   `json:"eventId,string"`
    ItemIDs []int64 `json:"itemIdList"`
}
