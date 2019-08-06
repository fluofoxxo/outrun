package responses

import (
    "github.com/fluofoxxo/outrun/playerdata"
)

type ChaoStateResponse struct {
    BaseResponse
    ChaoState []pdata.Chao `json:"chaoState"`
}

func NewChaoStateResponse(base BaseInfo, chao []pdata.Chao) ChaoStateResponse {
    br := NewBaseResponse(base)
    csr := ChaoStateResponse{
        br,
        chao,
    }
    return csr
}
