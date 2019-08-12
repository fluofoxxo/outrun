package objects

import (
	"strconv"

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

type MileageReward struct {
	Type      int64 `json:"type"`  // this field is asked for, but doesn't seem to be used in the game...
	Point     int64 `json:"point"` // point on the map? (MileageMapDataManager.GetMileageReward [m_point usage])
	ItemID    int64 `json:"itemId"`
	NumItem   int64 `json:"numItem"`
	LimitTime int64 `json:"limitTime"` // seconds until something with balloons and timers happen..?
}

func NewMileageReward(t, point, itemID, numItem, limitTime int64) MileageReward {
	return MileageReward{
		t,
		point,
		itemID,
		numItem,
		limitTime,
	}
}

func DefaultMileageReward() MileageReward {
	t := int64(0)
	point := int64(2) // TODO: this may cause errors if it indeed does utilize the map
	iid, _ := strconv.Atoi(consts.ITEM_IDS[0])
	itemID := int64(iid)
	numItem := int64(13)
	limitTime := int64(120)
	return NewMileageReward(t, point, itemID, numItem, limitTime)
}
