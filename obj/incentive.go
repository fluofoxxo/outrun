package obj

type Incentive struct {
    Item
    IncentiveCount int64 `json:"numIncentiveCont"`
}

func NewIncentive(item Item, count int64) Incentive {
    return Incentive{
        item,
        count,
    }
}
