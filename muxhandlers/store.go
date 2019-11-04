package muxhandlers

import (
	"encoding/json"
	"strconv"

	"github.com/fluofoxxo/outrun/analytics"
	"github.com/fluofoxxo/outrun/analytics/factors"
	"github.com/fluofoxxo/outrun/db"
	"github.com/fluofoxxo/outrun/emess"
	"github.com/fluofoxxo/outrun/helper"
	"github.com/fluofoxxo/outrun/obj"
	"github.com/fluofoxxo/outrun/obj/constobjs"
	"github.com/fluofoxxo/outrun/requests"
	"github.com/fluofoxxo/outrun/responses"
	"github.com/fluofoxxo/outrun/status"
)

func GetRedStarExchangeList(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.RedStarExchangeListRequest
	err := json.Unmarshal(recv, &request)
	if err != nil {
		helper.Err("Error unmarshalling", err)
		return
	}

	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	var response responses.RedStarExchangeListResponse
	var redStarItems []obj.RedStarItem
	helper.Out("Recv ItemType " + strconv.Itoa(int(request.ItemType)))
	if request.ItemType == 0 {
		redStarItems = []obj.RedStarItem{}
	} else if request.ItemType == 1 {
		redStarItems = constobjs.RedStarItemsType1
	} else if request.ItemType == 2 {
		redStarItems = constobjs.RedStarItemsType2
	} else if request.ItemType == 4 {
		redStarItems = constobjs.RedStarItemsType4
	} else {
		helper.Respond([]byte("Invalid request"))
		return
	}

	response = responses.RedStarExchangeList(baseInfo, redStarItems, 0, "1900-1-1")
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}

func RedStarExchange(helper *helper.Helper) {
	recv := helper.GetGameRequest()
	var request requests.RedStarExchange
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

	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	itemID := request.ItemID
	getItemType := func(iid string) (string, int64, bool) {
		var itemType string
		itemPrice, ok := constobjs.ShopRingPrices[iid]
		if !ok {
			// it's not a ring item
			itemPrice, ok = constobjs.ShopEnergyPrices[iid]
			if !ok {
				// it's not ring nor energy item
				// unrecognized item!
				return "", 0, false
			}
			itemType = "energy"
			return itemType, itemPrice, true
		}
		itemType = "ring"
		return itemType, itemPrice, true
	}

	itemType, itemPrice, found := getItemType(itemID)
	if found {
		switch itemType {
		case "ring":
			if player.PlayerState.NumRedRings-itemPrice < 0 {
				baseInfo.StatusCode = status.NotEnoughRedRings
				return
			}
			player.PlayerState.NumRedRings -= itemPrice
			player.PlayerState.NumRings += constobjs.ShopRingAmounts[itemID]
			db.SavePlayer(player)
			_, err = analytics.Store(player.ID, factors.AnalyticTypePurchaseRings)
			if err != nil {
				helper.WarnErr("Error storing analytics (AnalyticTypePurchaseRings)", err)
			}
		case "energy":
			if player.PlayerState.NumRedRings-itemPrice < 0 {
				baseInfo.StatusCode = status.NotEnoughRedRings
				return
			}
			player.PlayerState.NumRedRings -= itemPrice
			player.PlayerState.Energy += constobjs.ShopEnergyAmounts[itemID]
			db.SavePlayer(player)
			_, err = analytics.Store(player.ID, factors.AnalyticTypePurchaseEnergy)
			if err != nil {
				helper.WarnErr("Error storing analytics (AnalyticTypePurchaseEnergy)", err)
			}
		default:
			// this should never execute!
			baseInfo.StatusCode = status.MasterDataMismatch
			helper.Out("Default case executed... Something went wrong!")
		}
	} else {
		baseInfo.StatusCode = status.MasterDataMismatch
	}

	response := responses.DefaultRedStarExchange(baseInfo, player)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}
