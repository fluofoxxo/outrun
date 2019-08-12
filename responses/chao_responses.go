package responses

import (
	"encoding/json"

	"github.com/fluofoxxo/outrun/consts"
	"github.com/fluofoxxo/outrun/objects"
	"github.com/fluofoxxo/outrun/playerdata"
	"github.com/jinzhu/now"
)

type ChaoStateResponse struct {
	BaseResponse
	ChaoState []pdata.Chao `json:"chaoState"`
}

func NewChaoStateResponse(base BaseInfo, chao []pdata.Chao) ChaoStateResponse {
	br := NewBaseResponse(base)
	csr := ChaoStateResponse{
		br,
		chao,
	}
	return csr
}

type ChaoWheelOptionsResponse struct {
	BaseResponse
	ChaoWheelOptions objects.ChaoWheelOptions `json:"chaoWheelOptions"`
}

func NewChaoWheelOptionsResponse(base BaseInfo, chaoWheelOptions objects.ChaoWheelOptions) ChaoWheelOptionsResponse {
	br := NewBaseResponse(base)
	cwor := ChaoWheelOptionsResponse{
		br,
		chaoWheelOptions,
	}
	return cwor
}

func DefaultChaoWheelOptionsResponse(base BaseInfo) ChaoWheelOptionsResponse {
	startTime := now.BeginningOfDay().UTC().Unix() + 32400 // 12 AM + 9 hours = 9 AM
	endTime := startTime + 86399                           // 23:59:59 later
	return NewChaoWheelOptionsResponse(
		base,
		objects.ChaoWheelOptions{
			consts.CHAOWHEEL_RARITY,
			consts.CHAOWHEEL_ITEMWEIGHT,
			consts.CHAOWHEEL_SPINCOST,
			consts.CHAOWHEEL_CHAOROULETTETYPE,
			consts.CHAOWHEEL_NUMSPECIALEGG,
			consts.CHAOWHEEL_ROULETTEAVAILABLE,
			[]objects.ChaoCampaign{},
			consts.CHAOWHEEL_NUMCHAOROULETTE,
			consts.CHAOWHEEL_NUMCHAOROULETTETOKEN,
			startTime,
			endTime,
		},
	)
}

type PrizeChaoWheelResponse struct {
	BaseResponse
	PrizeList []objects.ChaoPrize `json:"prizeList"`
}

func NewPrizeChaoWheelResponse(base BaseInfo, prizeList []objects.ChaoPrize) PrizeChaoWheelResponse {
	br := NewBaseResponse(base)
	pcwr := PrizeChaoWheelResponse{
		br,
		prizeList,
	}
	return pcwr
}

func DefaultPrizeChaoWheelResponse(base BaseInfo) PrizeChaoWheelResponse {
	var prizeList []objects.ChaoPrize
	err := json.Unmarshal([]byte(consts.JSON_DEFAULT_PRIZECHAOWHEEL_PRIZELIST), &prizeList)
	if err != nil {
		panic(err)
	}
	return NewPrizeChaoWheelResponse(base, prizeList)
}

type EquipChaoResponse struct {
    BaseResponse
    PlayerState pdata.PlayerState `json:"playerState"`
}

func NewEquipChaoResponse(base BaseInfo, playerState pdata.PlayerState) EquipChaoResponse {
    br := NewBaseResponse(base)
    return EquipChaoResponse{
        br,
        playerState,
    }
}

func DefaultEquipChaoResponse(base BaseInfo, player pdata.Player) EquipChaoResponse {
    return NewEquipChaoResponse(base, player.PlayerState)
}
