package obj

type BattleStatus struct {
	Wins            int64 `json:"numWin"`           // battle wins
	Losses          int64 `json:"numLose"`          // battle losses
	Draws           int64 `json:"numDraw"`          // battle draws
	LossesByDefault int64 `json:"numLoseByDefault"` // ?
	GoOnWins        int64 `json:"goOnWin"`          // ?
	GoOnLosses      int64 `json:"goOnLosses"`       // ?
}

func DefaultBattleStatus() BattleStatus {
	return BattleStatus{
		1,
		2,
		3,
		4,
		5,
		6,
	}
}
