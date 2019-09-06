package config

import (
	"encoding/json"
	"io/ioutil"
)

// defaults
const (
	DPort               = "9001"
	DDoTimeLogging      = true
	DLogUnknownRequests = true
	DLogAllRequests     = false
	DDebug              = false
)

var CFile ConfigFile

type ConfigFile struct {
	Port               string `json:"port"`
	DoTimeLogging      bool   `json:"doTimeLogging"`
	LogUnknownRequests bool   `json:"logUnknownRequests"`
	LogAllRequests     bool   `json:"logAllRequests"`
	Debug              bool   `json:"debug"`
}

func Parse(filename string) error {
	port := DPort
	doTimeLogging := DDoTimeLogging
	logUnknownRequests := DLogUnknownRequests
	logAllRequests := DLogAllRequests
	debug := DDebug

	var ret error
	file, err := loadFile(filename)
	if err == nil {
		var cf ConfigFile
		err := json.Unmarshal(file, &cf)
		if err == nil {
			port = cf.Port
			doTimeLogging = cf.DoTimeLogging
			logUnknownRequests = cf.LogUnknownRequests
			logAllRequests = cf.LogAllRequests
			debug = cf.Debug
		} else {
			ret = err
		}
	} else {
		ret = err
	}

	CFile = ConfigFile{
		port,
		doTimeLogging,
		logUnknownRequests,
		logAllRequests,
		debug,
	}
	return ret
}

func loadFile(filename string) ([]byte, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return b, err
}
