package responses

import (
	"github.com/fluofoxxo/outrun/enums"
	"github.com/fluofoxxo/outrun/netobj"
	"github.com/fluofoxxo/outrun/obj"
	"github.com/fluofoxxo/outrun/obj/constobjs"
	"github.com/fluofoxxo/outrun/responses/responseobjs"
)

type ChaoWheelOptionsResponse struct {
	BaseResponse
	ChaoWheelOptions netobj.ChaoWheelOptions `json:"chaoWheelOptions"`
}

func ChaoWheelOptions(base responseobjs.BaseInfo, chaoWheelOptions netobj.ChaoWheelOptions) ChaoWheelOptionsResponse {
	baseResponse := NewBaseResponse(base)
	out := ChaoWheelOptionsResponse{
		baseResponse,
		chaoWheelOptions,
	}
	return out
}

func DefaultChaoWheelOptions(base responseobjs.BaseInfo) ChaoWheelOptionsResponse {
	// TODO: Assess if needed
	chaoWheelOptions := netobj.DefaultChaoWheelOptions() // TODO: !!! Change this to the player's LastChaoWheelOptions
	return ChaoWheelOptions(
		base,
		chaoWheelOptions,
	)
}

type PrizeChaoWheelResponse struct {
	BaseResponse
	PrizeList []obj.ChaoPrize `json:"prizeList"`
}

func PrizeChaoWheel(base responseobjs.BaseInfo, prizeList []obj.ChaoPrize) PrizeChaoWheelResponse {
	baseResponse := NewBaseResponse(base)
	out := PrizeChaoWheelResponse{
		baseResponse,
		prizeList,
	}
	return out
}

func DefaultPrizeChaoWheel(base responseobjs.BaseInfo) PrizeChaoWheelResponse {
	prizeList := constobjs.DefaultChaoPrizeWheelPrizeList
	return PrizeChaoWheel(base, prizeList)
}

type EquipChaoResponse struct {
	BaseResponse
	PlayerState netobj.PlayerState `json:"playerState"`
}

func EquipChao(base responseobjs.BaseInfo, playerState netobj.PlayerState) EquipChaoResponse {
	baseResponse := NewBaseResponse(base)
	return EquipChaoResponse{
		baseResponse,
		playerState,
	}
}

type ChaoWheelSpinResponse struct {
	BaseResponse
	PlayerState      netobj.PlayerState      `json:"playerState"`
	CharacterState   []netobj.Character      `json:"characterState"`
	ChaoState        []netobj.Chao           `json:"chaoState"` // also works with json:"chaoStatus"
	ChaoWheelOptions netobj.ChaoWheelOptions `json:"chaoWheelOptions"`
	ChaoSpinResults  []netobj.ChaoSpinResult `json:"chaoSpinResultList"` // Should only contain one element! Otherwise, ItemWon is interpreted as -1
}

func ChaoWheelSpin(base responseobjs.BaseInfo, playerState netobj.PlayerState, characterState []netobj.Character, chaoState []netobj.Chao, chaoWheelOptions netobj.ChaoWheelOptions, chaoSpinResults []netobj.ChaoSpinResult) ChaoWheelSpinResponse {
	baseResponse := NewBaseResponse(base)
	return ChaoWheelSpinResponse{
		baseResponse,
		playerState,
		characterState,
		chaoState,
		chaoWheelOptions,
		chaoSpinResults,
	}
}

func DefaultChaoWheelSpin(base responseobjs.BaseInfo, player netobj.Player) ChaoWheelSpinResponse {
	// WARN: Do not use for normal purposes!! This should only be used for debugging
	dummyPrize := netobj.CharacterIDToChaoSpinPrize(enums.CTStrShadow)
	chaoSpinResults := netobj.DefaultChaoSpinResultNoItems(dummyPrize)
	return ChaoWheelSpin(
		base,
		player.PlayerState,
		player.CharacterState,
		player.ChaoState,
		netobj.DefaultChaoWheelOptions(),
		[]netobj.ChaoSpinResult{chaoSpinResults},
	)
}
