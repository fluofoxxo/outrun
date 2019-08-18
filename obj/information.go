package obj

// TODO: create a param object that makes it easier to create informations

type Information struct {
    ID       int64  `json:"id"`
    Priority int64  `json:"priority"`
    Start    int64  `json:"start"`
    End      int64  `json:"end"`
    Param    string `json:"param"`
}

func NewInformation(id, priority, start, end int64, param string) Information {
    return Information{
        id,
        priority,
        start,
        end,
        param,
    }
}
