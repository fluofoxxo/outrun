package responses

import (
	"github.com/fluofoxxo/outrun/netobj"
	"github.com/fluofoxxo/outrun/obj"
	"github.com/fluofoxxo/outrun/obj/constobjs"
	"github.com/fluofoxxo/outrun/responses/responseobjs"
)

type ChaoWheelOptionsResponse struct {
	BaseResponse
	ChaoWheelOptions obj.ChaoWheelOptions `json:"chaoWheelOptions"`
}

func ChaoWheelOptions(base responseobjs.BaseInfo, chaoWheelOptions obj.ChaoWheelOptions) ChaoWheelOptionsResponse {
	baseResponse := NewBaseResponse(base)
	out := ChaoWheelOptionsResponse{
		baseResponse,
		chaoWheelOptions,
	}
	return out
}

func DefaultChaoWheelOptions(base responseobjs.BaseInfo) ChaoWheelOptionsResponse {
	chaoWheelOptions := obj.DefaultChaoWheelOptions()
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
