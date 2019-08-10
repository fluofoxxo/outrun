package responses

import (
	"log"
	"time"

	"github.com/fluofoxxo/outrun/objects"
	"github.com/fluofoxxo/outrun/playerdata"
)

type QuickActStartResponse struct {
	BaseResponse
	PlayerState  pdata.PlayerState  `json:"playerState"`
	CampaignList []objects.Campaign `json:"campaignList"`
}

func NewQuickActStartResponse(base BaseInfo, playerState pdata.PlayerState, campaignList []objects.Campaign) QuickActStartResponse {
	br := NewBaseResponse(base)
	qasr := QuickActStartResponse{
		br,
		playerState,
		campaignList,
	}
	return qasr
}

func DefaultQuickActStartResponse(base BaseInfo, playerState pdata.PlayerState) QuickActStartResponse {
	return NewQuickActStartResponse(base, playerState, []objects.Campaign{})
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
