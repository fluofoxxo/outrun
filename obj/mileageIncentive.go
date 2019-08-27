package obj

type MileageIncentive struct {
    Type     int64  `json:"type"`
    ItemID   string `json:"itemId"`
    FriendID string `json:"friendId,omitempty"`
    NumItem  int64  `json:"numItem"`
    PointID  int64  `json:"pointId"`
}

func DefaultMileageIncentive() MileageIncentive {
    // TODO: this is ONLY for debugging right now.
    return MileageIncentive{
        0,
        "110000",
        "", // for battling?
        1,
        0, // ???
    }
}
