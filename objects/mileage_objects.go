package objects

type MileageFriend struct { // still don't know what fields this has
}

type MileageMapState struct {
	Episode          int64 `json:"episode"`
	Chapter          int64 `json:"chapter"`
	Point            int64 `json:"point"`
	MapDistance      int64 `json:"mapDistance"`
	NumBossAttack    int64 `json:"numBossAttack"`
	StageDistance    int64 `json:"stageDistance"`
	StageTotalScore  int64 `json:"stageTotalScore"`
	StageMaxScore    int64 `json:"stageMaxScore"`
	ChapterStartTime int64 `json:"chapterStartTime"`
}
