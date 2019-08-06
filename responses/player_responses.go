package responses

import (
    "github.com/fluofoxxo/outrun/playerdata"
)

type PlayerStateResponse struct {
    BaseResponse
    pdata.PlayerState `json:"playerState"`
}

func NewPlayerStateResponse(base BaseInfo, playerState pdata.PlayerState) PlayerStateResponse {
    br := NewBaseResponse(base)
    psr := PlayerStateResponse{
        br,
        playerState,
    }
    return psr
}
