package obj

import (
	"strconv"
	"time"
)

type LeaderboardEntry struct {
	FriendID          string `json:"friendId"`
	Name              string `json:"name"`
	URL               string `json:"url"`
	Grade             int64  `json:"grade"`
	ExposeOnline      int64  `json:"exposeOnline"` // TODO: remove this, apparently it's not used
	RankingScore      int64  `json:"rankingScore"`
	RankChanged       int64  `json:"rankChanged"` // TODO: potentially remove this, doesn't seem to be used
	IsSentEnergy      int64  `json:"energyFlg"`
	ExpireTime        int64  `json:"expireTime"` // TODO: doesn't seem to be used
	NumRank           int64  `json:"numRank"`    // left adjusted with '0's to match length 3; unknown of use, called internally 'mapRank'
	LoginTime         int64  `json:"loginTime"`
	CharacterID       string `json:"charaId"`
	CharacterLevel    int64  `json:"characterLevel"`
	SubcharacterID    string `json:"subCharaId"`
	SubcharacterLevel int64  `json:"subCharaLevel"`
	MainChaoID        string `json:"mainChaoId"`
	MainChaoLevel     int64  `json:"mainChaoLevel"`
	SubChaoID         string `json:"subChaoId"`
	SubChaoLevel      int64  `json:"subChaoLevel"`
	Language          int64  `json:"language"` // enums.Lang*
	League            int64  `json:"league"`
	MaxScore          int64  `json:"maxScore"`
}

func NewLeaderboardEntry(fid, n, url string, g, eo, rs, rc, ise, et, nr, lt, cid, cl, schid, schl, mcid, mcl, scid, scl, lang, league, maxScore int64) LeaderboardEntry {
	return LeaderboardEntry{
		fid,
		n,
		url,
		g,
		eo,
		rs,
		rc,
		ise,
		et,
		nr,
		lt,
		strconv.Itoa(int(cid)),
		cl,
		strconv.Itoa(int(schid)),
		schl,
		strconv.Itoa(int(mcid)),
		mcl,
		strconv.Itoa(int(scid)),
		scl,
		lang,
		league,
		maxScore,
	}
}

func DefaultLeaderboardEntry(uid string) LeaderboardEntry {
	return NewLeaderboardEntry(
		uid,
		"",
		"",
		0,
		1,
		0,
		0,
		0,
		0,
		0,
		time.Now().Unix(), // this should be player.LastLogin!
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
	)
}
