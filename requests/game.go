package requests

type QuickPostGameResultsRequest struct {
	Base
	Score                  int64  `json:"score,string"`
	Rings                  int64  `json:"numRings,string"`
	FailureRings           int64  `json:"numFailureRings,string"`
	RedRings               int64  `json:"numRedStarRings,string"`
	Distance               int64  `json:"distance,string"`
	DailyChallengeValue    int64  `json:"dailyChallengeValue,string"`
	DailyChallengeComplete int64  `json:"dailyChallengeComplete"`
	Animals                int64  `json:"numAnimals,string"`
	MaxCombo               int64  `json:"maxCombo,string"`
	Closed                 int64  `json:"closed"`
	CheatResult            string `json:"cheatResult"`
}

type PostGameResultsRequest struct {
	QuickPostGameResultsRequest
	BossDestroyed int64 `json:"bossDestroyed"`
	ChapterClear  int64 `json:"chapterClear"`
	GetChaoEgg    int64 `json:"getChaoEgg"`
	NumBossAttack int64 `json:"numBossAttack,string"`
	ReachPoint    int64 `json:"reachPoint,string"`
}
