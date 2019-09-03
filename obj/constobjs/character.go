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
	0,             // unlocked from the start, no cost
	NumRedRings,   // ?
	0,             // synced with Cost
	PriceRedRings, // ?
}

var CharacterTails = obj.Character{
	strconv.Itoa(enums.CharaTypeTails),
	350,
	NumRedRings,
	350, // synced with Cost
	PriceRedRings,
}

var CharacterKnuckles = obj.Character{
	strconv.Itoa(enums.CharaTypeKnuckles),
	350,
	NumRedRings,
	350, // synced with Cost
	PriceRedRings,
}

var CharacterAmy = obj.Character{
	strconv.Itoa(enums.CharaTypeAmy),
	400,
	NumRedRings,
	400, // synced with Cost
	PriceRedRings,
}

var CharacterShadow = obj.Character{
	strconv.Itoa(enums.CharaTypeShadow),
	500,
	NumRedRings,
	500, // synced with Cost
	PriceRedRings,
}

var CharacterBlaze = obj.Character{
	strconv.Itoa(enums.CharaTypeBlaze),
	550,
	NumRedRings,
	550, // synced with Cost
	PriceRedRings,
}

var CharacterRouge = obj.Character{
	strconv.Itoa(enums.CharaTypeRouge),
	550,
	NumRedRings,
	550, // synced with Cost
	PriceRedRings,
}

var CharacterOmega = obj.Character{
	strconv.Itoa(enums.CharaTypeOmega),
	650,
	NumRedRings,
	650, // synced with Cost
	PriceRedRings,
}

var CharacterBig = obj.Character{
	strconv.Itoa(enums.CharaTypeBig),
	700,
	NumRedRings,
	700, // synced with Cost
	PriceRedRings,
}

var CharacterCream = obj.Character{
	strconv.Itoa(enums.CharaTypeCream),
	750,
	NumRedRings,
	750, // synced with Cost
	PriceRedRings,
}
var CharacterEspio = obj.Character{
	strconv.Itoa(enums.CharaTypeEspio),
	650,
	NumRedRings,
	650, // synced with Cost
	PriceRedRings,
}

var CharacterCharmy = obj.Character{
	strconv.Itoa(enums.CharaTypeCharmy),
	650,
	NumRedRings,
	650, // synced with Cost
	PriceRedRings,
}

var CharacterVector = obj.Character{
	strconv.Itoa(enums.CharaTypeVector),
	700,
	NumRedRings,
	700, // synced with Cost
	PriceRedRings,
}

var CharacterSilver = obj.Character{
	strconv.Itoa(enums.CharaTypeSilver),
	800,
	NumRedRings,
	800, // synced with Cost
	PriceRedRings,
}

var CharacterMetalSonic = obj.Character{
	strconv.Itoa(enums.CharaTypeMetalSonic),
	900,
	NumRedRings,
	900, // synced with Cost
	PriceRedRings,
}

var CharacterAmitieAmy = obj.Character{
	strconv.Itoa(enums.CharaTypeAmitieAmy),
	77000,
	NumRedRings,
	77000, // synced with Cost
	PriceRedRings,
}

var CharacterClassicSonic = obj.Character{
	strconv.Itoa(enums.CharaTypeClassicSonic),
	1000,
	NumRedRings,
	1000, // synced with Cost
	PriceRedRings,
}

var CharacterTikal = obj.Character{
	strconv.Itoa(enums.CharaTypeTikal),
	1100,
	NumRedRings,
	1100, // synced with Cost
	PriceRedRings,
}

var CharacterGothicAmy = obj.Character{
	strconv.Itoa(enums.CharaTypeGothicAmy),
	91000,
	NumRedRings,
	91000, // synced with Cost
	PriceRedRings,
}

var CharacterHalloweenShadow = obj.Character{
	strconv.Itoa(enums.CharaTypeHalloweenShadow),
	99000,
	NumRedRings,
	99000, // synced with Cost
	PriceRedRings,
}

var CharacterHalloweenRouge = obj.Character{
	strconv.Itoa(enums.CharaTypeHalloweenRouge),
	99000,
	NumRedRings,
	99000, // synced with Cost
	PriceRedRings,
}

var CharacterHalloweenOmega = obj.Character{
	strconv.Itoa(enums.CharaTypeHalloweenOmega),
	99000,
	NumRedRings,
	99000, // synced with Cost
	PriceRedRings,
}

var CharacterMephiles = obj.Character{
	strconv.Itoa(enums.CharaTypeMephiles),
	1550,
	NumRedRings,
	1550, // synced with Cost
	PriceRedRings,
}

var CharacterPSISilver = obj.Character{
	strconv.Itoa(enums.CharaTypePSISilver),
	2300,
	NumRedRings,
	2300, // synced with Cost
	PriceRedRings,
}

var CharacterXMasSonic = obj.Character{
	strconv.Itoa(enums.CharaTypeXMasSonic),
	85000,
	NumRedRings,
	85000, // synced with Cost
	PriceRedRings,
}

var CharacterXMasTails = obj.Character{
	strconv.Itoa(enums.CharaTypeXMasTails),
	85000,
	NumRedRings,
	85000, // synced with Cost
	PriceRedRings,
}

var CharacterXMasKnuckles = obj.Character{
	strconv.Itoa(enums.CharaTypeXMasKnuckles),
	85000,
	NumRedRings,
	85000, // synced with Cost
	PriceRedRings,
}

var CharacterWerehog = obj.Character{
	strconv.Itoa(enums.CharaTypeWerehog),
	800,
	NumRedRings,
	800, // synced with Cost
	PriceRedRings,
}

var CharacterSticks = obj.Character{
	strconv.Itoa(enums.CharaTypeSticks),
	750,
	NumRedRings,
	750, // synced with Cost
	PriceRedRings,
}
