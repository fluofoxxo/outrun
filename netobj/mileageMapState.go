package netobj

import (
	"github.com/fluofoxxo/outrun/obj"
)

type MileageMapState struct {
	obj.MapInfo
	Episode          int64 `json:"episode"`
	Chapter          int64 `json:"chapter"`
	Point            int64 `json:"point"`            // point in episode
	StageTotalScore  int64 `json:"stageTotalScore"`  // TODO: discover use. This is likely used for the total score gained in the current chapter.
	ChapterStartTime int64 `json:"chapterStartTime"` // TODO: discover use. _Maybe_ was used to record when the chapter was started
}
