package objects

import (
	"github.com/fluofoxxo/outrun/playerdata"
)

type Ticker struct { // TODO: find out what this contains!
}

type LoginBonusStatus struct {
	NumLogin      int64 `json:"numLogin"`
	NumBonus      int64 `json:"numBonus"`
	LastBonusTime int64 `json:"lastBonusTime"`
}

type SelectReward struct {
	ItemList []pdata.ItemInfo `json:"itemList"`
}

type LoginBonusReward struct {
	SelectRewardList []SelectReward `json:"selectRewardList"`
}
