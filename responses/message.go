package responses

import (
    "github.com/fluofoxxo/outrun/obj"
    "github.com/fluofoxxo/outrun/responses/responseobjs"
)

type MessageListResponse struct {
    BaseResponse
    MessageList           []obj.Message         `json:"messageList"`
    TotalMessages         int64                 `json:"totalMessage"`
    OperatorMessageList   []obj.OperatorMessage `json:"operatorMessageList"`
    TotalOperatorMessages int64                 `json:"totalOperatorMessage"`
}

func MessageList(base responseobjs.BaseInfo, msgl []obj.Message, opmsgl []obj.OperatorMessage) MessageListResponse {
    baseResponse := NewBaseResponse(base)
    out := MessageListResponse{
        baseResponse,
        msgl,
        int64(len(msgl)),
        opmsgl,
        int64(len(opmsgl)),
    }
    return out
}

func DefaultMessageList(base responseobjs.BaseInfo) MessageListResponse {
    return MessageList(
        base,
        []obj.Message{},
        []obj.OperatorMessage{},
    )
}
