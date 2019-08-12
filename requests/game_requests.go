package requests

import "github.com/fluofoxxo/outrun/objects"

type QuickActStartRequest struct {
	BasicRequest
	Modifier []int64 `json:"modifire"`           // Seems to be list of item IDs.
	Tutorial int64   `json:"tutorial,omitempty"` // will omit the field if not found
}

type ActStartRequest struct {
	QuickActStartRequest
	DistanceFriendList []objects.Friend `json:"distanceFriendList"` // TODO: Are we sure that it's not MileageFriend?
	Tutorial           int64            `json:"tutorial"`
}

type GameResultsRequestBase struct {
	BasicRequest
	Closed                 int64  `json:"closed"`
	Score                  int64  `json:"score,string"`
	NumRings               int64  `json:"numRings,string"`
	NumFailureRings        int64  `json:"numFailureRings,string"`
	NumRedStarRings        int64  `json:"numRedStarRings,string"`
	Distance               int64  `json:"distance,string"`
	DailyChallengeValue    int64  `json:"dailyChallengeValue,string"`
	DailyChallengeComplete int64  `json:"dailyChallengeComplete"`
	NumAnimals             int64  `json:"numAnimals,string"`
	CheatResult            string `json:"cheatResult"`
	MaxCombo               int64  `json:"maxCombo,string"`
}

type QuickPostGameResultsRequest struct {
	GameResultsRequestBase
}

type PostGameResultsRequest struct {
	GameResultsRequestBase
	BossDestroyed int64 `json:"bossDestroyed"`
	ChapterClear  int64 `json:"chapterClear"`
	GetChaoEgg    int64 `json:"getChaoEgg"`
	NumBossAttack int64 `json:"numBossAttack,string"`
	ReachPoint    int64 `json:"reachPoint,string"`
}

type MileageRewardRequest struct {
	BasicRequest
	Episode int64 `json:"episode,string"`
	Chapter int64 `json:"chapter,string"`
}
