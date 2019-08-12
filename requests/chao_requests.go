package requests

type EquipChaoRequest struct {
	BasicRequest
	MainChaoID string `json:"mainChaoId"`
	SubChaoID string `json:"subChaorId"`
}
