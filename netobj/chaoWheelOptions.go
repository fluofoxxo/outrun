package netobj

import (
	"github.com/fluofoxxo/outrun/consts"
	"github.com/fluofoxxo/outrun/enums"
	"github.com/fluofoxxo/outrun/obj"
	"github.com/jinzhu/now"
)

type ChaoWheelOptions struct {
	Rarity               []int64        `json:"rarity"`
	ItemWeight           []int64        `json:"itemWeight"`
	CampaignList         []obj.Campaign `json:"campaignList"`
	SpinCost             int64          `json:"spinCost"`
	ChaoRouletteType     int64          `json:"chaoRouletteType"` // value from enums.ChaoWheelType*
	NumSpecialEgg        int64          `json:"numSpecialEgg"`
	RouletteAvailable    int64          `json:"rouletteAvailable"`    // flag
	NumChaoRouletteToken int64          `json:"numChaoRouletteToken"` // number of premium roulette tickets
	NumChaoRoulette      int64          `json:"numChaoRoulette"`      // == 0 --> chaoWheelOptions.IsTutorial
	StartTime            int64          `json:"startTime"`            // TODO: Is this needed?
	EndTime              int64          `json:"endTime"`              // TODO: Is this needed?
}

func NewChaoWheelOptions(rarity, itemWeight []int64, campaignList []obj.Campaign, spinCost, chaoRouletteType, numSpecialEgg, rouletteAvailable, numChaoRouletteToken, numChaoRoulette, startTime, endTime int64) ChaoWheelOptions {
	return ChaoWheelOptions{
		rarity,
		itemWeight,
		campaignList,
		spinCost,
		chaoRouletteType,
		numSpecialEgg,
		rouletteAvailable,
		numChaoRouletteToken,
		numChaoRoulette,
		startTime,
		endTime,
	}
}

func DefaultChaoWheelOptions(playerState PlayerState) ChaoWheelOptions {
	rarity := []int64{2, 1, 100, 1, 2, 1, 100, 1} // l2 Chao, l1 Chao, character...
	//rarity := []int64{0, 1, 2, 100, 0, 1, 2, 100}      // TODO: REMOVE!
	itemWeight := []int64{6, 17, 5, 17, 16, 17, 5, 17} // Could possibly fake these, but the logic shouldn't allow it to happen
	campaignList := []obj.Campaign{}
	chaoRouletteType := enums.ChaoWheelTypeNormal
	numSpecialEgg := playerState.ChaoEggs
	rouletteAvailable := int64(1)
	numChaoRouletteToken := playerState.NumChaoRouletteTicket
	spinCost := consts.ChaoRouletteTicketCost
	if numChaoRouletteToken <= 0 { // if out of chao roulette tickets
		// use higher value for the red rings
		spinCost = consts.ChaoRouletteRedRingCost
	}
	numChaoRoulette := int64(1)
	startTime := now.BeginningOfDay().UTC().Unix() + 32400 // 12 AM + 9 hours = 9 AM
	endTime := startTime + 86399                           // 23:59:59 later
	return NewChaoWheelOptions(rarity, itemWeight, campaignList, spinCost, chaoRouletteType, numSpecialEgg, rouletteAvailable, numChaoRouletteToken, numChaoRoulette, startTime, endTime)
}
