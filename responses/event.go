package responses

import (
    "github.com/fluofoxxo/outrun/obj"
    "github.com/fluofoxxo/outrun/responses/responseobjs"
)

type EventListResponse struct {
    BaseResponse
    EventList []obj.Event `json:"eventList"`
}

func EventList(base responseobjs.BaseInfo, eventList []obj.Event) EventListResponse {
    baseResponse := NewBaseResponse(base)
    out := EventListResponse{
        baseResponse,
        eventList,
    }
    return out
}

func DefaultEventList(base responseobjs.BaseInfo) EventListResponse {
    return EventList(
        base,
        []obj.Event{},
    )
}
