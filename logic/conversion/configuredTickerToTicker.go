package conversion

import (
    "time"

    "github.com/fluofoxxo/outrun/config/infoconf"
    "github.com/fluofoxxo/outrun/obj"
    "github.com/jinzhu/now"
)

func ConfiguredTickerToTicker(index int64, ct infoconf.ConfiguredTicker) obj.Ticker {
    // Should be used by the game as soon as possible
    startTime := ct.StartTime
    switch startTime {
    case -2:
        startTime = now.BeginningOfDay().Unix()
    case -3:
        startTime = now.EndOfDay().Unix()
    case -4:
        startTime = time.Now().Unix() - 1
    }
    endTime := ct.EndTime
    switch endTime {
    case -2:
        endTime = now.BeginningOfDay().Unix()
    case -3:
        endTime = now.EndOfDay().Unix()
    case -4:
        endTime = time.Now().Add(24 * time.Hour).Unix()
    }
    newTicker := obj.Ticker{
        index,
        startTime,
        endTime,
        ct.Message,
    }
    return newTicker
}
