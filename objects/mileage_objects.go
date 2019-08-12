package objects

import (
	"github.com/fluofoxxo/outrun/consts"
)

type MileageFriend struct { // still don't know what fields this has
}

type MileageIncentive struct {
	Type     int64  `json:"type"`    // consts.MILEAGE_INCENTIVE_*
	ItemID   string `json:"itemId"`  // like pdata.ItemInfo?
	NumItem  int64  `json:"numItem"` // like pdata.ItemInfo?
	PointID  int64  `json:"pointId"`
	FriendID string `json:"friendId,omitempty"`
}

func DefaultMileageIncentive() MileageIncentive {
	// TODO: this is ONLY for debugging right now.
	return MileageIncentive{
		consts.MILEAGE_INCENTIVE_NONE,
		consts.ITEM_IDS[0],
		1,
		0,  // ???
		"", // for battling?
	}
}
