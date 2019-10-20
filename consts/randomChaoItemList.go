package consts

import "github.com/fluofoxxo/outrun/enums"

type PrizeInfo struct {
	AppearanceChance float64 // % chance for it to be chosen to be in wheel by the server
	Type             int64   // 0 for Chao, 1 for Character
}

// A 'load' as depicted below is the chance for the server to pick
// the associated item, where chosen is if randFloat(0, 100) < load.
// IMPORTANT: This load is exclusive to the rarity of the Chao that
// is being chosen by the server.

var RandomChaoWheelCharacterPrizes = map[string]float64{
	// characterID: load
	// Hopefully this should sum up to 100 just for
	// simplicity, but it shouldn't be a requirement.
	enums.CTStrSonic:        10.0,
	enums.CTStrTails:        10.0,
	enums.CTStrKnuckles:     10.0,
	enums.CTStrAmy:          10.0,
	enums.CTStrBig:          6.0,
	enums.CTStrBlaze:        6.0,
	enums.CTStrCharmy:       4.0,
	enums.CTStrCream:        6.0,
	enums.CTStrEspio:        4.0,
	enums.CTStrMephiles:     1.0,
	enums.CTStrOmega:        2.0,
	enums.CTStrPSISilver:    0.1, // This puts the total up to 100.1, but... [shrug]
	enums.CTStrRouge:        4.0,
	enums.CTStrShadow:       4.0,
	enums.CTStrSilver:       4.0,
	enums.CTStrSticks:       6.0,
	enums.CTStrTikal:        2.0,
	enums.CTStrVector:       4.0,
	enums.CTStrWerehog:      2.0,
	enums.CTStrClassicSonic: 2.0,
	enums.CTStrMetalSonic:   3.0,
}

var RandomChaoWheelChaoPrizes = map[string]float64{
	// TODO: Balance these
	enums.ChaoIDStrHeroChao:             5.0,
	enums.ChaoIDStrGoldChao:             5.0,
	enums.ChaoIDStrDarkChao:             5.0,
	enums.ChaoIDStrJewelChao:            4.5,
	enums.ChaoIDStrNormalChao:           4.5,
	enums.ChaoIDStrOmochao:              4.5,
	enums.ChaoIDStrRCMonkey:             3.5,
	enums.ChaoIDStrRCSpring:             3.5,
	enums.ChaoIDStrRCElectromagnet:      3.5,
	enums.ChaoIDStrBabyCyanWisp:         3.5,
	enums.ChaoIDStrBabyIndigoWisp:       3.5,
	enums.ChaoIDStrBabyYellowWisp:       3.5,
	enums.ChaoIDStrRCPinwheel:           3.0,
	enums.ChaoIDStrRCPiggyBank:          3.0,
	enums.ChaoIDStrRCBalloon:            3.0,
	enums.ChaoIDStrEasterChao:           2.0,
	enums.ChaoIDStrPurplePapurisu:       0.0, // Event (Puyo Puyo)
	enums.ChaoIDStrMagLv1:               0.0, // Event (Phantasy Star Online 2)
	enums.ChaoIDStrEggChao:              3.5,
	enums.ChaoIDStrPumpkinChao:          3.5,
	enums.ChaoIDStrSkullChao:            3.0,
	enums.ChaoIDStrYacker:               2.0,
	enums.ChaoIDStrRCGoldenPiggyBank:    1.5,
	enums.ChaoIDStrWizardChao:           1.0,
	enums.ChaoIDStrRCTurtle:             1.0,
	enums.ChaoIDStrRCUFO:                1.0,
	enums.ChaoIDStrRCBomber:             1.0,
	enums.ChaoIDStrEasterBunny:          1.0,
	enums.ChaoIDStrMagicLamp:            1.0,
	enums.ChaoIDStrStarShapedMissile:    1.0,
	enums.ChaoIDStrSuketoudara:          1.0, // Event (Puyo Puyo)
	enums.ChaoIDStrRappy:                1.0, // Event (Phantasy Star Online 2)
	enums.ChaoIDStrBlowfishTransporter:  1.0,
	enums.ChaoIDStrGenesis:              1.0,
	enums.ChaoIDStrCartridge:            1.0,
	enums.ChaoIDStrRCFighter:            1.0,
	enums.ChaoIDStrRCHovercraft:         1.0,
	enums.ChaoIDStrRCHelicopter:         1.0,
	enums.ChaoIDStrGreenCrystalMonsterS: 1.0,
	enums.ChaoIDStrGreenCrystalMonsterL: 1.0,
	enums.ChaoIDStrRCAirship:            1.0,
	enums.ChaoIDStrDesertChao:           1.0,
	enums.ChaoIDStrRCSatellite:          1.0,
	enums.ChaoIDStrMarineChao:           1.0,
	enums.ChaoIDStrNightopian:           1.0, // Event (NiGHTS)
	enums.ChaoIDStrOrca:                 1.0,
	enums.ChaoIDStrSonicOmochao:         1.0,
	enums.ChaoIDStrTailsOmochao:         1.0,
	enums.ChaoIDStrKnucklesOmochao:      1.0,
	enums.ChaoIDStrBoo:                  1.0,
	enums.ChaoIDStrHalloweenChao:        1.0,
	enums.ChaoIDStrHeavyBomb:            1.0, // Event (Fantasy Zone)
	enums.ChaoIDStrBlockBomb:            1.0,
	enums.ChaoIDStrHunkofMeat:           1.0,
	enums.ChaoIDStrYeti:                 1.0, // Event (Christmas)
	enums.ChaoIDStrSnowChao:             1.0, // Event (Christmas)
	enums.ChaoIDStrIdeya:                1.0, // Event (Christmas NiGHTS)
	enums.ChaoIDStrChristmasNightopian:  1.0, // Event (Christmas NiGHTS)
	enums.ChaoIDStrOrbot:                1.0,
	enums.ChaoIDStrCubot:                1.0,
	enums.ChaoIDStrLightChaos:           1.0,
	enums.ChaoIDStrHeroChaos:            1.0,
	enums.ChaoIDStrDarkChaos:            1.0,
	enums.ChaoIDStrChip:                 1.0,
	enums.ChaoIDStrShahra:               1.0,
	enums.ChaoIDStrCaliburn:             1.0,
	enums.ChaoIDStrKingArthursGhost:     1.0,
	enums.ChaoIDStrRCTornado:            1.0,
	enums.ChaoIDStrRCBattleCruiser:      1.0,
	enums.ChaoIDStrMerlina:              1.0, // Event (Windy Hill in Spring)
	enums.ChaoIDStrErazorDjinn:          1.0, // Event (Desert Ruins)
	enums.ChaoIDStrRCMoonMech:           1.0, // Raid Boss Roulette(?)
	enums.ChaoIDStrCarbuncle:            1.0, // Event (Puyo Puyo)
	enums.ChaoIDStrKuna:                 1.0, // Event (Phantasy Star Online 2)
	enums.ChaoIDStrChaos:                1.0,
	enums.ChaoIDStrDeathEgg:             1.0,
	enums.ChaoIDStrRedCrystalMonsterS:   1.0,
	enums.ChaoIDStrRedCrystalMonsterL:   1.0,
	enums.ChaoIDStrGoldenGoose:          1.0,
	enums.ChaoIDStrMotherWisp:           1.0, // Event (Tropical Coast)
	enums.ChaoIDStrRCPirateSpaceship:    1.0,
	enums.ChaoIDStrGoldenAngel:          1.0,
	enums.ChaoIDStrNiGHTS:               1.0, // Event (NiGHTS)
	enums.ChaoIDStrReala:                1.0, // Event (NiGHTS)
	enums.ChaoIDStrRCTornado2:           1.0,
	enums.ChaoIDStrChaoWalker:           1.0,
	enums.ChaoIDStrDarkQueen:            1.0,
	enums.ChaoIDStrKingBoomBoo:          1.0, // Event (Halloween)
	enums.ChaoIDStrOPapa:                1.0, // Event (Fantasy Zone)
	enums.ChaoIDStrOpaOpa:               1.0, // Event (Fantasy Zone)
	enums.ChaoIDStrRCBlockFace:          1.0,
	enums.ChaoIDStrChristmasYeti:        1.0, // Event (Christmas)
	enums.ChaoIDStrChristmasNiGHTS:      1.0, // Event (Christmas NiGHTS)
	//enums.ChaoIDStrDFekt:                1.0,
	enums.ChaoIDStrDarkChaoWalker: 1.0,
}
