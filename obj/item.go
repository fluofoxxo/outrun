package obj

type Item struct {
    ID     string `json:"itemId"`
    Amount int64  `json:"numItem,string"`
}

func NewItem(iid string, amount int64) Item {
    return Item{
        iid,
        amount,
    }
}
