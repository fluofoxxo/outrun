package responses

import (
	"strconv"

	"github.com/jinzhu/now"

	"github.com/RunnersRevival/outrun/consts"
	"github.com/RunnersRevival/outrun/logic"
	"github.com/RunnersRevival/outrun/netobj"
	"github.com/RunnersRevival/outrun/obj"
	"github.com/RunnersRevival/outrun/obj/constobjs"
	"github.com/RunnersRevival/outrun/responses/responseobjs"
)

type DailyChallengeDataResponse struct {
	BaseResponse
	IncentiveList          []obj.Incentive `json:"incentiveList"`
	IncentiveListCount     int64           `json:"incentiveListCont"`
	NumDailyChallengeCount int64           `json:"numDilayChalCont"`
	NumDailyChallengeDay   int64           `json:"numDailyChalDay"`
	MaxDailyChallengeDay   int64           `json:"maxDailyChalDay"`
	EndTime                int64           `json:"chalEndTime"`
}

func DailyChallengeData(base responseobjs.BaseInfo, numDailyChallenge, nextNumDailyChallenge int64) DailyChallengeDataResponse {
	ilSrc := consts.DailyMissionRewards
	//ilSrc := []int64{enums.ItemIDRing, enums.ItemIDBarrier, enums.ItemIDMagnet, enums.ItemIDTrampoline, enums.ItemIDAsteroid, enums.ItemIDDrill, enums.ItemIDRedRing} // must be length of seven!
	//ilSrc := []int64{900000, 900000, 900000, 900000, 900000, 900000, 900000}
	incentiveList := []obj.Incentive{}
	for amountSrc, id := range ilSrc {
		item := obj.NewItem(strconv.Itoa(int(id)), consts.DailyMissionRewardCounts[amountSrc])
		incentive := obj.NewIncentive(
			item,
			int64(amountSrc+1),
		)
		incentiveList = append(incentiveList, incentive)
	}
	incentiveListCount := int64(len(incentiveList))
	numDailyChallengeCount := int64(numDailyChallenge)
	maxDailyChallengeDay := int64(7)
	numDailyChallengeDay := int64(maxDailyChallengeDay - nextNumDailyChallenge)
	//endTime := int64(1470322800)      // 08/04/2016 @ 3:00PM (UTC)
	endTime := now.EndOfDay().UTC().Unix()
	baseResponse := NewBaseResponse(base)
	return DailyChallengeDataResponse{
		baseResponse,
		incentiveList,
		incentiveListCount,
		numDailyChallengeCount,
		numDailyChallengeDay,
		maxDailyChallengeDay,
		endTime,
	}
}

type CostListResponse struct {
	BaseResponse
	ConsumedCostList []obj.ConsumedItem `json:"consumedCostList"`
}

func CostList(base responseobjs.BaseInfo, consumedCostList []obj.ConsumedItem) CostListResponse {
	baseResponse := NewBaseResponse(base)
	out := CostListResponse{
		baseResponse,
		consumedCostList,
	}
	return out
}

func DefaultCostList(base responseobjs.BaseInfo) CostListResponse {
	return CostList(
		base,
		constobjs.DefaultCostList,
	)
}

type MileageDataResponse struct {
	BaseResponse
	MileageFriendList []netobj.MileageFriend `json:"mileageFriendList"`
	MileageMapState   netobj.MileageMapState `json:"mileageMapState"`
}

func MileageData(base responseobjs.BaseInfo, mileageFriendList []netobj.MileageFriend, mileageMapState netobj.MileageMapState) MileageDataResponse {
	baseResponse := NewBaseResponse(base)
	out := MileageDataResponse{
		baseResponse,
		mileageFriendList,
		mileageMapState,
	}
	return out
}

func DefaultMileageData(base responseobjs.BaseInfo, player netobj.Player) MileageDataResponse {
	mileageFriendList := player.MileageFriends
	mileageMapState := player.MileageMapState
	return MileageData(
		base,
		mileageFriendList,
		mileageMapState,
	)
}

type CampaignListResponse struct {
	BaseResponse
	CampaignList []obj.Campaign `json:"campaignList"`
}

func CampaignList(base responseobjs.BaseInfo, campaignList []obj.Campaign) CampaignListResponse {
	baseResponse := NewBaseResponse(base)
	out := CampaignListResponse{
		baseResponse,
		campaignList,
	}
	return out
}

func DefaultCampaignList(base responseobjs.BaseInfo) CampaignListResponse {
	campaignList := []obj.Campaign{}
	return CampaignList(
		base,
		campaignList,
	)
}

type ActStartBaseResponse struct {
	BaseResponse
	PlayerState  netobj.PlayerState `json:"playerState"`
	CampaignList []obj.Campaign     `json:"campaignList"`
}

func ActStartBase(base responseobjs.BaseInfo, playerState netobj.PlayerState, campaignList []obj.Campaign) ActStartBaseResponse {
	baseResponse := NewBaseResponse(base)
	return ActStartBaseResponse{
		baseResponse,
		playerState,
		campaignList,
	}
}

type QuickActStartResponse struct {
	ActStartBaseResponse
}

func QuickActStart(base responseobjs.BaseInfo, playerState netobj.PlayerState, campaignList []obj.Campaign) QuickActStartResponse {
	actStartBase := ActStartBase(base, playerState, campaignList)
	return QuickActStartResponse{
		actStartBase,
	}
}

func DefaultQuickActStart(base responseobjs.BaseInfo, player netobj.Player, campaignList []obj.Campaign) QuickActStartResponse {
	playerState := player.PlayerState
	return QuickActStart(
		base,
		playerState,
		campaignList,
	)
}

type ActStartResponse struct {
	ActStartBaseResponse
	DistanceFriendList []netobj.MileageFriend `json:"distanceFriendList"` // TODO: Discover if correct type
}

func ActStart(base responseobjs.BaseInfo, playerState netobj.PlayerState, campaignList []obj.Campaign, distFriends []netobj.MileageFriend) ActStartResponse {
	actStartBase := ActStartBase(base, playerState, campaignList)
	return ActStartResponse{
		actStartBase,
		distFriends,
	}
}

func DefaultActStart(base responseobjs.BaseInfo, player netobj.Player, campaignList []obj.Campaign) ActStartResponse {
	playerState := player.PlayerState
	distFriends := []netobj.MileageFriend{}
	return ActStart(
		base,
		playerState,
		campaignList,
		distFriends,
	)
}

type QuickPostGameResultsResponse struct {
	BaseResponse
	PlayerState             netobj.PlayerState    `json:"playerState"`
	ChaoState               []netobj.Chao         `json:"chaoState"`
	DailyChallengeIncentive []obj.Incentive       `json:"dailyChallengeIncentive"` // should be obj.Item, but game doesn't care
	MessageList             []obj.Message         `json:"messageList"`
	OperatorMessageList     []obj.OperatorMessage `json:"operatorMessageList"`
	TotalMessage            int64                 `json:"totalMessage"`
	TotalOperatorMessage    int64                 `json:"totalOperatorMessage"`
	PlayCharacterState      []netobj.Character    `json:"playCharacterState"` // Character can substitute PlayCharacter
}

func QuickPostGameResults(base responseobjs.BaseInfo, player netobj.Player, dci []obj.Incentive, ml []obj.Message, oml []obj.OperatorMessage, pcs []netobj.Character) QuickPostGameResultsResponse {
	baseResponse := NewBaseResponse(base)
	playerState := player.PlayerState
	chaoState := player.ChaoState
	dailyChallengeIncentive := dci
	messageList := []obj.Message{}
	operatorMessageList := []obj.OperatorMessage{}
	totalMessage := int64(len(messageList))
	totalOperatorMessage := int64(len(operatorMessageList))
	playCharacterState := pcs
	return QuickPostGameResultsResponse{
		baseResponse,
		playerState,
		chaoState,
		dailyChallengeIncentive,
		messageList,
		operatorMessageList,
		totalMessage,
		totalOperatorMessage,
		playCharacterState,
	}
}

func DefaultQuickPostGameResults(base responseobjs.BaseInfo, player netobj.Player, pcs []netobj.Character) QuickPostGameResultsResponse {
	dci := []obj.Incentive{}
	ml := []obj.Message{}
	oml := []obj.OperatorMessage{
		obj.DefaultOperatorMessage(),
	}
	/*
		mainC, err := player.GetMainChara()
		if err != nil {
			// TODO: use better error handling!
			log.Println("[ERR] (DefaultQuickPostGameResults) Error getting main character: ", err)
		}
		subC, err := player.GetSubChara()
		if err != nil {
			// TODO: use better error handling!
			log.Println("[ERR] (DefaultQuickPostGameResults) Error getting sub character: ", err)
		}
		pcs := []netobj.Character{
			mainC,
			subC,
		}
	*/
	return QuickPostGameResults(
		base,
		player,
		dci,
		ml,
		oml,
		pcs,
	)
}

type PostGameResultsResponse struct {
	QuickPostGameResultsResponse
	MileageMapState      netobj.MileageMapState `json:"mileageMapState"`
	MileageIncentiveList []obj.MileageIncentive `json:"mileageIncentiveList"`
	WheelOptions         netobj.WheelOptions    `json:"wheelOptions"`
}

func PostGameResults(base responseobjs.BaseInfo, player netobj.Player, dci []obj.Incentive, ml []obj.Message, oml []obj.OperatorMessage, pcs []netobj.Character, mms netobj.MileageMapState, mil []obj.MileageIncentive, wo netobj.WheelOptions) PostGameResultsResponse {
	baseResponse := NewBaseResponse(base)
	playerState := player.PlayerState
	chaoState := player.ChaoState
	dailyChallengeIncentive := dci
	messageList := []obj.Message{}
	operatorMessageList := []obj.OperatorMessage{}
	totalMessage := int64(len(messageList))
	totalOperatorMessage := int64(len(operatorMessageList))
	playCharacterState := pcs
	qpgrr := QuickPostGameResultsResponse{
		baseResponse,
		playerState,
		chaoState,
		dailyChallengeIncentive,
		messageList,
		operatorMessageList,
		totalMessage,
		totalOperatorMessage,
		playCharacterState,
	}
	return PostGameResultsResponse{
		qpgrr,
		mms,
		mil,
		wo,
	}
}

func DefaultPostGameResults(base responseobjs.BaseInfo, player netobj.Player, pcs []netobj.Character, incentives []obj.MileageIncentive) PostGameResultsResponse {
	qpgrr := DefaultQuickPostGameResults(base, player, pcs)
	mms := player.MileageMapState
	//mil := []obj.MileageIncentive{}
	//wo := netobj.DefaultWheelOptions(player.PlayerState.NumRouletteTicket, player.RouletteInfo.RouletteCountInPeriod)
	// TODO: Remove logic from response!!
	player.LastWheelOptions = logic.WheelRefreshLogic(player, player.LastWheelOptions)
	wo := player.LastWheelOptions
	return PostGameResultsResponse{
		qpgrr,
		mms,
		incentives,
		wo,
	}
}

type PostGameResultsEventResponse struct {
	PostGameResultsResponse
	EventState         netobj.EventState `json:"eventState"`
	EventIncentiveList []obj.Item        `json:"eventIncentiveList"`
}

func PostGameResultsEvent(base responseobjs.BaseInfo, player netobj.Player, dci []obj.Incentive, ml []obj.Message, oml []obj.OperatorMessage, pcs []netobj.Character, mms netobj.MileageMapState, mil []obj.MileageIncentive, wo netobj.WheelOptions, eil []obj.Item, es netobj.EventState) PostGameResultsEventResponse {
	baseResponse := NewBaseResponse(base)
	playerState := player.PlayerState
	chaoState := player.ChaoState
	dailyChallengeIncentive := dci
	messageList := []obj.Message{}
	operatorMessageList := []obj.OperatorMessage{}
	totalMessage := int64(len(messageList))
	totalOperatorMessage := int64(len(operatorMessageList))
	playCharacterState := pcs
	qpgrr := QuickPostGameResultsResponse{
		baseResponse,
		playerState,
		chaoState,
		dailyChallengeIncentive,
		messageList,
		operatorMessageList,
		totalMessage,
		totalOperatorMessage,
		playCharacterState,
	}
	pgrr := PostGameResultsResponse{
		qpgrr,
		mms,
		mil,
		wo,
	}
	return PostGameResultsEventResponse{
		pgrr,
		es,
		eil,
	}
}

func DefaultPostGameResultsEvent(base responseobjs.BaseInfo, player netobj.Player, pcs []netobj.Character, incentives []obj.MileageIncentive, eventIncentives []obj.Item, eventState netobj.EventState) PostGameResultsEventResponse {
	qpgrr := DefaultQuickPostGameResults(base, player, pcs)
	mms := player.MileageMapState
	//mil := []obj.MileageIncentive{}
	//wo := netobj.DefaultWheelOptions(player.PlayerState.NumRouletteTicket, player.RouletteInfo.RouletteCountInPeriod)
	// TODO: Remove logic from response!!
	player.LastWheelOptions = logic.WheelRefreshLogic(player, player.LastWheelOptions)
	wo := player.LastWheelOptions
	pgrr := PostGameResultsResponse{
		qpgrr,
		mms,
		incentives,
		wo,
	}
	return PostGameResultsEventResponse{
		pgrr,
		eventState,
		eventIncentives,
	}
}

type FreeItemListResponse struct {
	BaseResponse
	FreeItemList []obj.Item `json:"freeItemList"`
}

func FreeItemList(base responseobjs.BaseInfo, freeItemList []obj.Item) FreeItemListResponse {
	baseResponse := NewBaseResponse(base)
	return FreeItemListResponse{
		baseResponse,
		freeItemList,
	}
}

func DefaultFreeItemList(base responseobjs.BaseInfo) FreeItemListResponse {
	freeItemList := constobjs.AllItems
	return FreeItemList(
		base,
		freeItemList,
	)
}

type MileageRewardResponse struct {
	BaseResponse
	MileageRewards []obj.MileageReward `json:"mileageMapRewardList"`
}

func MileageReward(base responseobjs.BaseInfo, mileageRewards []obj.MileageReward) MileageRewardResponse {
	baseResponse := NewBaseResponse(base)
	return MileageRewardResponse{
		baseResponse,
		mileageRewards,
	}
}

func DefaultMileageReward(base responseobjs.BaseInfo, chapter, episode int64) MileageRewardResponse {
	return MileageReward(
		base,
		constobjs.GetAreaReward(chapter, episode),
	)
}
