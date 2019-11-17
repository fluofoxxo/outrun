package conversion

import (
	"time"

	"github.com/fluofoxxo/outrun/enums"
	"github.com/fluofoxxo/outrun/netobj"
	"github.com/fluofoxxo/outrun/obj"
)

func PlayerToLeaderboardEntry(player netobj.Player) obj.LeaderboardEntry {
	friendID := player.ID // TODO: is this right?
	name := player.Username
	url := player.Username + "_findme" // TODO: only used for testing right now
	grade := player.PlayerState.Rank   // TODO: _probably_ what this is, but unsure
	exposeOnline := int64(0)
	rankingScore := player.PlayerState.HighScore // TODO: this probably differs based on mode...
	rankChanged := int64(2)
	isSentEnergy := int64(0)
	expireTime := time.Now().UTC().Unix() + 12345
	numRank := player.PlayerState.Rank + 1
	loginTime := player.LastLogin
	mainCharaID := player.PlayerState.MainCharaID
	mainCharaLevel := int64(12) // TODO: remove testing
	subCharaID := player.PlayerState.SubCharaID
	subCharaLevel := int64(34) // TODO: remove testing
	mainChaoID := player.PlayerState.MainChaoID
	mainChaoLevel := int64(5) // TODO: remove testing
	subChaoID := player.PlayerState.SubChaoID
	subChaoLevel := int64(6) // TODO: remove testing
	language := int64(enums.LangEnglish)
	league := player.PlayerState.Rank // TODO: check if this is right
	maxScore := player.PlayerState.HighScore
	return obj.LeaderboardEntry{
		friendID,
		name,
		url,
		grade,
		exposeOnline,
		rankingScore,
		rankChanged,
		isSentEnergy,
		expireTime,
		numRank,
		loginTime,
		mainCharaID,
		mainCharaLevel,
		subCharaID,
		subCharaLevel,
		mainChaoID,
		mainChaoLevel,
		subChaoID,
		subChaoLevel,
		language,
		league,
		maxScore,
	}
}
