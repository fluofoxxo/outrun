package responses

import (
	"github.com/fluofoxxo/outrun/netobj"
	"github.com/fluofoxxo/outrun/responses/responseobjs"
)

type ChangeCharacterResponse struct {
	BaseResponse
	PlayerState netobj.PlayerState `json:"playerState"`
}

func ChangeCharacter(base responseobjs.BaseInfo, playerState netobj.PlayerState) ChangeCharacterResponse {
	baseResponse := NewBaseResponse(base)
	return ChangeCharacterResponse{
		baseResponse,
		playerState,
	}
}
