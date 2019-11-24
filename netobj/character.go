package netobj

import (
	"github.com/fluofoxxo/outrun/config/gameconf"
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
	Status            int64          `json:"status"` // value from enums.CharacterStatus*
	Level             int64          `json:"level"`
	Exp               int64          `json:"exp"`
	Star              int64          `json:"star"`
	StarMax           int64          `json:"starMax"`
	LockCondition     int64          `json:"lockCondition"` // value from enums.LockCondition*
	CampaignList      []obj.Campaign `json:"campaignList"`
	AbilityLevel      []int64        `json:"abilityLevel"`    // levels for each ability
	AbilityNumRings   []int64        `json:"abilityNumRings"` // where is this being checked? I can't find the string using dnSpy...
	AbilityLevelUp    []int64        `json:"abilityLevelup"`  // this is a list of items using enums.ItemID*
	AbilityLevelUpExp []int64        `json:"abilityLevelupExp,omitempty"`
}

var tick = 0

func DefaultCharacter(char obj.Character) Character {
	status := int64(enums.CharacterStatusUnlocked)
	level := int64(0)
	exp := int64(0)
	star := int64(0)     // Limit breaks
	starMax := int64(10) // Max number of limit breaks?
	lockCondition := int64(enums.LockConditionOpen)
	campaignList := []obj.Campaign{}
	abilityLevel := []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // 11 abilities?
	abilityNumRings := []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
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
		abilityNumRings,
		abilityLevelUp,
		abilityLevelUpExp,
	}
}

func DefaultLockedCharacter(char obj.Character) Character {
	ch := DefaultCharacter(char)
	ch.LockCondition = int64(enums.LockConditionRingOrRedRing)
	ch.Status = int64(enums.CharacterStatusLocked)
	return ch
}

func DefaultStageLockedCharacter(char obj.Character) Character {
	ch := DefaultCharacter(char)
	ch.LockCondition = int64(enums.LockConditionMileageEpisode)
	ch.Status = int64(enums.CharacterStatusLocked)
	return ch
}

func DefaultRouletteLockedCharacter(char obj.Character) Character {
	ch := DefaultCharacter(char)
	ch.LockCondition = int64(enums.LockConditionRoulette)
	ch.Status = int64(enums.CharacterStatusLocked)
	return ch
}

func UnlockedCharacterState() []Character { // every character
	// TODO: It looks like the game only wants 300000-300020, otherwise an index error is created. Investigate this in game!
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
		DefaultCharacter(constobjs.CharacterClassicSonic),
		DefaultCharacter(constobjs.CharacterWerehog),
		DefaultCharacter(constobjs.CharacterSticks),
		DefaultCharacter(constobjs.CharacterTikal),
		DefaultCharacter(constobjs.CharacterMephiles),
		DefaultCharacter(constobjs.CharacterPSISilver),
	}
}

func DefaultCharacterState() []Character {
	if gameconf.CFile.AllCharactersUnlocked {
		return UnlockedCharacterState()
	}
	return []Character{
		DefaultCharacter(constobjs.CharacterSonic),
		DefaultStageLockedCharacter(constobjs.CharacterTails),    // Episode 11
		DefaultStageLockedCharacter(constobjs.CharacterKnuckles), // 'Campaign Map 17'
		DefaultLockedCharacter(constobjs.CharacterAmy),
		DefaultLockedCharacter(constobjs.CharacterShadow),
		DefaultLockedCharacter(constobjs.CharacterBlaze),
		DefaultLockedCharacter(constobjs.CharacterRouge),
		DefaultLockedCharacter(constobjs.CharacterOmega),
		DefaultLockedCharacter(constobjs.CharacterBig),
		DefaultLockedCharacter(constobjs.CharacterCream),
		DefaultLockedCharacter(constobjs.CharacterEspio),
		DefaultLockedCharacter(constobjs.CharacterCharmy),
		DefaultLockedCharacter(constobjs.CharacterVector),
		DefaultLockedCharacter(constobjs.CharacterSilver),
		DefaultLockedCharacter(constobjs.CharacterMetalSonic),
		DefaultLockedCharacter(constobjs.CharacterClassicSonic),
		DefaultLockedCharacter(constobjs.CharacterWerehog),
		DefaultLockedCharacter(constobjs.CharacterSticks),
		DefaultLockedCharacter(constobjs.CharacterTikal),
		DefaultLockedCharacter(constobjs.CharacterMephiles),
		DefaultLockedCharacter(constobjs.CharacterPSISilver),
	}
}
