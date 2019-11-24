package eventconf

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/fluofoxxo/outrun/enums"
)

// defaults
// default variable names MUST be "D" + (nameOfVariable)
var Defaults = map[string]interface{}{
	"DAllowEvents":   false,
	"DCurrentEvents": []ConfiguredEvent{},
	"DEnforceGlobal": false,
}

var EventTypes = map[string]int64{
	"specialStage":  enums.EventIDSpecialStage,
	"raidBoss":      enums.EventIDRaidBoss,
	"collectObject": enums.EventIDCollectObject,
	"gacha":         enums.EventIDGacha,
	"advert":        enums.EventIDAdvert,
	"quick":         enums.EventIDQuick,
	"bgm":           enums.EventIDBGM,
}

type ConfiguredEvent struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	StartTime int64  `json:"startTime"` // *
	EndTime   int64  `json:"endTime"`   // *
}

/*
* StartTime and EndTime can use certain negative values to represent time
functions that are not otherwise easily possible using a static
JSON file. -2 will be the beginning of the system's day, and -3 will
be the end of the system's day.
A value of -4 will activate the event no matter what. For StartTime, it
will use the current time minus 1 in order to say the event has just started.
For EndTime, it will use 24 hours into the future. This isn't foolproof, as
the event could expire and yield odd results, but if you have a player that
is playing for long enough for that to happen, you have a _different_ problem
on your hands.
*/

func (c ConfiguredEvent) RealType() int64 {
	return EventTypes[c.Type]
}
func (c ConfiguredEvent) HasValidType() bool {
	_, ok := EventTypes[c.Type]
	return ok
}

var CFile ConfigFile

type ConfigFile struct {
	AllowEvents   bool              `json:"allowEvents,omitempty"`
	CurrentEvents []ConfiguredEvent `json:"currentEvents,omitempty"`
	EnforceGlobal bool              `json:"enforceGlobal,omitempty"`
}

func Parse(filename string) error {
	CFile = ConfigFile{
		Defaults["DAllowEvents"].(bool),
		Defaults["DCurrentEvents"].([]ConfiguredEvent),
		Defaults["DEnforceGlobal"].(bool),
	}
	file, err := loadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &CFile)
	if err != nil {
		return err
	}
	newEvents := []ConfiguredEvent{}
	for i, ce := range CFile.CurrentEvents {
		if !ce.HasValidType() {
			log.Printf("[WARN] Invalid event type %s at index %v, ignoring\n", ce.Type, i)
			continue
		}
		newEvents = append(newEvents, ce)
	}
	CFile.CurrentEvents = newEvents
	return nil
}

func loadFile(filename string) ([]byte, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return b, err
}
