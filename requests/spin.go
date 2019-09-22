package requests

type CommitWheelSpinRequest struct {
    Base
    Count int64 `json:"count,string"`
}
