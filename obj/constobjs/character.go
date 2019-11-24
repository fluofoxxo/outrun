package constobjs

import (
	"strconv"

	"github.com/fluofoxxo/outrun/enums"
	"github.com/fluofoxxo/outrun/obj"
)

/*
All values are placeholders unless otherwise marked (Ex.: Sonic).
This should be changed when real values are found, or if we decide
that having custom values would be better for the balance of the game.

Multiple fields also have no currently known purposes, so these fields
are replaced with numbers that should be very easy to spot as 'abnormal'
in gameplay, thus giving credence to the idea that these values are
being actively used in gameplay. They may also have underlying issues,
which can be detected through a logcat reading.
*/

const NumRedRings = 1337
const PriceRedRings = 9001

// TODO: replace strconv.Itoa conversions to their string equivalents in enums. This should be done after #10 is solved and closed!

var CharacterSonic = obj.Character{
	strconv.Itoa(enums.CharaTypeSonic),
	0,           // unlocked from the start, no cost
	NumRedRings, // ?
	40000,       // used for limit breaking
	16,          // red rings used for limit breaking
}

var CharacterTails = obj.Character{
	strconv.Itoa(enums.CharaTypeTails),
	350,
	NumRedRings,
	30000, // used for limit breaking
	12,    // red rings used for limit breaking
}

var CharacterKnuckles = obj.Character{
	strconv.Itoa(enums.CharaTypeKnuckles),
	350,
	NumRedRings,
	30000, // used for limit breaking
	12,    // red rings used for limit breaking
}

var CharacterAmy = obj.Character{
	strconv.Itoa(enums.CharaTypeAmy),
	400,
	NumRedRings,
	31000, // used for limit breaking
	13,    // red rings used for limit breaking
}

var CharacterShadow = obj.Character{
	strconv.Itoa(enums.CharaTypeShadow),
	500,
	NumRedRings,
	35000, // used for limit breaking
	14,    // red rings used for limit breaking
}

var CharacterBlaze = obj.Character{
	strconv.Itoa(enums.CharaTypeBlaze),
	550,
	NumRedRings,
	39500, // used for limit breaking
	16,    // red rings used for limit breaking
}

var CharacterRouge = obj.Character{
	strconv.Itoa(enums.CharaTypeRouge),
	550,
	NumRedRings,
	39500, // used for limit breaking
	16,    // red rings used for limit breaking
}

var CharacterOmega = obj.Character{
	strconv.Itoa(enums.CharaTypeOmega),
	650,
	NumRedRings,
	46000, // used for limit breaking
	19,    // red rings used for limit breaking
}

var CharacterBig = obj.Character{
	strconv.Itoa(enums.CharaTypeBig),
	700,
	NumRedRings,
	49500, // used for limit breaking
	20,    // red rings used for limit breaking
}

var CharacterCream = obj.Character{
	strconv.Itoa(enums.CharaTypeCream),
	750,
	NumRedRings,
	49500, // used for limit breaking
	20,    // red rings used for limit breaking
}
var CharacterEspio = obj.Character{
	strconv.Itoa(enums.CharaTypeEspio),
	650,
	NumRedRings,
	46000, // used for limit breaking
	19,    // red rings used for limit breaking
}

var CharacterCharmy = obj.Character{
	strconv.Itoa(enums.CharaTypeCharmy),
	650,
	NumRedRings,
	46000, // used for limit breaking
	19,    // red rings used for limit breaking
}

var CharacterVector = obj.Character{
	strconv.Itoa(enums.CharaTypeVector),
	700,
	NumRedRings,
	49500, // used for limit breaking
	20,    // red rings used for limit breaking
}

var CharacterSilver = obj.Character{
	strconv.Itoa(enums.CharaTypeSilver),
	800,
	NumRedRings,
	52500, // used for limit breaking
	22,    // red rings used for limit breaking
}

var CharacterMetalSonic = obj.Character{
	strconv.Itoa(enums.CharaTypeMetalSonic),
	900,
	NumRedRings,
	57000, // used for limit breaking
	24,    // red rings used for limit breaking
}

var CharacterAmitieAmy = obj.Character{
	strconv.Itoa(enums.CharaTypeAmitieAmy),
	77000,
	NumRedRings,
	77000, // used for limit breaking
	32,    // red rings used for limit breaking
}

var CharacterClassicSonic = obj.Character{
	strconv.Itoa(enums.CharaTypeClassicSonic),
	1000,
	NumRedRings,
	67000, // used for limit breaking
	28,    // red rings used for limit breaking
}

var CharacterTikal = obj.Character{
	strconv.Itoa(enums.CharaTypeTikal),
	1100,
	NumRedRings,
	69000, // used for limit breaking
	29,    // red rings used for limit breaking
}

var CharacterGothicAmy = obj.Character{
	strconv.Itoa(enums.CharaTypeGothicAmy),
	91000,
	NumRedRings,
	91000, // used for limit breaking
	38,    // red rings used for limit breaking
}

var CharacterHalloweenShadow = obj.Character{
	strconv.Itoa(enums.CharaTypeHalloweenShadow),
	99000,
	NumRedRings,
	99000, // used for limit breaking
	41,    // red rings used for limit breaking
}

var CharacterHalloweenRouge = obj.Character{
	strconv.Itoa(enums.CharaTypeHalloweenRouge),
	99000,
	NumRedRings,
	99000, // used for limit breaking
	41,    // red rings used for limit breaking
}

var CharacterHalloweenOmega = obj.Character{
	strconv.Itoa(enums.CharaTypeHalloweenOmega),
	99000,
	NumRedRings,
	99000, // used for limit breaking
	41,    // red rings used for limit breaking
}

var CharacterMephiles = obj.Character{
	strconv.Itoa(enums.CharaTypeMephiles),
	1550,
	NumRedRings,
	76000, // used for limit breaking
	31,    // red rings used for limit breaking
}

var CharacterPSISilver = obj.Character{
	strconv.Itoa(enums.CharaTypePSISilver),
	2300,
	NumRedRings,
	98000, // used for limit breaking
	41,    // red rings used for limit breaking
}

var CharacterXMasSonic = obj.Character{
	strconv.Itoa(enums.CharaTypeXMasSonic),
	85000,
	NumRedRings,
	85000, // used for limit breaking
	35,    // red rings used for limit breaking
}

var CharacterXMasTails = obj.Character{
	strconv.Itoa(enums.CharaTypeXMasTails),
	85000,
	NumRedRings,
	85000, // used for limit breaking
	35,    // red rings used for limit breaking
}

var CharacterXMasKnuckles = obj.Character{
	strconv.Itoa(enums.CharaTypeXMasKnuckles),
	85000,
	NumRedRings,
	85000, // used for limit breaking
	35,    // red rings used for limit breaking
}

var CharacterWerehog = obj.Character{
	strconv.Itoa(enums.CharaTypeWerehog),
	800,
	NumRedRings,
	52500, // used for limit breaking
	22,    // red rings used for limit breaking
}

var CharacterSticks = obj.Character{
	strconv.Itoa(enums.CharaTypeSticks),
	750,
	NumRedRings,
	49500, // used for limit breaking
	20,    // red rings used for limit breaking
}
