package responses

import (
	"log"
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

type QuickPostGameResultsResponse struct {
	BaseResponse
	PlayerState             pdata.PlayerState            `json:"playerState"`
	DailyChallengeIncentive []objects.Incentive          `json:"dailyChallengeIncentive"` // TODO: Confirm this is the correct type
	MessageList             []objects.Message            `json:"messageList"`
	TotalMessage            int64                        `json:"totalMessage"`
	OperatorMessageList     []objects.OperatorMessage    `json:"operatorMessageList"`
	TotalOperatorMessage    int64                        `json:"totalOperatorMessage"`
	PlayCharacterState      []objects.PlayCharacterState `json:"playCharacterState"`
}

func NewQuickPostGameResultsResponse(base BaseInfo, ps pdata.PlayerState, dci []objects.Incentive, ml []objects.Message, oml []objects.OperatorMessage, pcs []objects.PlayCharacterState) QuickPostGameResultsResponse {
	br := NewBaseResponse(base)
	return QuickPostGameResultsResponse{
		br,
		ps,
		dci,
		ml,
		int64(len(ml)),
		oml,
		int64(len(oml)),
		pcs,
	}
}

func DefaultQuickPostGameResultsResponse(base BaseInfo, player pdata.Player) (QuickPostGameResultsResponse, error) {
	// TODO: const this!
	ps := player.PlayerState
	dci := []objects.Incentive{}
	ml := []objects.Message{}
	mitem := objects.NewMessageItem(
		"900000",
		5,
		0,
		0,
	)
	oml := []objects.OperatorMessage{
		objects.NewOperatorMessage(
			"8575819",
			"A daily challenge reward.",
			mitem,
			time.Now().Unix()+3600,
		),
	}
	mchar, err := player.GetMainCharacter()
	if err != nil {
		log.Println("[MAJOR ERR] Something is wrong. GetMainCharacter couldn't find the main char. MainCharaID: " + player.PlayerState.MainCharaID)
	}
	pcs := []objects.PlayCharacterState{
		objects.NewPlayCharacterState(
			mchar,
			[]int64{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			[]int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		),
	}
	return NewQuickPostGameResultsResponse(base, ps, dci, ml, oml, pcs), nil
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
