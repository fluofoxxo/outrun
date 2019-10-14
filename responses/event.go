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
        []obj.Event{
            /*
               obj.NewEvent(
                   //enums.EventIDSpecialStage+10002, // game subtracts one from number?
                   //enums.EventIDAdvert+50002, // 50002 converts to ui_event_50005_Atlas_en?
                   //enums.EventIDBGM+70002, // 70002 goes to 70007
                   enums.EventIDQuick+60002, // 60002 goes to 60006
                   0,                        // event type
                   now.BeginningOfDay().Unix(),
                   now.EndOfDay().Unix(),
                   now.EndOfDay().Unix(),
               ),
            */
        },
    )
}
