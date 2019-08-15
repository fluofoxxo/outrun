package netobj

import (
	"github.com/fluofoxxo/outrun/obj"
)

type Character struct { // Can also be used as PlayCharacter
	obj.Character
	Status            int64    `json:"status"` // value from enums.CharacterStatus*
	Level             int64    `json:"level"`
	Exp               int64    `json:"exp"`
	Star              int64    `json:"star"`
	StarMax           int64    `json:"starMax"`
	LockCondition     int64    `json:"lockCondition"` // value from enums.LockCondition*
	CampaignList      []string `json:"campaignList"`
	AbilityLevel      []int64  `json:"abilityLevel"`    // levels for each ability
	AbilityNumRings   []int64  `json:"abilityNumRings"` // ?
	AbilityLevelUp    []int64  `json:"abilityLevelup"`  // this is a list of items using enums.ItemID*
	AbilityLevelUpExp []int64  `json:"abilityLevelupExp"`
}
