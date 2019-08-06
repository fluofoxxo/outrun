package responses

import (
    "github.com/fluofoxxo/outrun/objects"
)

type MessageListResponse struct {
    BaseResponse
    MessageList          []objects.Message `json:"messageList"`
    TotalMessage         int64             `json:"totalMessage"`
    OperatorMessageList  []objects.Message `json:"operatorMessageList"`
    TotalOperatorMessage int64             `json:"totalOperatorMessage"`
}

func NewMessageListResponse(base BaseInfo) MessageListResponse {
    br := NewBaseResponse(base)
    mlr := MessageListResponse{
        br,
        []objects.Message{}, // TODO: Find a good replacement for this
        0,
        []objects.Message{}, // TODO: Find a good replacement for this
        0,
    }
    return mlr
}
