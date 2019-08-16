package obj

type Campaign struct {
    Type       int64 `json:"campaignType"`       // value from enums.CampaignType*
    Content    int64 `json:"campaignContent"`    // Appears to be generic parameter for various use based on CampaignType. No known casting enums
    SubContent int64 `json:"campaignSubContent"` // Same as above, but used much more sparingly
    BeginDate  int64 `json:"campaignStarTime"`   // time at which campaign starts
    EndDate    int64 `json:"campaignEndTime"`    // time at which campaign ends
}
