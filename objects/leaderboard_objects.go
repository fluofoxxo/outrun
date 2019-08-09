package objects

import (
    "time"

    "github.com/fluofoxxo/outrun/playerdata"
)

type LeagueData struct {
    LeagueID           int64           `json:"leagueId,string"`
    GroupID            int64           `json:"groupId,string"`
    NumUp              int64           `json:"numUp,string"`
    NumDown            int64           `json:"numDown,string"`
    NumGroupMember     int64           `json:"numGroupMember,string"`
    HighScoreOperator  []OperatorScore `json:"highScoreOpe"`  // what is this used for?
    TotalScoreOperator []OperatorScore `json:"totalScoreOpe"` // ^^^^^^^^^^^^^^^^^^^^

}

type OperatorScore struct { // TODO: no documentation on what this is
    Operator    int64            `json:"operator,string"`
    Number      int64            `json:"number,string"`
    PresentList []pdata.ItemInfo `json:"presentList,string"`
}

type PlayerEntry struct {
    FriendID          string `json:"friendId"`
    Name              string `json:"name"`
    URL               string `json:"url"`
    Grade             int64  `json:"grade"`
    ExposeOnline      int64  `json:"exposeOnline"`
    RankingScore      int64  `json:"rankingScore"`
    RankChanged       int64  `json:"rankChanged"`
    IsSentEnergy      int64  `json:"energyFlg"`
    ExpireTime        int64  `json:"expireTime"`
    NumRank           int64  `json:"numRank"`
    LoginTime         int64  `json:"loginTime"`
    CharacterID       int64  `json:"charaId"`
    CharacterLevel    int64  `json:"characterLevel"`
    SubcharacterID    int64  `json:"subCharaId"`
    SubcharacterLevel int64  `json:"subCharaLevel"`
    MainChaoID        int64  `json:"mainChaoId"`
    MainChaoLevel     int64  `json:"mainChaoLevel"`
    SubChaoID         int64  `json:"subChaoId"`
    SubChaoLevel      int64  `json:"subChaoLevel"`
    Language          int64  `json:"language"`
    League            int64  `json:"league"`
    MaxScore          int64  `json:"maxScore"`
}

func DefaultPlayerEntry(uid string) PlayerEntry {
    // TODO: const the below
    return PlayerEntry{
        uid, // PlayerEntry.FriendID should be set to the current player's ID
        "",
        "",
        0,
        1,
        0,
        0,
        0,
        0,
        0,
        time.Now().Unix(), // TODO: this should be the player's login time for real!
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        1,
        0,
        0,
    }
}
