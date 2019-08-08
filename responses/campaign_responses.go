package responses

import "github.com/fluofoxxo/outrun/objects"

type CampaignListResponse struct {
	BaseResponse
	CampaignList []objects.Campaign `json:"campaignList"`
}

func NewCampaignListResponse(base BaseInfo, campaignList []objects.Campaign) CampaignListResponse {
	// TODO: verify that objects.Campaign is actually what the game is expecting
	br := NewBaseResponse(base)
	clr := CampaignListResponse{
		br,
		campaignList,
	}
	return clr
}
