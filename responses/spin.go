package responses

import (
    "github.com/fluofoxxo/outrun/netobj"
    "github.com/fluofoxxo/outrun/responses/responseobjs"
)

type WheelOptionsResponse struct {
    BaseResponse
    WheelOptions netobj.WheelOptions `json:"wheelOptions"`
}

func WheelOptions(base responseobjs.BaseInfo, wheelOptions netobj.WheelOptions) WheelOptionsResponse {
    baseResponse := NewBaseResponse(base)
    out := WheelOptionsResponse{
        baseResponse,
        wheelOptions,
    }
    return out
}
