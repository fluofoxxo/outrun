package requests

type ChangeCharacterRequest struct {
	Base
	MainCharaID string `json:"mainCharacterId"`
	SubCharaID  string `json:"subCharacterId"`
}
