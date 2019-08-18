package obj

type ConsumedItem struct {
    Item
    ID string `json:"consumedItemId"`
}

func NewConsumedItem(item Item, id string) ConsumedItem {
    return ConsumedItem{
        item,
        id,
    }
}
