package obj

import (
    "strconv"
    "time"

    "github.com/fluofoxxo/outrun/enums"
)

type OperatorMessage struct {
    ID         string      `json:"messageId"`
    Content    string      `json:"contents"`
    Item       MessageItem `json:"item"`
    ExpireTime int64       `json:"expireTime"`
}

func DefaultOperatorMessage() OperatorMessage {
    id := "8575819"
    content := "A daily challenge reward."
    item := NewMessageItem(
        strconv.Itoa(int(enums.ItemIDInvincible)),
        135,
        0,
        0,
    )
    expireTime := time.Now().Unix() + 12600 // three and a half hours from now
    return OperatorMessage{
        id,
        content,
        item,
        expireTime,
    }
}
