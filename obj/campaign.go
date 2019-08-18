package obj

type Campaign struct {
    Type       int64 `json:"campaignType"`
    Content    int64 `json:"campaignContent"`
    SubContent int64 `json:"campaignSubContent"`
    StartTime  int64 `json:"campaignStartTime"`
    EndTime    int64 `json:"campaignEndTime"`
}

func NewCampaign(ctype, content, subcontent, startTime, endTime int64) Campaign {
    return Campaign{
        ctype,
        content,
        subcontent,
        startTime,
        endTime,
    }
}

func DefaultCampaign(ctype, content, subcontent int64) Campaign {
    return NewCampaign(
        ctype,
        content,
        subcontent,
        1565499998,
        1565499999,
    )
}
