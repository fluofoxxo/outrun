package gameconf

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/fluofoxxo/outrun/enums"
)

var CharacterMap = map[string]string{
	"sonic":           enums.CTStrSonic,
	"tails":           enums.CTStrTails,
	"knuckles":        enums.CTStrTails,
	"amy":             enums.CTStrAmy,
	"shadow":          enums.CTStrShadow,
	"blaze":           enums.CTStrBlaze,
	"rouge":           enums.CTStrRouge,
	"omega":           enums.CTStrOmega,
	"big":             enums.CTStrBig,
	"cream":           enums.CTStrCream,
	"espio":           enums.CTStrEspio,
	"charmy":          enums.CTStrCharmy,
	"vector":          enums.CTStrVector,
	"silver":          enums.CTStrSilver,
	"metalSonic":      enums.CTStrMetalSonic,
	"classicSonic":    enums.CTStrClassicSonic,
	"werehog":         enums.CTStrWerehog,
	"sticks":          enums.CTStrSticks,
	"tikal":           enums.CTStrTikal,
	"mephiles":        enums.CTStrMephiles,
	"psiSilver":       enums.CTStrPSISilver,
	"espSilver":       enums.CTStrPSISilver,
	"amitieAmy":       enums.CTStrAmitieAmy,
	"gothicAmy":       enums.CTStrGothicAmy,
	"halloweenShadow": enums.CTStrHalloweenShadow,
	"halloweenRouge":  enums.CTStrHalloweenRouge,
	"halloweenOmega":  enums.CTStrHalloweenOmega,
	"xmasSonic":       enums.CTStrXMasSonic,
	"xmasTails":       enums.CTStrXMasTails,
	"xmasKnuckles":    enums.CTStrXMasKnuckles,
	"empty":           "-1",
	"none":            "-1",
}

// defaults
var Defaults = map[string]interface{}{
	"DAllCharactersUnlocked": true,
	"DDefaultMainCharacter":  "sonic",
	"DDefaultSubCharacter":   "empty",
	"DStartingRings":         int64(5000),
	"DStartingRedRings":      int64(25),
	"DStartingEnergy":        int64(5),
}

var CFile ConfigFile

type ConfigFile struct {
	AllCharactersUnlocked bool   `json:"allCharactersUnlocked,omitempty"`
	DefaultMainCharacter  string `json:"defaultMainCharacter,omitempty"`
	DefaultSubCharacter   string `json:"defaultSubCharacter,omitempty"`
	StartingRings         int64  `json:"startingRings,omitempty"`
	StartingRedRings      int64  `json:"startingRedRings,omitempty"`
	StartingEnergy        int64  `json:"startingEnergy,omitempty"`
}

func Parse(filename string) error {
	CFile = ConfigFile{
		Defaults["DAllCharactersUnlocked"].(bool),
		Defaults["DDefaultMainCharacter"].(string),
		Defaults["DDefaultSubCharacter"].(string),
		Defaults["DStartingRings"].(int64),
		Defaults["DStartingRedRings"].(int64),
		Defaults["DStartingEnergy"].(int64),
	}
	file, err := loadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &CFile)
	if err != nil {
		return err
	}
	var ok bool
	CFile.DefaultMainCharacter, ok = CharacterMap[CFile.DefaultMainCharacter]
	if !ok {
		log.Printf("[WARN] Invalid main character '%s', defaulting to Sonic\n", CFile.DefaultMainCharacter)
		CFile.DefaultMainCharacter = CharacterMap["sonic"]
	}
	CFile.DefaultSubCharacter, ok = CharacterMap[CFile.DefaultSubCharacter]
	if !ok {
		log.Printf("[WARN] Invalid sub character '%s', defaulting to None\n", CFile.DefaultSubCharacter)
		CFile.DefaultSubCharacter = CharacterMap["none"]
	}
	return nil
}

func loadFile(filename string) ([]byte, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return b, err
}
