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
	1000,
	NumRedRings,
	1000, // synced with Cost
	PriceRedRings,
}

var CharacterKnuckles = obj.Character{
	strconv.Itoa(enums.CharaTypeKnuckles),
	2500,
	NumRedRings,
	2500, // synced with Cost
	PriceRedRings,
}

var CharacterAmy = obj.Character{
	strconv.Itoa(enums.CharaTypeAmy),
	4000,
	NumRedRings,
	4000, // synced with Cost
	PriceRedRings,
}

var CharacterShadow = obj.Character{
	strconv.Itoa(enums.CharaTypeShadow),
	8000,
	NumRedRings,
	8000, // synced with Cost
	PriceRedRings,
}

var CharacterBlaze = obj.Character{
	strconv.Itoa(enums.CharaTypeBlaze),
	12500,
	NumRedRings,
	12500, // synced with Cost
	PriceRedRings,
}

var CharacterRouge = obj.Character{
	strconv.Itoa(enums.CharaTypeRouge),
	15000,
	NumRedRings,
	15000, // synced with Cost
	PriceRedRings,
}

var CharacterOmega = obj.Character{
	strconv.Itoa(enums.CharaTypeOmega),
	18500,
	NumRedRings,
	18500, // synced with Cost
	PriceRedRings,
}

var CharacterBig = obj.Character{
	strconv.Itoa(enums.CharaTypeBig),
	22500,
	NumRedRings,
	22500, // synced with Cost
	PriceRedRings,
}

var CharacterCream = obj.Character{
	strconv.Itoa(enums.CharaTypeCream),
	28000,
	NumRedRings,
	28000, // synced with Cost
	PriceRedRings,
}
var CharacterEspio = obj.Character{
	strconv.Itoa(enums.CharaTypeEspio),
	35000,
	NumRedRings,
	35000, // synced with Cost
	PriceRedRings,
}

var CharacterCharmy = obj.Character{
	strconv.Itoa(enums.CharaTypeCharmy),
	41500,
	NumRedRings,
	41500, // synced with Cost
	PriceRedRings,
}

var CharacterVector = obj.Character{
	strconv.Itoa(enums.CharaTypeVector),
	47500,
	NumRedRings,
	47500, // synced with Cost
	PriceRedRings,
}

var CharacterSilver = obj.Character{
	strconv.Itoa(enums.CharaTypeSilver),
	53500,
	NumRedRings,
	53500, // synced with Cost
	PriceRedRings,
}

var CharacterMetalSonic = obj.Character{
	strconv.Itoa(enums.CharaTypeMetalSonic),
	60000,
	NumRedRings,
	60000, // synced with Cost
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
	78500,
	NumRedRings,
	78500, // synced with Cost
	PriceRedRings,
}

var CharacterTikal = obj.Character{
	strconv.Itoa(enums.CharaTypeTikal),
	83000,
	NumRedRings,
	83000, // synced with Cost
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
	125000,
	NumRedRings,
	125000, // synced with Cost
	PriceRedRings,
}

var CharacterPSISilver = obj.Character{
	strconv.Itoa(enums.CharaTypePSISilver),
	185000,
	NumRedRings,
	185000, // synced with Cost
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
	96500,
	NumRedRings,
	96500, // synced with Cost
	PriceRedRings,
}

var CharacterSticks = obj.Character{
	strconv.Itoa(enums.CharaTypeSticks),
	110000,
	NumRedRings,
	110000, // synced with Cost
	PriceRedRings,
}
