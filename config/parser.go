package config

import (
	"encoding/json"
	"io/ioutil"
)

// defaults
// default variable names MUST be "D" + (nameOfVariable)
var Defaults = map[string]interface{}{
	"DPort":                     "9001",
	"DDoTimeLogging":            true,
	"DLogUnknownRequests":       true,
	"DLogAllRequests":           false,
	"DLogAllResponses":          false,
	"DDebug":                    false,
	"DDebugPrints":              false,
	"DEnableRPC":                false,
	"DRPCPort":                  "23432",
	"DEnablePublicStats":        false,
	"DEndpointPrefix":           "",
	"DEnableAnalytics":          false,
	"DPrintPlayerNames":         false,
	"DEventConfigFilename":      "event_config.json",
	"DSilenceEventConfigErrors": false,
	"DInfoConfigFilename":       "info_config.json",
	"DSilenceInfoConfigErrors":  false,
	"DGameConfigFilename":       "game_config.json",
	"DSilenceGameConfigErrors":  false,
}

var CFile ConfigFile

type ConfigFile struct {
	Port                     string `json:"port,omitempty"`
	DoTimeLogging            bool   `json:"doTimeLogging,omitempty"`
	LogUnknownRequests       bool   `json:"logUnknownRequests,omitempty"`
	LogAllRequests           bool   `json:"logAllRequests,omitempty"`
	LogAllResponses          bool   `json:"logAllResponses,omitempty"`
	Debug                    bool   `json:"debug,omitempty"`
	DebugPrints              bool   `json:"debugPrints,omitempty"`
	EnableRPC                bool   `json:"enableRPC,omitempty"`
	RPCPort                  string `json:"rpcPort,omitempty"`
	EnablePublicStats        bool   `json:"enablePublicStats,omitempty"`
	EndpointPrefix           string `json:"endpointPrefix,omitempty"`
	EnableAnalytics          bool   `json:"enableAnalytics,omitempty"`
	PrintPlayerNames         bool   `json:"printPlayerNames,omitempty"`
	EventConfigFilename      string `json:"eventConfigFilename,omitempty"`
	SilenceEventConfigErrors bool   `json:"silenceEventConfigErrors,omitempty"`
	InfoConfigFilename       string `json:"infoConfigFilename,omitempty"`
	SilenceInfoConfigErrors  bool   `json:"silenceInfoConfigErrors,omitempty"`
	GameConfigFilename       string `json:"gameConfigFilename,omitempty"`
	SilenceGameConfigErrors  bool   `json:"silenceGameConfigErrors,omitempty"`
}

func Parse(filename string) error {
	CFile = ConfigFile{
		Defaults["DPort"].(string),
		Defaults["DDoTimeLogging"].(bool),
		Defaults["DLogUnknownRequests"].(bool),
		Defaults["DLogAllRequests"].(bool),
		Defaults["DLogAllResponses"].(bool),
		Defaults["DDebug"].(bool),
		Defaults["DDebugPrints"].(bool),
		Defaults["DEnableRPC"].(bool),
		Defaults["DRPCPort"].(string),
		Defaults["DEnablePublicStats"].(bool),
		Defaults["DEndpointPrefix"].(string),
		Defaults["DEnableAnalytics"].(bool),
		Defaults["DPrintPlayerNames"].(bool),
		Defaults["DEventConfigFilename"].(string),
		Defaults["DSilenceEventConfigErrors"].(bool),
		Defaults["DInfoConfigFilename"].(string),
		Defaults["DSilenceInfoConfigErrors"].(bool),
		Defaults["DGameConfigFilename"].(string),
		Defaults["DSilenceGameConfigErrors"].(bool),
	}
	file, err := loadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &CFile)
	if err != nil {
		return err
	}
	return nil
}

func loadFile(filename string) ([]byte, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return b, err
}
