package requests

type EquipChaoRequest struct {
    Base
    MainChaoID string `json:"mainChaoId"`
    SubChaoID  string `json:"subChaoId"`
}

type CommitChaoWheelSpinRequest struct {
    Base
    Count int64 `json:"count,string"`
}
