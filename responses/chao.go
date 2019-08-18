package responses

import (
    "github.com/fluofoxxo/outrun/obj"
    "github.com/fluofoxxo/outrun/responses/responseobjs"
)

type ChaoWheelOptionsResponse struct {
    BaseResponse
    ChaoWheelOptions obj.ChaoWheelOptions `json:"chaoWheelOptions"`
}

func ChaoWheelOptions(base responseobjs.BaseInfo, chaoWheelOptions obj.ChaoWheelOptions) ChaoWheelOptionsResponse {
    baseResponse := NewBaseResponse(base)
    out := ChaoWheelOptionsResponse{
        baseResponse,
        chaoWheelOptions,
    }
    return out
}

func DefaultChaoWheelOptions(base responseobjs.BaseInfo) ChaoWheelOptionsResponse {
    chaoWheelOptions := obj.DefaultChaoWheelOptions()
    return ChaoWheelOptions(
        base,
        chaoWheelOptions,
    )
}
