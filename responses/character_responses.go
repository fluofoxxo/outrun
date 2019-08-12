package responses

import (
    "github.com/fluofoxxo/outrun/playerdata"
)

type CharacterStateResponse struct {
    BaseResponse
    CharacterState []pdata.CharacterState `json:"characterState"`
}

func NewCharacterStateResponse(base BaseInfo, cStates []pdata.CharacterState) CharacterStateResponse {
    br := NewBaseResponse(base)
    csr := CharacterStateResponse{
        br,
        cStates,
    }
    return csr
}

type ChangeCharacterResponse struct {
    BaseResponse
    PlayerState pdata.PlayerState `json:"playerState"`
}

func NewChangeCharacterResponse(base BaseInfo, playerState pdata.PlayerState) ChangeCharacterResponse {
    br := NewBaseResponse(base)
    return ChangeCharacterResponse{
        br,
        playerState,
    }
}

func DefaultChangeCharacterResponse(base BaseInfo, player pdata.Player) ChangeCharacterResponse {
    return NewChangeCharacterResponse(base, player.PlayerState)
}
