package obj

type MessageItem struct {
    ID              string `json:"itemId"`
    Amount          int64  `json:"numItem"`
    AdditionalInfo1 int64  `json:"additionalInfo1"`
    AdditionalInfo2 int64  `json:"additionalInfo2"`
}

func NewMessageItem(id string, amount, ai1, ai2 int64) MessageItem {
    return MessageItem{
        id,
        amount,
        ai1,
        ai2,
    }
}
