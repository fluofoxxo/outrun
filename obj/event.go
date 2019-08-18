package obj

type Event struct {
    ID        int64 `json:"eventId"`        // used to get the type... For some reason...
    Type      int64 `json:"eventType"`      // Dead code in game but still must be satisfied
    StartTime int64 `json:"eventStartTime"` // UTC time
    EndTime   int64 `json:"eventEndTime"`   // UTC time
    CloseTime int64 `json:"eventCloseTime"` // UTC time
}

func NewEvent(id, startTime, endTime, closeTime int64) Event {
    return Event{
        id,
        0,
        startTime,
        endTime,
        closeTime,
    }
}
