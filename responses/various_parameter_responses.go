package responses

import (
    "github.com/fluofoxxo/outrun/playerdata"
)

type VariousParameterResponse struct {
    BaseResponse
    pdata.PlayerVarious
}

func MakeVariousParameterResponse(base BaseInfo, cmSkipCount, energyRecoveryMax, energyRecveryTime, onePlayCmCount, onePlayContinueCount, isPurchased int64) VariousParameterResponse {
    pv := pdata.MakePlayerVarious(cmSkipCount, energyRecoveryMax, energyRecveryTime, onePlayCmCount, onePlayContinueCount, isPurchased)
    br := NewBaseResponse(base)
    vpr := VariousParameterResponse{
        br,
        pv,
    }
    return vpr
}

func NewVariousParameterResponse(base BaseInfo) VariousParameterResponse {
    // uses constants to fill out.
    // TODO: put in consts
    vpr := MakeVariousParameterResponse(
        base,
        1,
        5,
        900,
        1,
        4,
        0,
    )
    return vpr
}
