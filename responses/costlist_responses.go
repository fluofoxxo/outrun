package responses

import (
    "github.com/fluofoxxo/outrun/objects"
)

type ConsumedCostListResponse struct {
    BaseResponse
    ConsumedCostList []objects.ConsumedItem `json:"consumedCostList"`
}

func NewConsumedCostListResponse(base BaseInfo, consumedCostList []objects.ConsumedItem) ConsumedCostListResponse {
    br := NewBaseResponse(base)
    cclr := ConsumedCostListResponse{
        br,
        consumedCostList,
    }
    return cclr
}
