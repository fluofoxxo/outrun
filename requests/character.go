package requests

type ChangeCharacterRequest struct {
    Base
    MainCharaID string `json:"mainCharacterId"`
    SubCharaID  string `json:"subCharacterId"`
}

type UpgradeCharacterRequest struct {
    Base
    AbilityID   string `json:"abilityId"`
    CharacterID string `json:"characterId"`
}
