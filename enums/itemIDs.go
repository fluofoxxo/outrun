package enums

import "strconv"

const (
	ItemIDNone = -1
)
const (
	ItemIDBoostScore = int64(iota) + 110000
	ItemIDBoostTrampoline
	ItemIDBoostSubChara
)
const (
	ItemIDInvincible = int64(iota) + 120000
	ItemIDBarrier
	ItemIDMagnet
	ItemIDTrampoline
	ItemIDCombo
	ItemIDLaser
	ItemIDDrill
	ItemIDAsteroid
	ItemIDRingBonus
	ItemIDDistanceBonus
	ItemIDAnimalBonus
)

var (
	ItemIDStrBoostScore      = strconv.Itoa(int(ItemIDBoostScore))
	ItemIDStrBoostTrampoline = strconv.Itoa(int(ItemIDBoostTrampoline))
	ItemIDStrBoostSubChara   = strconv.Itoa(int(ItemIDBoostSubChara))
)

var (
	ItemIDStrInvincible    = strconv.Itoa(int(ItemIDInvincible))
	ItemIDStrBarrier       = strconv.Itoa(int(ItemIDBarrier))
	ItemIDStrMagnet        = strconv.Itoa(int(ItemIDMagnet))
	ItemIDStrTrampoline    = strconv.Itoa(int(ItemIDTrampoline))
	ItemIDStrCombo         = strconv.Itoa(int(ItemIDCombo))
	ItemIDStrLaser         = strconv.Itoa(int(ItemIDLaser))
	ItemIDStrDrill         = strconv.Itoa(int(ItemIDDrill))
	ItemIDStrAsteroid      = strconv.Itoa(int(ItemIDAsteroid))
	ItemIDStrRingBonus     = strconv.Itoa(int(ItemIDRingBonus))
	ItemIDStrDistanceBonus = strconv.Itoa(int(ItemIDDistanceBonus))
	ItemIDStrAnimalBonus   = strconv.Itoa(int(ItemIDAnimalBonus))
)

const (
	ItemIDPackedInvincible0 = int64(iota) + 120100
	ItemIDPackedBarrier0
	ItemIDPackedMagnet0
	ItemIDPackedTrampoline0
	ItemIDPackedCombo0
	ItemIDPackedLaser0
	ItemIDPackedDrill0
	ItemIDPackedAsteroid0
	ItemIDPackedRingBonus0
	ItemIDPackedDistanceBonus0
	ItemIDPackedAnimalBonus0
)
const (
	ItemIDPackedInvincible1 = int64(iota) + 121000
	ItemIDPackedBarrier1
	ItemIDPackedMagnet1
	ItemIDPackedTrampoline1
	ItemIDPackedCombo1
	ItemIDPackedLaser1
	ItemIDPackedDrill1
	ItemIDPackedAsteroid1
	ItemIDPackedRingBonus1
	ItemIDPackedDistanceBonus1
	ItemIDPackedAnimalBonus1
)

const ( // Mileage rewards
	ItemIDRedRing  = 900000
	ItemIDRedRing0 = 900010
	ItemIDRedRing1 = 900030
	ItemIDRedRing2 = 900060
	ItemIDRedRing3 = 900210
	ItemIDRedRing4 = 900380
	ItemIDRing     = 910000
	ItemIDRing0    = 910021
	ItemIDRing1    = 910045
	ItemIDRing2    = 910094
	ItemIDRing3    = 910147
	ItemIDRing4    = 910204
	ItemIDRing5    = 910265
)
