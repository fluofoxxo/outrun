package muxhandlers

import (
	"fmt"
	"time"

	"github.com/fluofoxxo/outrun/config"
	"github.com/fluofoxxo/outrun/config/eventconf"
	"github.com/fluofoxxo/outrun/emess"
	"github.com/fluofoxxo/outrun/helper"
	"github.com/fluofoxxo/outrun/obj"
	"github.com/fluofoxxo/outrun/responses"
	"github.com/fluofoxxo/outrun/status"
	"github.com/jinzhu/now"
)

func GetEventList(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	// construct event list
	eventList := []obj.Event{}
	if eventconf.CFile.AllowEvents {
		if eventconf.CFile.EnforceGlobal || len(player.PersonalEvents) == 0 {
			for _, confEvent := range eventconf.CFile.CurrentEvents {
				startTime := confEvent.StartTime
				switch startTime {
				case -2:
					startTime = now.BeginningOfDay().Unix()
				case -3:
					startTime = now.EndOfDay().Unix()
				case -4:
					startTime = time.Now().Unix() - 1
				}
				endTime := confEvent.EndTime
				switch endTime {
				case -2:
					endTime = now.BeginningOfDay().Unix()
				case -3:
					endTime = now.EndOfDay().Unix()
				case -4:
					endTime = time.Now().Add(24 * time.Hour).Unix()
				}
				newEvent := obj.Event{
					// Ex. ID: 712340000
					(confEvent.ID * 10000) + confEvent.RealType(), // make a real event ID
					confEvent.RealType(),
					startTime,
					endTime,
					endTime,
				}
				eventList = append(eventList, newEvent)
			}
		} else {
			eventList = append(eventList, player.PersonalEvents...)
		}
	}
	if config.CFile.DebugPrints {
		helper.Out("Personal event list: " + fmt.Sprintln(player.PersonalEvents))
		helper.Out("Global event list: " + fmt.Sprintln(eventconf.CFile.CurrentEvents))
		helper.Out("Event list: " + fmt.Sprintln(eventList))
	}
	response := responses.EventList(baseInfo, eventList)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}
