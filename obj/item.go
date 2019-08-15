package obj

type Item struct {
	ID     string `json:"itemId"`
	Amount int64  `json:"numItem,string"`
}
