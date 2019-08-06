package responses

import (
    "github.com/fluofoxxo/outrun/objects"
)

type EventListResponse struct {
    BaseResponse
    EventList []objects.Event `json:"eventList"`
}

func NewEventListResponse(base BaseInfo, eventList []objects.Event) EventListResponse {
    br := NewBaseResponse(base)
    elr := EventListResponse{
        br,
        eventList,
    }
    return elr
}
