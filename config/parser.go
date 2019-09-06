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
)

var CFile ConfigFile

type ConfigFile struct {
    Port               string `json:"port"`
    DoTimeLogging      bool   `json:"doTimeLogging"`
    LogUnknownRequests bool   `json:"logUnknownRequests"`
}

func Parse(filename string) error {
    port := DPort
    doTimeLogging := DDoTimeLogging
    logUnknownRequests := DLogUnknownRequests

    var ret error
    file, err := loadFile(filename)
    if err == nil {
        var cf ConfigFile
        err := json.Unmarshal(file, &cf)
        if err == nil {
            port = cf.Port
            doTimeLogging = cf.DoTimeLogging
            logUnknownRequests = cf.LogUnknownRequests
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
