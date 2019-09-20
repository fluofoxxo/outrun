package obj

import (
    "time"

    "github.com/fluofoxxo/outrun/enums"
)

type DailyBattleData struct {
    UserID         string `json:"userId"`
    Name           string `json:"name"`
    MainChaoID     string `json:"mainChaoId"`
    SubChaoID      string `json:"subChaoId"`
    MainCharaID    string `json:"charaId"`
    SubCharaID     string `json:"subCharaId"`
    MainChaoLevel  int64  `json:"mainChaoLevel"`
    SubChaoLevel   int64  `json:"subChaoLevel"`
    MainCharaLevel int64  `json:"charaLevel"`
    SubCharaLevel  int64  `json:"subCharaLevel"`
    Rank           int64  `json:"numRank"`
    HighScore      int64  `json:"maxScore"`
    League         int64  `json:"league"`
    LoginTime      int64  `json:"loginTime"`
    IsSentEnergy   int64  `json:"energyFlg"` // flag
    GoOnWins       int64  `json:"goOnWin"`   // ??? RankingUtil.rankIndex, but what does this mean?
    Language       int64  `json:"language"`  // enums.Language*
}

func DefaultDailyBattleData() DailyBattleData {
    return DailyBattleData{
        "1230001234",
        "La-li-lu-le-lo",
        "400000",
        "400001",
        enums.CTStrSonic,
        enums.CTStrTails,
        2,
        3,
        4,
        5,
        //371,
        1,
        1731731,
        enums.RankingLeagueB,
        time.Now().Unix() - 555, // a little less than 9 minutes prior
        0,
        //7,
        1,
        enums.LanguageEnglish,
    }
}
