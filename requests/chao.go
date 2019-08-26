package requests

type EquipChaoRequest struct {
	Base
	MainChaoID string `json:"mainChaoId"`
	SubChaoID  string `json:"subChaoId"`
}
