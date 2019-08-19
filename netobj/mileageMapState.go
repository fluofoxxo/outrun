package netobj

import (
    "time"

    "github.com/fluofoxxo/outrun/obj"
)

type MileageMapState struct {
    obj.MapInfo
    Episode          int64 `json:"episode"`
    Chapter          int64 `json:"chapter"`
    Point            int64 `json:"point"`            // point in episode
    StageTotalScore  int64 `json:"stageTotalScore"`  // TODO: discover use. This is very likely used for the total score gained in the current chapter, which means this value MUST be set to the total chapter score
    ChapterStartTime int64 `json:"chapterStartTime"` // TODO: discover use. Appears to be used for point item expiry?
}

func (m MileageMapState) AdvanceChapter() {
    m.StageTotalScore = int64(0)
    m.ChapterStartTime = time.Now().Unix()
    m.Chapter += 1
    m.Episode = int64(0)
    m.Point = int64(0)
}
func (m MileageMapState) AddScore(score int64) {
    m.StageTotalScore += score
}

func DefaultMileageMapState() MileageMapState {
    // implies that this is a new account
    mapInfo := obj.DefaultMapInfo()
    return MileageMapState{
        mapInfo,
        1,
        1,
        0,
        0,
        time.Now().Unix(),
    }
}
