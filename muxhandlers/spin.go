package muxhandlers

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/fluofoxxo/outrun/analytics"
	"github.com/fluofoxxo/outrun/analytics/factors"
	"github.com/fluofoxxo/outrun/consts"
	"github.com/fluofoxxo/outrun/db"
	"github.com/fluofoxxo/outrun/emess"
	"github.com/fluofoxxo/outrun/enums"
	"github.com/fluofoxxo/outrun/helper"
	"github.com/fluofoxxo/outrun/logic"
	"github.com/fluofoxxo/outrun/netobj"
	"github.com/fluofoxxo/outrun/requests"
	"github.com/fluofoxxo/outrun/responses"
	"github.com/fluofoxxo/outrun/status"
)

func GetWheelOptions(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)

	//player.LastWheelOptions = netobj.DefaultWheelOptions(player.PlayerState) // generate new wheel for 'reroll' mechanic
	helper.DebugOut("Time now: %s", time.Now().Unix())
	helper.DebugOut("RoulettePeriodEnd: %s", player.RouletteInfo.RoulettePeriodEnd)
	// check if we need to reset the end period
	endPeriod := player.RouletteInfo.RoulettePeriodEnd
	if time.Now().Unix() > endPeriod {
		player.RouletteInfo = netobj.DefaultRouletteInfo() // Effectively reset everything, set new end time
	}

	// refresh wheel
	player.LastWheelOptions = logic.WheelRefreshLogic(player, player.LastWheelOptions)

	response := responses.WheelOptions(baseInfo, player.LastWheelOptions)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func CommitWheelSpin(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.CommitWheelSpinRequest
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
	helper.DebugOut("request.Count: %s", request.Count)

	endPeriod := player.RouletteInfo.RoulettePeriodEnd
	helper.DebugOut("Time now: %s", time.Now().Unix())
	helper.DebugOut("End period: %s", endPeriod)
	if time.Now().Unix() > endPeriod {
		player.RouletteInfo = netobj.DefaultRouletteInfo() // Effectively reset everything, set new end time
		helper.DebugOut("New roulette period")
		helper.DebugOut("RouletteCountInPeriod: %s", player.RouletteInfo.RouletteCountInPeriod)
	}

	responseStatus := status.OK
	hasTickets := player.PlayerState.NumRouletteTicket > 0
	hasFreeSpins := player.RouletteInfo.RouletteCountInPeriod < consts.RouletteFreeSpins
	helper.DebugOut("Has tickets: %s", hasTickets)
	helper.DebugOut("Number of tickets: %s", player.PlayerState.NumRouletteTicket)
	helper.DebugOut("Has free spins: %s", hasFreeSpins)
	helper.DebugOut("Roulette count: %s", player.RouletteInfo.RouletteCountInPeriod)
	if hasTickets || hasFreeSpins {
		//if player.LastWheelOptions.NumRemainingRoulette > 0 {
		wonItem := player.LastWheelOptions.Items[player.LastWheelOptions.ItemWon]
		itemExists := player.IndexOfItem(wonItem) != -1
		if itemExists {
			amountOfItemWon := player.LastWheelOptions.Item[player.LastWheelOptions.ItemWon]
			helper.DebugOut("wonItem: %s", wonItem)
			helper.DebugOut("amountOfItemWon: %s", amountOfItemWon)
			itemIndex := player.IndexOfItem(wonItem)
			helper.DebugOut("Amount of item player has: %s", player.PlayerState.Items[itemIndex].Amount)
			player.PlayerState.Items[itemIndex].Amount += amountOfItemWon
			helper.DebugOut("New amount of item player has: %s", player.PlayerState.Items[itemIndex].Amount)
		} else {
			if wonItem == strconv.Itoa(enums.IDTypeItemRouletteWin) {
				// Jackpot
				player.PlayerState.NumRings += player.LastWheelOptions.NumJackpotRing
			} else if wonItem == strconv.Itoa(enums.IDTypeRedRing) {
				// Red rings
				player.PlayerState.NumRedRings += player.LastWheelOptions.Item[player.LastWheelOptions.ItemWon]
			} else {
				helper.Warn("item '" + wonItem + "' not found")
			}
		}

		helper.DebugOut("Time now: %s", time.Now().Unix())
		helper.DebugOut("RoulettePeriodEnd: %s", player.RouletteInfo.RoulettePeriodEnd)
		endPeriod := player.RouletteInfo.RoulettePeriodEnd
		helper.DebugOut("Time now (passed): %s", time.Now().Unix())
		helper.DebugOut("End period (passed): %s", endPeriod)
		if time.Now().Unix() > endPeriod { // TODO: Do we still need this?
			player.RouletteInfo = netobj.DefaultRouletteInfo() // Effectively reset everything, set new end time
			helper.DebugOut("New roulette period")
			helper.DebugOut("RouletteCountInPeriod", player.RouletteInfo.RouletteCountInPeriod)
		}

		// generate NEXT! wheel
		player.RouletteInfo.RouletteCountInPeriod++ // we've spun an additional time
		if player.RouletteInfo.RouletteCountInPeriod > consts.RouletteFreeSpins {
			// we've run out of free spins for the period
			player.PlayerState.NumRouletteTicket--
		}
		numRouletteTicket := player.PlayerState.NumRouletteTicket
		rouletteCount := player.RouletteInfo.RouletteCountInPeriod                             // get amount of times we've spun the wheel today
		player.LastWheelOptions = netobj.DefaultWheelOptions(numRouletteTicket, rouletteCount) // create wheel
		if player.RouletteInfo.GotJackpotThisPeriod {
			player.LastWheelOptions.NumJackpotRing = 1
		}
		if wonItem == strconv.Itoa(enums.IDTypeItemRouletteWin) {
			player.RouletteInfo.GotJackpotThisPeriod = true
		}
	} else {
		// do not modify the wheel, set error status
		responseStatus = status.RouletteUseLimit
	}

	baseInfo := helper.BaseInfo(emess.OK, responseStatus)
	response := responses.WheelSpin(baseInfo, player.PlayerState, player.CharacterState, player.ChaoState, player.LastWheelOptions)

	err = db.SavePlayer(player)
	if err != nil {
		helper.InternalErr("Error saving player", err)
		return
	}

	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
		return
	}
	_, err = analytics.Store(player.ID, factors.AnalyticTypeSpinItemRoulette)
	if err != nil {
		helper.WarnErr("Error storing analytics (AnalyticTypeSpinItemRoulette)", err)
	}
}
