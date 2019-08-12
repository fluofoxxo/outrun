package responses

import (
	"time"

	"github.com/fluofoxxo/outrun/consts"
	"github.com/fluofoxxo/outrun/objects"
	"github.com/fluofoxxo/outrun/playerdata"
)

type ActResponseBase struct {
	BaseResponse
	PlayerState  pdata.PlayerState  `json:"playerState"`
	CampaignList []objects.Campaign `json:"campaignList"`
}

func NewActResponseBase(base BaseInfo, playerState pdata.PlayerState, campaignList []objects.Campaign) ActResponseBase {
	br := NewBaseResponse(base)
	return ActResponseBase{
		br,
		playerState,
		campaignList,
	}
}

type QuickActStartResponse struct {
	ActResponseBase
}

func NewQuickActStartResponse(base BaseInfo, playerState pdata.PlayerState, campaignList []objects.Campaign) QuickActStartResponse {
	arb := NewActResponseBase(base, playerState, campaignList)
	qasr := QuickActStartResponse{
		arb,
	}
	return qasr
}

func DefaultQuickActStartResponse(base BaseInfo, playerState pdata.PlayerState) QuickActStartResponse {
	return NewQuickActStartResponse(base, playerState, []objects.Campaign{})
}

type ActStartResponse struct {
	ActResponseBase
	DistanceFriendList []objects.Friend `json:"distanceFriendList"` // TODO: discover if this is really what is needed
}

func NewActStartResponse(base BaseInfo, playerState pdata.PlayerState, campaignList []objects.Campaign, distFriends []objects.Friend) ActStartResponse {
	arb := NewActResponseBase(base, playerState, campaignList)
	asr := ActStartResponse{
		arb,
		distFriends, // TODO: find out if distFriends should just be the DistanceFriendList received from request
	}
	return asr
}

func DefaultActStartResponse(base BaseInfo, playerState pdata.PlayerState) ActStartResponse {
	return NewActStartResponse(base, playerState, []objects.Campaign{}, []objects.Friend{})
}

type FreeItemListResponse struct {
	BaseResponse
	FreeItemList []objects.FreeItem `json:"freeItemList"`
	ResetTime    int64              `json:"resetTime"`
}

func NewFreeItemListResponse(base BaseInfo, fil []objects.FreeItem, rt int64) FreeItemListResponse {
	br := NewBaseResponse(base)
	return FreeItemListResponse{
		br,
		fil,
		rt,
	}
}

var defaultFreeItems = makeDefaultFreeItems()

func DefaultFreeItemListResponse(base BaseInfo) FreeItemListResponse {
	resetTime := time.Now().UTC().Unix() + 7200 // two hours from now (For debugging)
	filr := NewFreeItemListResponse(
		base,
		defaultFreeItems,
		resetTime,
	)
	return filr
}

func makeDefaultFreeItems() []objects.FreeItem {
	dfi := []objects.FreeItem{}
	for _, id := range consts.ITEM_IDS {
		fi := objects.NewFreeItem(id, 1, 1)
		dfi = append(dfi, fi)
	}
	return dfi
}
