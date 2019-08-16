package netobj

import (
	"github.com/fluofoxxo/outrun/enums"
	"github.com/fluofoxxo/outrun/obj"
	"github.com/fluofoxxo/outrun/obj/constobjs"
)

/*
Notes:
  - I believe stars are used as "prestige" for the characters, if all skills are maxed out
    - starMax may be the max prestige
*/

type Character struct { // Can also be used as PlayCharacter
	obj.Character
	Status        int64          `json:"status"` // value from enums.CharacterStatus*
	Level         int64          `json:"level"`
	Exp           int64          `json:"exp"`
	Star          int64          `json:"star"`
	StarMax       int64          `json:"starMax"`
	LockCondition int64          `json:"lockCondition"` // value from enums.LockCondition*
	CampaignList  []obj.Campaign `json:"campaignList"`
	AbilityLevel  []int64        `json:"abilityLevel"` // levels for each ability
	//AbilityNumRings []int64 `json:"abilityNumRings"`  // this was in the original code, but isn't in the game code. this may be a point of failure
	AbilityLevelUp    []int64 `json:"abilityLevelup"` // this is a list of items using enums.ItemID*
	AbilityLevelUpExp []int64 `json:"abilityLevelupExp,omitempty"`
}

func DefaultCharacter(char obj.Character) Character {
	status := int64(enums.CharacterStatusUnlocked)
	level := int64(1)
	exp := int64(31415)
	star := int64(2)       // star value of the character? Differs based on rarity, I guess
	starMax := int64(1331) // ????
	lockCondition := int64(enums.LockConditionOpen)
	campaignList := []obj.Campaign{}
	abilityLevel := []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // 11 abilities?
	abilityLevelUp := []int64{enums.ItemIDInvincible}
	abilityLevelUpExp := []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	return Character{
		char,
		status,
		level,
		exp,
		star,
		starMax,
		lockCondition,
		campaignList,
		abilityLevel,
		abilityLevelUp,
		abilityLevelUpExp,
	}
}

func DefaultCharacterState() []Character { // every character
	return []Character{
		DefaultCharacter(constobjs.CharacterSonic),
		DefaultCharacter(constobjs.CharacterTails),
		DefaultCharacter(constobjs.CharacterKnuckles),
		DefaultCharacter(constobjs.CharacterAmy),
		DefaultCharacter(constobjs.CharacterShadow),
		DefaultCharacter(constobjs.CharacterBlaze),
		DefaultCharacter(constobjs.CharacterRouge),
		DefaultCharacter(constobjs.CharacterOmega),
		DefaultCharacter(constobjs.CharacterBig),
		DefaultCharacter(constobjs.CharacterCream),
		DefaultCharacter(constobjs.CharacterEspio),
		DefaultCharacter(constobjs.CharacterCharmy),
		DefaultCharacter(constobjs.CharacterVector),
		DefaultCharacter(constobjs.CharacterSilver),
		DefaultCharacter(constobjs.CharacterMetalSonic),
		DefaultCharacter(constobjs.CharacterAmitieAmy),
		DefaultCharacter(constobjs.CharacterClassicSonic),
		DefaultCharacter(constobjs.CharacterTikal),
		DefaultCharacter(constobjs.CharacterGothicAmy),
		DefaultCharacter(constobjs.CharacterHalloweenShadow),
		DefaultCharacter(constobjs.CharacterHalloweenRouge),
		DefaultCharacter(constobjs.CharacterHalloweenOmega),
		DefaultCharacter(constobjs.CharacterMephiles),
		DefaultCharacter(constobjs.CharacterPSISilver),
		DefaultCharacter(constobjs.CharacterXMasSonic),
		DefaultCharacter(constobjs.CharacterXMasTails),
		DefaultCharacter(constobjs.CharacterXMasKnuckles),
		DefaultCharacter(constobjs.CharacterWerehog),
		DefaultCharacter(constobjs.CharacterSticks),
	}
}
