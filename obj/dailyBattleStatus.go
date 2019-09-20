package obj

type DailyBattleStatus struct {
    Wins            int64 `json:"numWin"`
    Losses          int64 `json:"numLose"`
    Draws           int64 `json:"numDraw"`
    LossesByDefault int64 `json:"numLoseByDefault"`
    GoOnWins        int64 `json:"goOnWin"`  // ???
    GoOnLosses      int64 `json:"goOnLose"` // ???
}

func DefaultDailyBattleStatus() DailyBattleStatus {
    return DailyBattleStatus{
        54,
        56,
        57,
        58,
        59,
        63,
    }
}
