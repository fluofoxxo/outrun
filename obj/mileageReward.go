package obj

import (
    "strconv"

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

func NewMileageReward(itemID int64, point int64) MileageReward {
    // TODO: Should not be used for legitimate purposes, only rapid development!
    return MileageReward{
        enums.ItemTypeUnknown,
        strconv.Itoa(int(itemID)),
        1,
        point,
        690, // 11 minutes, 30 seconds
    }
}
