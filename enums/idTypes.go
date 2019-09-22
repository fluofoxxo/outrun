package enums

// Game divides these numbers by 10000 in order to determine
// the kind of item that the server sends its way.

const (
	IDTypeNone = -1
)
const (
	IDTypeBoostItem = iota + 110000
	IDTypeEquipItem
)
const (
	IDTypeItemRouletteWin = iota + 200000
	IDTypeRouletteToken
	IDTypeEggItem
	IDTypePremiumRouletteTicket
	IDTypeItemRouletteTicket
)
const (
	IDTypeChara = 300000
	IDTypeChao  = 400000
)
const (
	IDTypeRedRing = iota + 900000
	IDTypeRing
	IDTypeEnergy
	IDTypeEnergyMax
)
const (
	IDTypeRaidRing = 960000
)
