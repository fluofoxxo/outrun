package requests

type ChangeCharacterRequest struct {
	BasicRequest
	MainCharacterID string `json:"mainCharacterId"`
	SubCharacterID string `json:"subCharacterId"`
}
