package obj

import (
    "strconv"
    "time"

    "github.com/fluofoxxo/outrun/enums"
)

type MileageReward struct {
    Type      int64  `json:"type"`   // never used in game?
    ItemID    string `json:"itemId"` // TODO: integrate obj.Item as field instead of itemId and numItem?
    NumItem   int64  `json:"numItem"`
    Point     int64  `json:"point"`
    LimitTime int64  `json:"limitTime"` // timespan (sec.)
}

func DefaultMileageReward(point int64) MileageReward {
    return MileageReward{
        enums.ItemTypeInvincible,
        strconv.Itoa(int(enums.ItemIDInvincible)),
        1,
        point,
        690, // 11 minutes, 30 seconds
    }
}

func NewMileageReward(itype, itemID, numItem, point int64) MileageReward {
    // TODO: Should not be used for legitimate purposes, only rapid development!
    return MileageReward{
        itype, // should be enums.IncentiveType*
        strconv.Itoa(int(itemID)),
        numItem,
        point,
        time.Now().UTC().Unix() + 690, // 11 minutes, 30 seconds
    }
}
