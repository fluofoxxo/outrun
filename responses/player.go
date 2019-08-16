package responses

import (
    "github.com/fluofoxxo/outrun/netobj"
    "github.com/fluofoxxo/outrun/responses/responseobjs"
)

type PlayerStateResponse struct {
    BaseResponse
    PlayerState netobj.PlayerState `json:"playerState"`
}

func PlayerState(base responseobjs.BaseInfo, playerState netobj.PlayerState) PlayerStateResponse {
    baseResponse := NewBaseResponse(base)
    out := PlayerStateResponse{
        baseResponse,
        playerState,
    }
    return out
}
