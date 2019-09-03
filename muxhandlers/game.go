package muxhandlers

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/fluofoxxo/outrun/consts"
	"github.com/fluofoxxo/outrun/db"
	"github.com/fluofoxxo/outrun/emess"
	"github.com/fluofoxxo/outrun/helper"
	"github.com/fluofoxxo/outrun/netobj"
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
	playCharacters := []netobj.Character{
		mainC,
		subC,
	}
	fmt.Printf("Main Character Level (First): %v\n", mainC.Level)
	fmt.Printf("Sub Character Level (First): %v\n", subC.Level)
	if request.Closed == 0 { // If the game wasn't exited out of
		player.PlayerState.NumRings += request.Rings
		player.PlayerState.NumRedRings += request.RedRings
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
		if mainC.Level < 100 {
			mainC.Exp += expIncrease
			for mainC.Exp >= mainC.Cost { // while loop allows for multiple level ups
				// level up!
				levelIncrease, ok := consts.UpgradeIncreases[mainC.ID]
				if !ok {
					helper.InternalErr("Error getting level increase", fmt.Errorf("key '%s' not found in consts.UpgradeIncreases", mainC.ID))
					return
				}
				mainC.Level++
				mainC.Exp -= mainC.Cost
				mainC.Cost += levelIncrease
				mainC.AbilityLevel[abilityIndex]++
				if mainC.Level >= 100 {
					break
				}
			}
		}
		if subC.Level < 100 {
			for subC.Exp >= subC.Cost { // while loop allows for multiple level ups
				subC.Exp += expIncrease
				// level up!
				levelIncrease, ok := consts.UpgradeIncreases[subC.ID]
				if !ok {
					helper.InternalErr("Error getting level increase", fmt.Errorf("key '%s' not found in consts.UpgradeIncreases", subC.ID))
					return
				}
				subC.Level++
				subC.Exp -= subC.Cost
				subC.Cost += levelIncrease
				subC.AbilityLevel[abilityIndex]++
				if subC.Level >= 100 {
					break
				}
			}
		}
		playCharacters = []netobj.Character{ // TODO: check if this redefinition is needed
			mainC,
			subC,
		}
		//err = db.SavePlayer(player)
	}

	/*
		if err != nil {
			helper.InternalErr("Error saving player", err)
			return
		}
	*/

	mainCIndex := player.IndexOfChara(mainC.ID) // TODO: check if -1
	subCIndex := player.IndexOfChara(subC.ID)   // TODO: check if -1

	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultQuickPostGameResults(baseInfo, player, playCharacters)
	// apply the save after the response so that we don't break the leveling
	player.CharacterState[mainCIndex] = mainC
	player.CharacterState[subCIndex] = subC
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}

	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
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
	fmt.Printf("%v\n", mainC.Level)
	if request.Closed == 0 { // If the game wasn't exited out of
		player.PlayerState.NumRings += request.Rings
		player.PlayerState.NumRedRings += request.RedRings
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
		if mainC.Level < 100 {
			mainC.Exp += expIncrease
			for mainC.Exp >= mainC.Cost { // while loop allows for multiple level ups
				// level up!
				levelIncrease, ok := consts.UpgradeIncreases[mainC.ID]
				if !ok {
					helper.InternalErr("Error getting level increase", fmt.Errorf("key '%s' not found in consts.UpgradeIncreases", mainC.ID))
					return
				}
				mainC.Level++
				mainC.Exp -= mainC.Cost
				mainC.Cost += levelIncrease
				mainC.AbilityLevel[abilityIndex]++
				if mainC.Level >= 100 {
					break
				}
			}
		}
		if subC.Level < 100 {
			for subC.Exp >= subC.Cost { // while loop allows for multiple level ups
				subC.Exp += expIncrease
				// level up!
				levelIncrease, ok := consts.UpgradeIncreases[subC.ID]
				if !ok {
					helper.InternalErr("Error getting level increase", fmt.Errorf("key '%s' not found in consts.UpgradeIncreases", subC.ID))
					return
				}
				subC.Level++
				subC.Exp -= subC.Cost
				subC.Cost += levelIncrease
				subC.AbilityLevel[abilityIndex]++
				if subC.Level >= 100 {
					break
				}
			}
		}
		playCharacters = []netobj.Character{ // TODO: check if this redefinition is needed
			mainC,
			subC,
		}
		//err = db.SavePlayer(player)
	}

	/*
		if err != nil {
			helper.InternalErr("Error saving player", err)
			return
		}
	*/

	mainCIndex := player.IndexOfChara(mainC.ID) // TODO: check if -1
	subCIndex := player.IndexOfChara(subC.ID)   // TODO: check if -1

	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultPostGameResults(baseInfo, player, playCharacters)
	// apply the save after the response so that we don't break the leveling
	player.CharacterState[mainCIndex] = mainC
	player.CharacterState[subCIndex] = subC
	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}

	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetFreeItemList(helper *helper.Helper) {
	// Probably agnostic...
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultFreeItemList(baseInfo)
	err := helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}
