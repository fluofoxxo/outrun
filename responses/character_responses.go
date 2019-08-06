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
