package obj

type OperatorScore struct {
    Operator    int64  `json:"operator,string"`
    Number      int64  `json:"number,string"`
    PresentList []Item `json:"presentList,string"`
}

func NewOperatorScore(operator, number int64, presentList []Item) OperatorScore {
    return OperatorScore{
        operator,
        number,
        presentList,
    }
}
