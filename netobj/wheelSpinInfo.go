package netobj

import (
    "time"

    "github.com/fluofoxxo/outrun/enums"
)

type WheelSpinInfo struct {
    ID    string `json:"id"`
    Start int64  `json:"start"`
    End   int64  `json:"end"`
    Param string `json:"param"`
}

func NewWheelSpinInfo(id, param string) WheelSpinInfo {
    return WheelSpinInfo{
        id,
        time.Now().UTC().Unix(),
        time.Now().UTC().Unix() + 7300, // 2 hours + 100s from now
        param,
    }
}

func DefaultWheelSpinInfoList() []WheelSpinInfo {
    return []WheelSpinInfo{
        NewWheelSpinInfo(enums.ItemIDStrAsteroid, "This"),
        NewWheelSpinInfo(enums.ItemIDStrAsteroid, "is"),
        NewWheelSpinInfo(enums.ItemIDStrAsteroid, "a"),
        NewWheelSpinInfo(enums.ItemIDStrAsteroid, "test"),
        NewWheelSpinInfo(enums.ItemIDStrAsteroid, "message,"),
        NewWheelSpinInfo(enums.ItemIDStrAsteroid, "but"),
        NewWheelSpinInfo(enums.ItemIDStrAsteroid, "not"),
        NewWheelSpinInfo(enums.ItemIDStrAsteroid, "joined!"),
    }
}
