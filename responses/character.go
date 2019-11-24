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

type UpgradeCharacterResponse struct {
	BaseResponse
	PlayerState    netobj.PlayerState `json:"playerState"`
	CharacterState []netobj.Character `json:"characterState"`
}

func UpgradeCharacter(base responseobjs.BaseInfo, playerState netobj.PlayerState, characterState []netobj.Character) UpgradeCharacterResponse {
	baseResponse := NewBaseResponse(base)
	return UpgradeCharacterResponse{
		baseResponse,
		playerState,
		characterState,
	}
}

func DefaultUpgradeCharacter(base responseobjs.BaseInfo, player netobj.Player) UpgradeCharacterResponse {
	playerState := player.PlayerState
	characterState := player.CharacterState
	return UpgradeCharacter(
		base,
		playerState,
		characterState,
	)
}

// UnlockedCharacter's response is the exact same as UpgradeCharacterResponse, it seems
