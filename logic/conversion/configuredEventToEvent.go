package conversion

import (
    "time"

    "github.com/fluofoxxo/outrun/config/eventconf"
    "github.com/fluofoxxo/outrun/obj"
    "github.com/jinzhu/now"
)

func ConfiguredEventToEvent(ce eventconf.ConfiguredEvent) obj.Event {
    // Should be used by the game as soon as possible
    startTime := ce.StartTime
    switch startTime {
    case -2:
        startTime = now.BeginningOfDay().Unix()
    case -3:
        startTime = now.EndOfDay().Unix()
    case -4:
        startTime = time.Now().Unix() - 1
    }
    endTime := ce.EndTime
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
        (ce.ID * 10000) + ce.RealType(), // make a real event ID
        ce.RealType(),
        startTime,
        endTime,
        endTime,
    }
    return newEvent
}
