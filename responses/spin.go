package responses

import (
	"github.com/fluofoxxo/outrun/netobj"
	"github.com/fluofoxxo/outrun/responses/responseobjs"
)

type WheelOptionsResponse struct {
	BaseResponse
	WheelOptions netobj.WheelOptions `json:"wheelOptions"`
}

func WheelOptions(base responseobjs.BaseInfo, wheelOptions netobj.WheelOptions) WheelOptionsResponse {
	baseResponse := NewBaseResponse(base)
	out := WheelOptionsResponse{
		baseResponse,
		wheelOptions,
	}
	return out
}

type WheelSpinResponse struct {
	BaseResponse
	PlayerState    netobj.PlayerState  `json:"playerState"`
	CharacterState []netobj.Character  `json:"characterState"`
	ChaoState      []netobj.Chao       `json:"chaoState"`
	WheelOptions   netobj.WheelOptions `json:"wheelOptions"`
}

func WheelSpin(base responseobjs.BaseInfo, playerState netobj.PlayerState, characterState []netobj.Character, chaoState []netobj.Chao, wheelOptions netobj.WheelOptions) WheelSpinResponse {
	baseResponse := NewBaseResponse(base)
	return WheelSpinResponse{
		baseResponse,
		playerState,
		characterState,
		chaoState,
		wheelOptions,
	}
}

func DefaultWheelSpin(base responseobjs.BaseInfo, player netobj.Player) WheelSpinResponse {
	// TODO: remove me! I am no longer being used.
	wheelOptions := netobj.DefaultWheelOptions(player.PlayerState.NumRouletteTicket, player.RouletteInfo.RouletteCountInPeriod)
	playerState := player.PlayerState
	characterState := player.CharacterState
	chaoState := player.ChaoState
	return WheelSpin(
		base,
		playerState,
		characterState,
		chaoState,
		wheelOptions,
	)
}
