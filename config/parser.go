package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

// defaults
// default variable names MUST be "D" + (nameOfVariable)
var Defaults = map[string]interface{}{
	"DPort":               "9001",
	"DDoTimeLogging":      true,
	"DLogUnknownRequests": true,
	"DLogAllRequests":     false,
	"DDebug":              false,
	"DDebugPrints":        false,
	"DRPCPort":            "23432",
	"DEnablePublicStats":  false,
}

var CFile ConfigFile

type ConfigFile struct {
	Port               string `json:"port,omitempty"`
	DoTimeLogging      bool   `json:"doTimeLogging,omitempty"`
	LogUnknownRequests bool   `json:"logUnknownRequests,omitempty"`
	LogAllRequests     bool   `json:"logAllRequests,omitempty"`
	Debug              bool   `json:"debug,omitempty"`
	DebugPrints        bool   `json:"debugPrints,omitempty"`
	RPCPort            string `json:"rpcPort,omitempty"`
	EnablePublicStats  bool   `json:"enablePublicStats,omitempty"`
}

func Parse(filename string) error {
	var ret error
	file, err := loadFile(filename)
	if err == nil {
		var cf ConfigFile
		err := json.Unmarshal(file, &cf)
		if err == nil {
			vo := reflect.ValueOf(cf)
			rtype := vo.Type()
			for i := 0; i < vo.NumField(); i++ {
				fieldname := rtype.Field(i).Name
				fieldVal := vo.Field(i).Interface()
				defaultName := "D" + fieldname
				defaultVal := Defaults[defaultName]
				isZero, err := isZeroVal(fieldVal)
				if err != nil {
					ret = fmt.Errorf("error getting zero value: %s", err)
					return ret
				}
				if isZero {
					reflect.ValueOf(&cf).Elem().Field(i).Set(reflect.ValueOf(defaultVal)) // assign the variable with its default
				}
			}

			CFile = cf
			return ret
		} else {
			ret = err
		}
	} else {
		ret = err
	}

	CFile = ConfigFile{
		Defaults["DPort"].(string),
		Defaults["DDoTimeLogging"].(bool),
		Defaults["DLogUnknownRequests"].(bool),
		Defaults["DLogAllRequests"].(bool),
		Defaults["DDebug"].(bool),
		Defaults["DDebugPrints"].(bool),
		Defaults["DRPCPort"].(string),
		Defaults["DEnablePublicStats"].(bool),
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

func isZeroVal(val interface{}) (bool, error) {
	vartype := reflect.TypeOf(val)
	if !vartype.Comparable() {
		return false, fmt.Errorf("error comparing type %v", vartype)
	}
	return reflect.Zero(vartype).Interface() == val, nil
}
