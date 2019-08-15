package obj

const (
	DefaultMapDistance   = 0
	DefaultNumBossAttack = 1
	DefaultStageDistance = 1337000
	DefaultStageMaxScore = 8008135
)

type MapInfo struct {
	MapDistance   int64 `json:"mapDistance"`   // used sparingly in game...?
	NumBossAttack int64 `json:"numBossAttack"` // TODO: number of boss fights? Check how often this is used in game
	StageDistance int64 `json:"stageDistance"` // TODO: discover use
	StageMaxScore int64 `json:"stageMaxScore"` // TODO: discover use
}

func DefaultMapInfo() MapInfo {
	return MapInfo{
		DefaultMapDistance,
		DefaultNumBossAttack,
		DefaultStageDistance,
		DefaultStageMaxScore,
	}
}
