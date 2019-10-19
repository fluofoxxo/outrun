package muxhandlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/fluofoxxo/outrun/config"
	"github.com/fluofoxxo/outrun/consts"
	"github.com/fluofoxxo/outrun/db"
	"github.com/fluofoxxo/outrun/emess"
	"github.com/fluofoxxo/outrun/enums"
	"github.com/fluofoxxo/outrun/helper"
	"github.com/fluofoxxo/outrun/logic/roulette"
	"github.com/fluofoxxo/outrun/netobj"
	"github.com/fluofoxxo/outrun/obj"
	"github.com/fluofoxxo/outrun/requests"
	"github.com/fluofoxxo/outrun/responses"
	"github.com/fluofoxxo/outrun/status"
)

func GetChaoWheelOptions(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultChaoWheelOptions(baseInfo, player)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func GetPrizeChaoWheelSpin(helper *helper.Helper) {
	// agnostic
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.DefaultPrizeChaoWheel(baseInfo)
	err := helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func EquipChao(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.EquipChaoRequest
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

	mainChaoID := request.MainChaoID
	subChaoID := request.SubChaoID
	if mainChaoID != "-1" {
		player.PlayerState.MainChaoID = mainChaoID
	}
	if subChaoID != "-1" {
		player.PlayerState.SubChaoID = subChaoID
	}
	if config.CFile.DebugPrints {
		helper.Out("Main Chao: " + mainChaoID)
		helper.Out("Sub Chao: " + subChaoID)
	}
	if config.CFile.Debug {
		// TODO: remove
		player.PlayerState.NumRedRings += 150
	}
	db.SavePlayer(player)

	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.EquipChao(baseInfo, player.PlayerState)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func CommitChaoWheelSpin(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}

	data := helper.GetGameRequest()
	var request requests.CommitChaoWheelSpinRequest
	err = json.Unmarshal(data, &request)
	if err != nil {
		helper.InternalErr("Error unmarshalling", err)
	}

	items := player.ChaoRouletteGroup.WheelChao
	weights := player.ChaoRouletteGroup.ChaoWheelOptions.ItemWeight
	availStatus := status.OK
	// set initial prize
	prize := netobj.CharacterIDToChaoSpinPrize("0") // This will almost certainly give the game errors if improperly counting payment!
	spinResults := []netobj.ChaoSpinResult{}        // TODO: Find out why it's an array

	if config.CFile.DebugPrints {
		spf := func(a ...interface{}) string {
			return fmt.Sprintf("%v", a...)
		}
		helper.Out("PRE")
		helper.Out("Items: " + spf(items))
		helper.Out("Weights: " + spf(weights))
		helper.Out("Chao Eggs (Player): " + spf(player.PlayerState.ChaoEggs))
		helper.Out("Chao Eggs (ChaoWheelOptions): " + spf(player.ChaoRouletteGroup.ChaoWheelOptions.NumSpecialEgg))
		helper.Out("Chao Roulette tickets (Player): " + spf(player.PlayerState.NumChaoRouletteTicket))
		helper.Out("Chao Roulette tickets (ChaoWheelOptions): " + spf(player.ChaoRouletteGroup.ChaoWheelOptions.NumChaoRouletteToken))
		helper.Out("Chao Roulette spin cost: " + spf(player.ChaoRouletteGroup.ChaoWheelOptions.SpinCost))
		helper.Out("Tails stars: " + spf(player.CharacterState[1].Star)) // TODO: volatile, remove
		helper.Out("---------------------------------------")
	}

	// reset ChaoRouletteInfo if needed
	rightNow := time.Now().Unix()
	if rightNow > player.ChaoRouletteGroup.ChaoRouletteInfo.RoulettePeriodEnd { // if past period
		player.ChaoRouletteGroup.ChaoRouletteInfo = netobj.DefaultRouletteInfo() // reset all values
	}

	// spin logic
	primaryLogic := func(usingTickets bool) {
		if usingTickets { // paying with ticket(s)
			player.PlayerState.NumChaoRouletteTicket -= consts.ChaoRouletteTicketCost * request.Count // spend ticket(s)
		} else { // paying with red ring(s)
			player.PlayerState.NumRedRings -= consts.ChaoRouletteRedRingCost * request.Count // spend red ring(s)
		}
		player.ChaoRouletteGroup.ChaoRouletteInfo.RouletteCountInPeriod++ // increment times spun in timer; TODO: Should we count request.Count?
		actions := request.Count
		for actions > 0 {
			actions--
			gottenItemIndex, err := roulette.ChooseChaoRouletteItemIndex(items, weights) // pick a potential item index (used for later)
			if err != nil {
				helper.Err("Error choosing Chao roulette item", err)
				return
			}
			gottenItem := items[gottenItemIndex]                       // ID of prize
			gottenPrize := netobj.GenericIDToChaoSpinPrize(gottenItem) // convert ID to prize
			prize = gottenPrize
			spinResult := netobj.ChaoSpinResult{
				prize,
				[]obj.Item{},           // TODO: Research purpose
				int64(gottenItemIndex), // This might be incorrect (ItemWon)
			}
			if prize.Rarity == 100 { // Character
				// increase character level by (amount)
				charIndex := player.IndexOfChara(prize.ID)
				if charIndex == -1 { // character index not found, should never happen
					helper.InternalErr("cannot get index of character '"+strconv.Itoa(charIndex)+"'", err)
					return
				}
				starUpCount := consts.ChaoRouletteCharacterStarIncrease
				for starUpCount > 0 && player.CharacterState[charIndex].Star < 10 { // 10 is max amount of stars a character can have before game breaks
					starUpCount--
					player.CharacterState[charIndex].Star++
				}
				/*
					// The following code is for leveling up characters when gotten by the
					// roulette. This has been dumped in favor of giving a star to the
					// character.
					// TODO: remove
					levelUpCharacter := func() error {
						levelIncrease, ok := consts.UpgradeIncreases[player.CharacterState[charIndex].ID]
						if !ok {
							return fmt.Errorf("key '%v' not found in consts.UpgradeIncreases", player.CharacterState[charIndex].ID)
						}
						player.CharacterState[charIndex].AbilityLevel[rand.Intn(len(player.CharacterState[charIndex].AbilityLevel))]++ // upgrade random ability
						player.CharacterState[charIndex].Level += 1
						player.CharacterState[charIndex].Exp = 0
						player.CharacterState[charIndex].Cost += levelIncrease
						return nil
					}
					levelUpCount := consts.ChaoRouletteCharacterLevelIncrease
					for levelUpCount > 0 { // level up (consts.ChaoRouletteCharacterLevelIncrease) times
						levelUpCount--
						err := levelUpCharacter()
						if err != nil {
							helper.InternalErr("Error levelling up character", err)
							return
						}

						if player.CharacterState[charIndex].Level > 100 { // if limit break
							// reset all level based character values, do limit break
							player.CharacterState[charIndex].Level = 0
							player.CharacterState[charIndex].AbilityLevel = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
							player.CharacterState[charIndex].AbilityNumRings = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
							player.CharacterState[charIndex].AbilityLevelUpExp = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
							player.CharacterState[charIndex].Star++
							if player.CharacterState[charIndex].Star >= player.CharacterState[charIndex].StarMax { // if exceeded max amount of stars
								// TODO: then what?
								player.CharacterState[charIndex].Star = player.CharacterState[charIndex].StarMax
							}
						}
					}
				*/
				spinResult.WonPrize.Level = player.CharacterState[charIndex].Level // set level of prize to character level
			} else if prize.Rarity == 2 || prize.Rarity == 1 || prize.Rarity == 0 { // Chao
				chaoIndex := player.IndexOfChao(prize.ID)
				if chaoIndex == -1 { // chao index not found, should never happen
					helper.InternalErr("cannot get index of chao '"+strconv.Itoa(chaoIndex)+"'", err)
					return
				}
				highRange := int(consts.ChaoRouletteChaoLevelIncreaseHigh)
				lowRange := int(consts.ChaoRouletteChaoLevelIncreaseLow)
				prizeChaoLevel := int64(rand.Intn(highRange-lowRange+1) + lowRange) // This level is added to the current Chao level
				player.ChaoState[chaoIndex].Level += prizeChaoLevel
				if player.ChaoState[chaoIndex].Level > 10 { // if max chao level (https://www.deviantart.com/vocaloidbrsfreak97/journal/So-Sonic-Runners-just-recently-updated-574789098)
					excess := player.ChaoState[chaoIndex].Level - 10 // get amount gone over
					prizeChaoLevel -= excess                         // shave it from prize level
					player.ChaoState[chaoIndex].Level = 10           // reset to maximum
				}
				spinResult.WonPrize.Level = player.ChaoState[chaoIndex].Level
			} else { // Should never happen!
				helper.InternalErr("unknown prize rarity '"+strconv.Itoa(int(prize.Rarity))+"'", fmt.Errorf("")) // TODO: Probably shouldn't use a blank error?
			}
			spinResults = append(spinResults, spinResult) // add spin result to results list (See spinResults declaration)
		}
		// create a new wheel; must be done after ALL player operations are done
		chaoCanBeLevelled := !player.AllChaoMaxLevel()
		charactersCanBeLevelled := !player.AllCharactersMaxLevel()
		if config.CFile.DebugPrints {
			helper.Out("Chao can be levelled: " + strconv.FormatBool(chaoCanBeLevelled))
			helper.Out("Characters can be levelled: " + strconv.FormatBool(charactersCanBeLevelled))
		}
		fixRarities := func(rarities []int64) ([]int64, bool) {
			newRarities := []int64{}
			if !chaoCanBeLevelled && !charactersCanBeLevelled {
				// Wow, they can't upgrade _anything!_
				return newRarities, false
			}
			if config.CFile.Debug {
				player.PlayerState.NumRedRings += 150
				//return []int64{100, 100, 100, 100, 100, 100, 100, 100}, true
				return []int64{0, 0, 0, 0, 0, 0, 0, 0}, true
			}
			for _, r := range rarities {
				if r == 0 || r == 1 || r == 2 { // Chao
					if chaoCanBeLevelled {
						newRarities = append(newRarities, r)
					} else {
						newRarities = append(newRarities, 100) // append a character
					}
				} else if r == 100 { // character
					if charactersCanBeLevelled {
						newRarities = append(newRarities, r)
					} else {
						newRarities = append(newRarities, int64(rand.Intn(3))) // append random rarity Chao
					}
				} else { // should never happen
					panic(fmt.Errorf("invalid rarity '" + strconv.Itoa(int(r)) + "'")) // TODO: use better way to handle
				}
			}
			return newRarities, true
		}
		player.ChaoRouletteGroup.ChaoWheelOptions = netobj.DefaultChaoWheelOptions(player.PlayerState) // create a new wheel
		newRarities, ok := fixRarities(player.ChaoRouletteGroup.ChaoWheelOptions.Rarity)
		if !ok { // if player is entirely unable to upgrade anything
			// TODO: this is probably not the right way to do this!
			player.ChaoRouletteGroup.ChaoWheelOptions.SpinCost = player.PlayerState.NumChaoRouletteTicket + player.PlayerState.NumRedRings // make it impossible for player to use roulette
		} else { // if player can upgrade
			player.ChaoRouletteGroup.ChaoWheelOptions.Rarity = newRarities
		}
		//newItems, err := roulette.GetRandomChaoRouletteItems(player.ChaoRouletteGroup.ChaoWheelOptions.Rarity, player.GetAllMaxLevelIDs()) // create new wheel items
		//newItems, err := roulette.GetRandomChaoRouletteItems(player.ChaoRouletteGroup.ChaoWheelOptions.Rarity, player.GetAllNonMaxedChaoAndCharacters()) // create new wheel items
		newItems, newRarities, err := roulette.GetRandomChaoRouletteItems(player.ChaoRouletteGroup.ChaoWheelOptions.Rarity, player.GetAllNonMaxedCharacters(), player.GetAllNonMaxedChao())
		if err != nil {
			helper.InternalErr("Error getting new items", err)
			return
		}
		player.ChaoRouletteGroup.WheelChao = newItems
		player.ChaoRouletteGroup.ChaoWheelOptions.Rarity = newRarities
		if config.CFile.DebugPrints {
			helper.Out(fmt.Sprintf("%v", newRarities))
		}
		if config.CFile.Debug {
			player.ChaoRouletteGroup.WheelChao = []string{enums.CTStrTails, enums.CTStrTails, enums.CTStrTails, enums.CTStrTails, enums.CTStrTails, enums.CTStrTails, enums.CTStrTails, enums.CTStrTails}
		}
	}

	hasTickets := player.PlayerState.NumChaoRouletteTicket >= consts.ChaoRouletteTicketCost*request.Count
	hasAvailableRings := player.PlayerState.NumRedRings >= consts.ChaoRouletteRedRingCost*request.Count

	if hasTickets { // if tickets to spend
		primaryLogic(true)
	} else if hasAvailableRings { // if no tickets, but sufficient red rings
		primaryLogic(false)
	} else { // no tickets nor sufficient red rings
		availStatus = status.RouletteUseLimit
	}

	if config.CFile.DebugPrints {
		spf := func(a ...interface{}) string {
			return fmt.Sprintf("%v", a...)
		}
		helper.Out("POST")
		helper.Out("Items: " + spf(player.ChaoRouletteGroup.WheelChao))
		helper.Out("Weights: " + spf(player.ChaoRouletteGroup.ChaoWheelOptions.ItemWeight))
		helper.Out("Chao Eggs (Player): " + spf(player.PlayerState.ChaoEggs))
		helper.Out("Chao Eggs (ChaoWheelOptions): " + spf(player.ChaoRouletteGroup.ChaoWheelOptions.NumSpecialEgg))
		helper.Out("Chao Roulette tickets (Player): " + spf(player.PlayerState.NumChaoRouletteTicket))
		helper.Out("Chao Roulette tickets (ChaoWheelOptions): " + spf(player.ChaoRouletteGroup.ChaoWheelOptions.NumChaoRouletteToken))
		helper.Out("Chao Roulette spin cost: " + spf(player.ChaoRouletteGroup.ChaoWheelOptions.SpinCost))
		helper.Out("Tails stars: " + spf(player.CharacterState[1].Star)) // TODO: volatile, remove
	}

	baseInfo := helper.BaseInfo(emess.OK, availStatus)
	response := responses.ChaoWheelSpin(baseInfo, player.PlayerState, player.CharacterState, player.ChaoState, player.ChaoRouletteGroup.ChaoWheelOptions, spinResults)

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
