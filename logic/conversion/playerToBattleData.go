package conversion

import (
	"time"

	"github.com/fluofoxxo/outrun/enums"
	"github.com/fluofoxxo/outrun/netobj"
	"github.com/fluofoxxo/outrun/obj"
)

func DebugPlayerToBattleData(player netobj.Player) obj.BattleData {
	uid := player.ID
	username := player.Username
	maxScore := player.PlayerState.HighScore
	league := int64(enums.RankingLeagueC)
	loginTime := time.Now().UTC().Unix() - 60
	mainChaoID := player.PlayerState.MainChaoID
	mainChaoLevel := int64(2) // TODO: this may be problematic if the game does checks
	subChaoID := player.PlayerState.SubChaoID
	subChaoLevel := int64(3) // TODO: this may be problematic if the game does checks
	rank := int64(1)         // ??
	mainCharaID := player.PlayerState.MainCharaID
	mainCharaLevel := int64(4) // TODO: this may be problematic if the game does checks
	subCharaID := player.PlayerState.SubCharaID
	subCharaLevel := int64(5) // TODO: this may be problematic if the game does checks
	goOnWin := int64(0)
	isSentEnergy := int64(0)
	language := int64(enums.LangEnglish)
	return obj.BattleData{
		uid,
		username,
		maxScore,
		league,
		loginTime,
		mainChaoID,
		mainChaoLevel,
		subChaoID,
		subChaoLevel,
		rank,
		mainCharaID,
		mainCharaLevel,
		subCharaID,
		subCharaLevel,
		goOnWin,
		isSentEnergy,
		language,
	}
}
