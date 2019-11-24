package consts

import (
	"math/rand"
	"strconv"

	"github.com/fluofoxxo/outrun/enums"
)

type AmountRange struct {
	Min  int64
	Max  int64
	Step int64
}

func (a AmountRange) GetRandom() int64 {
	// construct random list first
	randomSelections := []int64{}
	diff := int64(0)
	currMin := a.Min
	for diff >= 0 {
		randomSelections = append(randomSelections, currMin)
		currMin += a.Step
		diff = a.Max - currMin
	}
	selectionIndex := rand.Intn(len(randomSelections))
	selection := randomSelections[selectionIndex]
	return selection
}

// The game does not support RingBonus, DistanceBonus,
// or AnimalBonus on the normal wheel.

// NOTE: If you remove an item from NormalWheelItemAmountRange
// but don't remove it from RandomItemListNormalWheel, you're going
// to create a memory leak.

var RandomItemListNormalWheel = []string{
	enums.ItemIDStrInvincible,
	enums.ItemIDStrBarrier,
	enums.ItemIDStrMagnet,
	enums.ItemIDStrTrampoline,
	enums.ItemIDStrCombo,
	enums.ItemIDStrLaser,
	enums.ItemIDStrDrill,
	enums.ItemIDStrAsteroid,
	strconv.Itoa(enums.IDTypeRedRing),
	//strconv.Itoa(enums.IDTypeItemRouletteWin),
}

var NormalWheelItemAmountRange = map[string]AmountRange{
	enums.ItemIDStrInvincible:         AmountRange{5, 10, 1},
	enums.ItemIDStrBarrier:            AmountRange{5, 10, 1},
	enums.ItemIDStrMagnet:             AmountRange{5, 10, 1},
	enums.ItemIDStrTrampoline:         AmountRange{5, 10, 1},
	enums.ItemIDStrCombo:              AmountRange{5, 10, 1},
	enums.ItemIDStrLaser:              AmountRange{5, 10, 1},
	enums.ItemIDStrDrill:              AmountRange{5, 10, 1},
	enums.ItemIDStrAsteroid:           AmountRange{5, 10, 1},
	strconv.Itoa(enums.IDTypeRedRing): AmountRange{20, 100, 20},
	//strconv.Itoa(enums.IDTypeItemRouletteWin): AmountRange{1, 1, 1},
}
