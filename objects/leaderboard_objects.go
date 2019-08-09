package objects

import (
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
