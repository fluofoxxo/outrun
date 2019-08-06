package objects

type Campaign struct {
    CampaignType       int64 `json:"campaignType"`
    CampaignContent    int64 `json:"campaignContent"`
    CampaignSubContent int64 `json:"campaignSubContent"`
    CampaignStartTime  int64 `json:"campaignStartTime"` // 6 days from 8/3/2019 at 9:00. Indicates week long event
    CampaignEndTime    int64 `json:"campaignEndTime"`   // 15 days from 8/3/2019 at 8:59...?
}

func MakeCampaign(ctype, content, subcontent, startTime, endTime int64) Campaign {
    c := Campaign{
        ctype,
        content,
        subcontent,
        startTime,
        endTime,
    }
    return c
}

func NewCampaign(ctype, content, subcontent int64) Campaign {
    // TODO: uses placeholder numbers. Make these _not_ placeholder numbers.
    c := MakeCampaign(
        ctype,
        content,
        subcontent,
        1565499998,
        1565499999,
    )
    return c
}
