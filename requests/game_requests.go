package requests

type QuickActStartRequest struct {
	BasicRequest
	Modifier []int64 `json:"modifire"`           // Seems to be list of item IDs.
	Tutorial int64   `json:"tutorial,omitempty"` // will omit the field if not found
}

type QuickPostGameResultsRequest struct {
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
