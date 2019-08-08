package objects

type Information struct {
    ID       int64  `json:"id"`
    Priority int64  `json:"priority"`
    Start    int64  `json:"start"`
    End      int64  `json:"end"`
    Param    string `json:"param"`
}

func NewInformation(id, priority, start, end int64, param string) Information {
    i := Information{
        id,
        priority,
        start,
        end,
        param,
    }
    return i
}

type OperatorInfo struct { // Appears to be an administrative message to a specific player or players.
}
