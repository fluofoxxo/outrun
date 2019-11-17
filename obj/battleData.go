package obj

import (
	"time"

	"github.com/fluofoxxo/outrun/enums"
)

type BattleData struct {
	UserID         string `json:"userId"`     // user ID
	Name           string `json:"name"`       // username?
	MaxScore       int64  `json:"maxScore"`   // highest score gained
	League         int64  `json:"league"`     // enums.RankingLeague*
	LoginTime      int64  `json:"loginTime"`  // used by game in ranking window with GetCurrentTime() - loginTime
	MainChaoID     string `json:"mainChaoId"` // TODO: if this struct is a problem for the game, the game actually only expects the first two strings to be real strings; these are ints
	MainChaoLevel  int64  `json:"mainChaoLevel"`
	SubChaoID      string `json:"subChaoId"`
	SubChaoLevel   int64  `json:"subChaoLevel"`
	Rank           int64  `json:"numRank"` // left justified by '0's in the game; unknown what it's sourced from
	MainCharaID    string `json:"charaId"`
	MainCharaLevel int64  `json:"charaLevel"`
	SubCharaID     string `json:"subCharaId"`
	SubCharaLevel  int64  `json:"subCharaLevel"`
	GoOnWin        int64  `json:"goOnWin"`   // I think this is either win count or successive win count?
	IsSentEnergy   int64  `json:"energyFlg"` // you can't send energy to friend if true (is only friend if social connection); enables custom icon
	Language       int64  `json:"language"`  // enums.Lang*
}

func DebugRivalBattleData() BattleData {
	return BattleData{
		"8118118111",
		"Adamska",
		1020304,
		enums.RankingLeagueS,
		time.Now().UTC().Unix() - 60000,
		"400009",
		4,
		"400000",
		9,
		2,
		enums.CTStrShadow, // TODO: test if this will take event characters...?
		12,
		enums.CTStrTails,
		34,
		79,
		0,
		enums.LangEnglish,
	}
}
