package conversion

import (
	"time"

	"github.com/fluofoxxo/outrun/enums"
	"github.com/fluofoxxo/outrun/netobj"
	"github.com/fluofoxxo/outrun/obj"
)

func PlayerToLeaderboardEntry(player netobj.Player, mode int64) obj.LeaderboardEntry {
	friendID := player.ID
	name := player.Username
	url := player.Username + "_findme" // TODO: only used for testing right now
	grade := int64(1) // TODO: make this specifiable in the future (this is the high score ranking it seems)
	exposeOnline := int64(0)
	rankingScore := player.PlayerState.HighScore
	rankChanged := int64(0)
	isSentEnergy := int64(0)
	expireTime := now.EndOfWeek().UTC().Unix()
	numRank := player.PlayerState.Rank
	loginTime := player.LastLogin
	mainCharaID := player.PlayerState.MainCharaID
	mainCharaLevel := int64(0)
	subCharaID := player.PlayerState.SubCharaID
	subCharaLevel := int64(0)
	mainChaoID := player.PlayerState.MainChaoID
	mainChaoLevel := int64(0)
	subChaoID := player.PlayerState.SubChaoID
	subChaoLevel := int64(0)
	if player.IndexOfChara(mainCharaID) != -1 {
		mainCharaLevel = player.CharacterState[player.IndexOfChara(mainCharaID)].Level
	}
	if player.IndexOfChara(subCharaID) != -1 {
		subCharaLevel = player.CharacterState[player.IndexOfChara(subCharaID)].Level
	}
	if player.IndexOfChao(mainChaoID) != -1 {
		mainChaoLevel = player.ChaoState[player.IndexOfChao(mainChaoID)].Level
	}
	if player.IndexOfChao(subChaoID) != -1 {
		subChaoLevel = player.ChaoState[player.IndexOfChao(subChaoID)].Level
	}
	language := int64(enums.LangEnglish)
	league := player.PlayerState.RankingLeague
	maxScore := player.PlayerState.HighScore
	if mode == 1 { //timed mode?
		rankingScore = player.PlayerState.TimedHighScore
		league = player.PlayerState.QuickRankingLeague
		maxScore = player.PlayerState.TimedHighScore
	}
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
