package muxhandlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/fluofoxxo/outrun/analytics"
	"github.com/fluofoxxo/outrun/analytics/factors"
	"github.com/fluofoxxo/outrun/config"
	"github.com/fluofoxxo/outrun/consts"
	"github.com/fluofoxxo/outrun/db"
	"github.com/fluofoxxo/outrun/emess"
	"github.com/fluofoxxo/outrun/enums"
	"github.com/fluofoxxo/outrun/helper"
	"github.com/fluofoxxo/outrun/logic/campaign"
	"github.com/fluofoxxo/outrun/netobj"
	"github.com/fluofoxxo/outrun/obj/constobjs"
	"github.com/fluofoxxo/outrun/requests"
	"github.com/fluofoxxo/outrun/responses"
	"github.com/fluofoxxo/outrun/status"
)

func GetDailyChallengeData(helper *helper.Helper) {
	// no player, agnostic
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DailyChallengeData(baseInfo)
	err := helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetCostList(helper *helper.Helper) {
	// no player, agonstic
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultCostList(baseInfo)
	err := helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetMileageData(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting player", err) // TODO: see if InternalErr is consistent with other usage of this context
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultMileageData(baseInfo, player)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetCampaignList(helper *helper.Helper) {
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultCampaignList(baseInfo)
	err := helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func QuickActStart(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultQuickActStart(baseInfo, player)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
	_, err = analytics.Store(player.ID, factors.AnalyticTypeTimedStarts)
	if err != nil {
		helper.WarnErr("Error storing analytics (AnalyticTypeTimedStarts)", err)
	}
}

func ActStart(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultActStart(baseInfo, player)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
	_, err = analytics.Store(player.ID, factors.AnalyticTypeStoryStarts)
	if err != nil {
		helper.WarnErr("Error storing analytics (AnalyticTypeStoryStarts)", err)
	}
}

func ActRetry(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	player.PlayerState.NumRedRings -= 5
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.NewBaseResponse(baseInfo)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
	_, err = analytics.Store(player.ID, factors.AnalyticTypeRevives)
	if err != nil {
		helper.WarnErr("Error storing analytics (AnalyticTypeRevives)", err)
	}
}

func QuickPostGameResults(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.QuickPostGameResultsRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}

	mainC, err := player.GetMainChara()
	if err != nil {
		helper.InternalErr("Error getting main character", err)
		return
	}
	subC, err := player.GetSubChara()
	if err != nil {
		helper.InternalErr("Error getting sub character", err)
		return
	}
	mainCIndex := player.IndexOfChara(mainC.ID) // TODO: check if -1
	subCIndex := player.IndexOfChara(subC.ID)   // TODO: check if -1
	playCharacters := []netobj.Character{
		mainC,
		subC,
	}
	if request.Closed == 0 { // If the game wasn't exited out of
		player.PlayerState.NumRings += request.Rings
		player.PlayerState.NumRedRings += request.RedRings
		player.PlayerState.NumRouletteTicket += request.RedRings // TODO: URGENT! Remove as soon as possible!
		player.PlayerState.Animals += request.Animals
		playerTimedHighScore := player.PlayerState.TimedHighScore
		if request.Score > playerTimedHighScore {
			player.PlayerState.TimedHighScore = request.Score
		}
		//player.PlayerState.TotalDistance += request.Distance  // We don't do this in timed mode!
		// increase character(s)'s experience
		expIncrease := request.Rings + request.FailureRings // all rings collected
		abilityIndex := 1
		for abilityIndex == 1 { // unused ability is at index 1
			abilityIndex = rand.Intn(len(mainC.AbilityLevel))
		}
		// check that increases exist
		_, ok := consts.UpgradeIncreases[mainC.ID]
		if !ok {
			helper.InternalErr("Error getting upgrade increase", fmt.Errorf("no key '%s' in consts.UpgradeIncreases", mainC.ID))
			return
		}
		_, ok = consts.UpgradeIncreases[subC.ID]
		if !ok {
			helper.InternalErr("Error getting upgrade increase", fmt.Errorf("no key '%s' in consts.UpgradeIncreases", subC.ID))
			return
		}
		if playCharacters[0].Level < 100 {
			playCharacters[0].Exp += expIncrease
			for playCharacters[0].Exp >= playCharacters[0].Cost {
				// more exp than cost = level up
				playCharacters[0].Level++                                               // increase level
				playCharacters[0].AbilityLevel[abilityIndex]++                          // increase ability level
				playCharacters[0].Exp -= playCharacters[0].Cost                         // remove cost from exp
				playCharacters[0].Cost += consts.UpgradeIncreases[playCharacters[0].ID] // increase cost
			}
		}
		if playCharacters[1].Level < 100 {
			playCharacters[1].Exp += expIncrease
			for playCharacters[1].Exp >= playCharacters[1].Cost {
				// more exp than cost = level up
				playCharacters[1].Level++                                               // increase level
				playCharacters[1].AbilityLevel[abilityIndex]++                          // increase ability level
				playCharacters[1].Exp -= playCharacters[1].Cost                         // remove cost from exp
				playCharacters[1].Cost += consts.UpgradeIncreases[playCharacters[1].ID] // increase cost
			}
		}

		helper.DebugOut("Old mainC Exp: " + strconv.Itoa(int(mainC.Exp)) + " / " + strconv.Itoa(int(mainC.Cost)))
		helper.DebugOut("Old mainC Level: " + strconv.Itoa(int(mainC.Level)))
		helper.DebugOut("Old subC Exp: " + strconv.Itoa(int(subC.Exp)) + " / " + strconv.Itoa(int(subC.Cost)))
		helper.DebugOut("Old subC Level: " + strconv.Itoa(int(subC.Level)))
		helper.DebugOut("New mainC Exp: " + strconv.Itoa(int(playCharacters[0].Exp)) + " / " + strconv.Itoa(int(playCharacters[0].Cost)))
		helper.DebugOut("New mainC Level: " + strconv.Itoa(int(playCharacters[0].Level)))
		helper.DebugOut("New subC Exp: " + strconv.Itoa(int(playCharacters[1].Exp)) + " / " + strconv.Itoa(int(playCharacters[1].Cost)))
		helper.DebugOut("New subC Level: " + strconv.Itoa(int(playCharacters[1].Level)))
	}

	/*
		if err != nil {
			helper.InternalErr("Error saving player", err)
			return
		}
	*/

	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultQuickPostGameResults(baseInfo, player, playCharacters)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}

	// apply the save after the response so that we don't break the leveling
	mainC = playCharacters[0]
	subC = playCharacters[1]
	player.CharacterState[mainCIndex] = mainC
	player.CharacterState[subCIndex] = subC
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}
	_, err = analytics.Store(player.ID, factors.AnalyticTypeTimedEnds)
	if err != nil {
		helper.WarnErr("Error storing analytics (AnalyticTypeTimedEnds)", err)
	}
}

func PostGameResults(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.PostGameResultsRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}

	mainC, err := player.GetMainChara()
	if err != nil {
		helper.InternalErr("Error getting main character", err)
		return
	}
	subC, err := player.GetSubChara()
	if err != nil {
		helper.InternalErr("Error getting sub character", err)
		return
	}
	playCharacters := []netobj.Character{
		mainC,
		subC,
	}
	helper.DebugOut("Pre-function")
	helper.DebugOut("Chapter: %v", player.MileageMapState.Chapter)
	helper.DebugOut("Episode: %v", player.MileageMapState.Episode)
	helper.DebugOut("StageTotalScore: %v", player.MileageMapState.StageTotalScore)
	helper.DebugOut("Point: %v", player.MileageMapState.Point)
	helper.DebugOut("request.Score: %v", request.Score)

	incentives := constobjs.GetMileageIncentives(player.MileageMapState.Episode, player.MileageMapState.Chapter) // Game wants incentives in _current_ episode-chapter
	var oldRewardEpisode, newRewardEpisode int64
	var oldRewardChapter, newRewardChapter int64
	var oldRewardPoint, newRewardPoint int64

	if request.Closed == 0 { // If the game wasn't exited out of
		oldRewardEpisode = player.MileageMapState.Episode
		oldRewardChapter = player.MileageMapState.Chapter
		oldRewardPoint = player.MileageMapState.Point
		player.PlayerState.NumRings += request.Rings
		player.PlayerState.NumRedRings += request.RedRings
		player.PlayerState.NumRouletteTicket += request.RedRings // TODO: URGENT! Remove as soon as possible!
		player.PlayerState.Animals += request.Animals
		playerHighScore := player.PlayerState.HighScore
		if request.Score > playerHighScore {
			player.PlayerState.HighScore = request.Score
		}
		player.PlayerState.TotalDistance += request.Distance
		// increase character(s)'s experience
		expIncrease := request.Rings + request.FailureRings // all rings collected
		abilityIndex := 1
		for abilityIndex == 1 { // unused ability is at index 1
			abilityIndex = rand.Intn(len(mainC.AbilityLevel))
		}
		// check that increases exist
		_, ok := consts.UpgradeIncreases[mainC.ID]
		if !ok {
			helper.InternalErr("Error getting upgrade increase", fmt.Errorf("no key '%s' in consts.UpgradeIncreases", mainC.ID))
			return
		}
		_, ok = consts.UpgradeIncreases[subC.ID]
		if !ok {
			helper.InternalErr("Error getting upgrade increase", fmt.Errorf("no key '%s' in consts.UpgradeIncreases", subC.ID))
			return
		}
		if playCharacters[0].Level < 100 {
			playCharacters[0].Exp += expIncrease
			for playCharacters[0].Exp >= playCharacters[0].Cost {
				// more exp than cost = level up
				playCharacters[0].Level++                                               // increase level
				playCharacters[0].AbilityLevel[abilityIndex]++                          // increase ability level
				playCharacters[0].Exp -= playCharacters[0].Cost                         // remove cost from exp
				playCharacters[0].Cost += consts.UpgradeIncreases[playCharacters[0].ID] // increase cost
			}
		}
		if playCharacters[1].Level < 100 {
			playCharacters[1].Exp += expIncrease
			for playCharacters[1].Exp >= playCharacters[1].Cost {
				// more exp than cost = level up
				playCharacters[1].Level++                                               // increase level
				playCharacters[1].AbilityLevel[abilityIndex]++                          // increase ability level
				playCharacters[1].Exp -= playCharacters[1].Cost                         // remove cost from exp
				playCharacters[1].Cost += consts.UpgradeIncreases[playCharacters[1].ID] // increase cost
			}
		}

		helper.DebugOut("Old mainC Exp: " + strconv.Itoa(int(mainC.Exp)) + " / " + strconv.Itoa(int(mainC.Cost)))
		helper.DebugOut("Old mainC Level: " + strconv.Itoa(int(mainC.Level)))
		helper.DebugOut("Old subC Exp: " + strconv.Itoa(int(subC.Exp)) + " / " + strconv.Itoa(int(subC.Cost)))
		helper.DebugOut("Old subC Level: " + strconv.Itoa(int(subC.Level)))
		helper.DebugOut("New mainC Exp: " + strconv.Itoa(int(playCharacters[0].Exp)) + " / " + strconv.Itoa(int(playCharacters[0].Cost)))
		helper.DebugOut("New mainC Level: " + strconv.Itoa(int(playCharacters[0].Level)))
		helper.DebugOut("New subC Exp: " + strconv.Itoa(int(playCharacters[1].Exp)) + " / " + strconv.Itoa(int(playCharacters[1].Cost)))
		helper.DebugOut("New subC Level: " + strconv.Itoa(int(playCharacters[1].Level)))

		/*playCharacters = []netobj.Character{ // TODO: check if this redefinition is needed
			mainC,
			subC,
		}*/

		player.MileageMapState.StageTotalScore += request.Score

		goToNextChapter := request.ChapterClear == 1
		//chaoEggs := request.GetChaoEgg
		// TODO: Add chao eggs to player
		newPoint := request.ReachPoint

		goToNextEpisode := true
		if goToNextChapter {
			// Assumed this just means next episode...
			maxChapters, episodeHasMultipleChapters := consts.EpisodeWithChapters[player.MileageMapState.Episode]
			if episodeHasMultipleChapters {
				goToNextEpisode = false
				player.MileageMapState.Chapter++
				player.MileageMapState.StageTotalScore = 0
				if player.MileageMapState.Chapter > maxChapters {
					// there's no more chapters for this episode!
					goToNextEpisode = true
				}
			}
			if goToNextEpisode {
				player.MileageMapState.Episode++
				player.MileageMapState.Chapter = 1
				player.MileageMapState.Point = 0
				player.MileageMapState.StageTotalScore = 0
				helper.DebugOut("goToNextEpisode -> Episode: %v", player.MileageMapState.Episode)
				if config.CFile.Debug {
					player.MileageMapState.Episode = 15
				}
			}
			if player.MileageMapState.Episode > 50 { // if beat game, reset to 50-1
				player.MileageMapState.Episode = 50
				player.MileageMapState.Chapter = 1
				player.MileageMapState.Point = 0
				player.MileageMapState.StageTotalScore = 0
				helper.DebugOut("goToNextEpisode: Player (%s) beat the game!", player.ID)
			}
		} else {
			player.MileageMapState.Point = newPoint
		}
		if config.CFile.Debug {
			if player.MileageMapState.Episode < 14 {
				player.MileageMapState.Episode = 14
			}
		}
		newRewardEpisode = player.MileageMapState.Episode
		newRewardChapter = player.MileageMapState.Chapter
		newRewardPoint = player.MileageMapState.Point
		// add rewards to PlayerState
		wonRewards := campaign.GetWonRewards(oldRewardEpisode, oldRewardChapter, oldRewardPoint, newRewardEpisode, newRewardChapter, newRewardPoint)
		helper.DebugOut("wonRewards length: %v", wonRewards)
		helper.DebugOut("Previous rings: %v", player.PlayerState.NumRings)
		newItems := player.PlayerState.Items
		for _, reward := range wonRewards { // TODO: This is O(n^2). Maybe alleviate this?
			helper.DebugOut("Reward: %s", reward.ItemID)
			helper.DebugOut("Reward amount: %v", reward.NumItem)
			if reward.ItemID[2:] == "12" { // ID is an item
				// check if the item is already in the player's inventory
				for _, item := range player.PlayerState.Items {
					if item.ID == reward.ItemID { // item found, increment amount
						item.Amount += reward.NumItem
						break
					}
				}
			} else if reward.ItemID == strconv.Itoa(enums.ItemIDRing) { // Rings
				player.PlayerState.NumRings += reward.NumItem
			} else if reward.ItemID == strconv.Itoa(enums.ItemIDRedRing) { // Red rings
				player.PlayerState.NumRedRings += reward.NumItem
			} else {
				helper.Out("Unknown reward '" + reward.ItemID + "', ignoring")
			}
			// TODO: allow for characters to join the cast, like Tails on 11-1.1
		}
		helper.DebugOut("Current rings: %v", player.PlayerState.NumRings)
		player.PlayerState.Items = newItems
	}

	helper.DebugOut("Chapter: %v", player.MileageMapState.Chapter)
	helper.DebugOut("Episode: %v", player.MileageMapState.Episode)
	helper.DebugOut("StageTotalScore: %v", player.MileageMapState.StageTotalScore)
	helper.DebugOut("Point: %v", player.MileageMapState.Point)
	helper.DebugOut("request.Score: %v", request.Score)

	mainCIndex := player.IndexOfChara(mainC.ID) // TODO: check if -1
	subCIndex := player.IndexOfChara(subC.ID)   // TODO: check if -1

	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultPostGameResults(baseInfo, player, playCharacters, incentives)

	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}

	// apply the save after the response so that we don't break the leveling
	mainC = playCharacters[0]
	subC = playCharacters[1]
	player.CharacterState[mainCIndex] = mainC
	player.CharacterState[subCIndex] = subC
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}
	_, err = analytics.Store(player.ID, factors.AnalyticTypeStoryEnds)
	if err != nil {
		helper.WarnErr("Error storing analytics (AnalyticTypeStoryEnds)", err)
	}
}

func GetFreeItemList(helper *helper.Helper) {
	// TODO: allow free items to be set via config
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultFreeItemList(baseInfo)
	//response := responses.FreeItemList(baseInfo, []obj.Item{}) // No free items
	err := helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetMileageReward(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.MileageRewardRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}
	/*
		player, err := helper.GetCallingPlayer()
		if err != nil {
			helper.InternalErr("Error getting calling player", err)
			return
		}
	*/
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultMileageReward(baseInfo, request.Chapter, request.Episode)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}
