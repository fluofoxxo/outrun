package responses

import (
	"log"
	"time"

	"github.com/fluofoxxo/outrun/objects"
	"github.com/fluofoxxo/outrun/playerdata"
)

type GameResultsResponseBase struct {
	BaseResponse
	PlayerState             pdata.PlayerState            `json:"playerState"`             // PlayerState
	ChaoState               []pdata.Chao                 `json:"chaoState"`               // The JSON can also apparently be "chaoStatus" without error // ChaoState
	DailyChallengeIncentive []objects.Incentive          `json:"dailyChallengeIncentive"` // TODO: Confirm this is the correct type  // DailyMissionIncentives
	CharacterState          []pdata.CharacterState       `json:"characterState"`          // CharacterState
	MessageList             []objects.Message            `json:"messageList"`             // MessageList
	TotalMessage            int64                        `json:"totalMessage"`
	OperatorMessageList     []objects.OperatorMessage    `json:"operatorMessageList"`
	TotalOperatorMessage    int64                        `json:"totalOperatorMessage"`
	PlayCharacterState      []objects.PlayCharacterState `json:"playCharacterState"` // PlayCharacterState
}

func NewGameResultsResponseBase(base BaseInfo, ps pdata.PlayerState, cs []pdata.Chao, dci []objects.Incentive, css []pdata.CharacterState, ml []objects.Message, oml []objects.OperatorMessage, pcs []objects.PlayCharacterState) GameResultsResponseBase {
	br := NewBaseResponse(base)
	return GameResultsResponseBase{
		br,
		ps,
		cs,
		dci,
		css,
		ml,
		int64(len(ml)),
		oml,
		int64(len(oml)),
		pcs,
	}
}

type QuickPostGameResultsResponse struct {
	GameResultsResponseBase
}

func NewQuickPostGameResultsResponse(base BaseInfo, ps pdata.PlayerState, cs []pdata.Chao, dci []objects.Incentive, css []pdata.CharacterState, ml []objects.Message, oml []objects.OperatorMessage, pcs []objects.PlayCharacterState) QuickPostGameResultsResponse {
	grrb := NewGameResultsResponseBase(base, ps, cs, dci, css, ml, oml, pcs)
	return QuickPostGameResultsResponse{
		grrb,
	}
}

func DefaultQuickPostGameResultsResponse(base BaseInfo, player pdata.Player) (QuickPostGameResultsResponse, error) {
	// TODO: const this!
	ps := player.PlayerState
	cs := player.ChaoState
	dci := []objects.Incentive{}
	css := player.CharacterStates
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
	return NewQuickPostGameResultsResponse(base, ps, cs, dci, css, ml, oml, pcs), nil
}

type PostGameResultsResponse struct {
	BaseResponse
	PlayerState             pdata.PlayerState            `json:"playerState"`
	CharacterState          []pdata.CharacterState       `json:"characterState"`
	ChaoState               []pdata.Chao                 `json:"chaoState"`
	PlayCharacterState      []objects.PlayCharacterState `json:"playCharacterState"`
	MileageMapState         pdata.MileageMapState        `json:"mileageMapState"`
	DailyChallengeIncentive []pdata.ItemInfo             `json:"dailyChallengeIncentive"`
	MileageIncentiveList    []objects.MileageIncentive   `json:"mileageIncentiveList"`
	MessageList             []objects.Message            `json:"messageList"`
	TotalMessage            int64                        `json:"totalMessage"`
	OperatorMessageList     []objects.OperatorMessage    `json:"operatorMessageList"`
	TotalOperatorMessage    int64                        `json:"totalOperatorMessage"`
	EventIncentiveList      []pdata.ItemInfo             `json:"eventIncentiveList"`
	WheelOptions            objects.WheelOptions         `json:"wheelOptions"`
}

func NewPostGameResultsResponse(base BaseInfo, ps pdata.PlayerState, cs []pdata.CharacterState, chs []pdata.Chao, pcs []objects.PlayCharacterState, mms pdata.MileageMapState, dci []pdata.ItemInfo, mil []objects.MileageIncentive, ml []objects.Message, oml []objects.OperatorMessage, eil []pdata.ItemInfo, wo objects.WheelOptions) PostGameResultsResponse {
	br := NewBaseResponse(base)
	return PostGameResultsResponse{
		br,
		ps,
		cs,
		chs,
		pcs,
		mms,
		dci,
		mil,
		ml,
		int64(len(ml)),
		oml,
		int64(len(oml)),
		eil,
		wo,
	}
}

func DefaultPostGameResultsResponse(base BaseInfo, player pdata.Player) PostGameResultsResponse {
	ps := player.PlayerState
	cs := player.CharacterStates
	chs := player.ChaoState
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
	mms := player.MileageMapState
	dci := pdata.DEFAULT_ITEMS // TODO: check if valid for game
	mil := []objects.MileageIncentive{objects.DefaultMileageIncentive()}
	ml := []objects.Message{}
	oml := []objects.OperatorMessage{}
	eil := []pdata.ItemInfo{}
	wo := objects.DefaultWheelOptions()
	return NewPostGameResultsResponse(
		base,
		ps,
		cs,
		chs,
		pcs,
		mms,
		dci,
		mil,
		ml,
		oml,
		eil,
		wo,
	)
}
