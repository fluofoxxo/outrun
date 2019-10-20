package roulette

import (
	"errors"
	"math/rand"
	"strconv"

	"github.com/fluofoxxo/outrun/consts"
	"github.com/fluofoxxo/outrun/enums"
)

type LogicChao struct { // TODO: Find a better way to replace this!
	ID       string `json:"chaoId"`
	Rarity   int64  `json:"rarity"`
	Hidden   int64  `json:"hidden"` // this value is required, but is never really used in game... Best to keep at "1"
	Status   int64  `json:"status"` // enums.ChaoStatus*
	Level    int64  `json:"level"`
	Dealing  int64  `json:"setStatus"` // enums.ChaoDealing*
	Acquired int64  `json:"acquired"`  // flag
}

func GetRandomChaoWheelCharacter(count int) ([]string, error) {
	// Reference: https://eli.thegreenplace.net/2010/01/22/weighted-random-generation-in-python/
	totalWeights := []float64{}
	charList := []string{}
	runningTotal := float64(0)
	for cid, weight := range consts.RandomChaoWheelCharacterPrizes {
		charList = append(charList, cid)
		runningTotal += weight
		totalWeights = append(totalWeights, runningTotal)
	}

	finalCharList := []string{}
	for count > 0 {
		count--
		gotChar := false
		rnd := rand.Float64() * runningTotal
		for i, total := range totalWeights {
			if rnd < total {
				finalCharList = append(finalCharList, charList[i])
				gotChar = true
				break
			}
		}
		if !gotChar {
			return []string{}, errors.New("should not have gotten here")
		}
	}
	return finalCharList, nil
}

func GetRandomChaoWheelChao(rarity int64, count int) ([]string, error) {
	// Reference: https://eli.thegreenplace.net/2010/01/22/weighted-random-generation-in-python/
	totalWeights := []float64{}
	chaoList := []string{}
	runningTotal := float64(0)
	for chid, weight := range consts.RandomChaoWheelChaoPrizes {
		if string(chid[2]) == strconv.Itoa(int(rarity)) {
			chaoList = append(chaoList, chid)
			runningTotal += weight
			totalWeights = append(totalWeights, runningTotal)
		}
	}

	finalChaoList := []string{}
	for count > 0 {
		count--
		gotChao := false
		rnd := rand.Float64() * runningTotal
		for i, total := range totalWeights {
			if rnd < total {
				finalChaoList = append(finalChaoList, chaoList[i])
				gotChao = true
				break
			}
		}
		if !gotChao {
			return []string{}, errors.New("should not have gotten here")
		}
	}
	return finalChaoList, nil
}

func GetRandomChaoRouletteItems(rarities []int64, allowedCharacters, allowedChao []string) ([]string, []int64, error) { // TODO: Possibly rename to GetRandomChaoWheelItems?
	//failsafeTimeout := int64(3) // seconds, used to break in case of infinite loop
	//debugOriginalTime := time.Now().Unix()

	/*dprint := func(a ...interface{}) {
		if config.CFile.DebugPrints {
			log.Println(a...)
		}
	}*/
	allowedRarity := func(x int) bool {
		/*
			if x == 100 { // character
				return len(allowedCharacters) != 0 // if any characters allowed, return true
			}
			searchingForRarity := strconv.Itoa(x)[0] // get byte
			for _, chao := range allowedChao {
				if chao[2] == searchingForRarity {
					return true
				}
			}
			return false
		*/
		return true // Don't check for duplicate Chao
	}
	//chaoAllowed := func(chid string) bool {
	_ = func(chid string) bool {
		// Checks if chid is in allowedChao
		for _, chao := range allowedChao {
			if chao == chid {
				return true
			}
		}
		return false
	}
	//characterAllowed := func(cid string) bool {
	_ = func(cid string) bool {
		// Checks if cid is in allowedCharacters
		/*
			for _, char := range allowedCharacters {
				if char == cid {
					return true
				}
			}
			return false
		*/
		return true // Don't check for duplicate characters
	}

	getUnusedChao := func(rarity int64) (string, error) {
		// Implies that there is a Chao of the given rarity allowed! Check with allowedRarity.
		chao, err := GetRandomChaoWheelChao(rarity, 1)
		// Don't check for unused Chao
		return chao[0], err
		/*
			if err != nil {
				return "", err
			}
			for !chaoAllowed(chao[0]) { // get allowed Chao
				if time.Now().Unix()-failsafeTimeout > debugOriginalTime { // break out of long loop (debug)
					return enums.ChaoIDStrHeroChao, nil
				}
				dprint("searching for allowed chao ", chao[0])
				chao, err = GetRandomChaoWheelChao(rarity, 1)
				if err != nil {
					return "", err
				}
			}
			return chao[0], nil
		*/
	}
	getUnusedCharacter := func() (string, error) {
		// Implies that there are characters available! Check with len(allowedCharacters) != 0 or allowedRarity.
		char, err := GetRandomChaoWheelCharacter(1)
		// Don't check for unused characters
		return char[0], err
		/*
			if err != nil {
				return "", err
			}
			for !characterAllowed(char[0]) {
				if time.Now().Unix()-failsafeTimeout > debugOriginalTime { // break out of long loop (debug)
					return enums.CTStrTails, nil
				}
				dprint("searching for allowed character ", char)
				char, err = GetRandomChaoWheelCharacter(1)
				if err != nil {
					return "", err
				}
			}
			return char[0], nil
		*/
	}

	newRarities := rarities
	items := []string{}
	for i, rarity := range rarities {
		if rarity == 100 { // Character
			if !allowedRarity(100) { // if character not possible
				rarity = 2 // degrade to rarity 2 Chao
			} else {
				char, err := getUnusedCharacter()
				if err != nil {
					return []string{}, []int64{}, err
				}
				items = append(items, char)
			}
		}
		if rarity == 2 { // Rarity 2 Chao
			if !allowedRarity(2) { // if impossible
				rarity = 1 // degrade
			} else {
				chao, err := getUnusedChao(2)
				if err != nil {
					return []string{}, []int64{}, err
				}
				items = append(items, chao)
			}
		}
		if rarity == 1 { // Rarity 1 Chao
			if !allowedRarity(1) { // if impossible
				rarity = 0
			} else {
				chao, err := getUnusedChao(1)
				if err != nil {
					return []string{}, []int64{}, err
				}
				items = append(items, chao)
			}
		}
		if rarity == 0 { // Rarity 1 Chao
			if !allowedRarity(0) { // if impossible
				// TODO: what do we do now?
				items = append(items, enums.ChaoIDStrHeroChao)
			} else {
				chao, err := getUnusedChao(0)
				if err != nil {
					return []string{}, []int64{}, err
				}
				items = append(items, chao)
			}
		}
		newRarities[i] = rarity
	}

	return items, newRarities, nil
}

func ChooseChaoRouletteItem(items []string, weights []int64) (string, error) {
	index, err := ChooseChaoRouletteItemIndex(items, weights)
	if err != nil {
		return "", err
	}
	return items[index], nil
}

func ChooseChaoRouletteItemIndex(items []string, weights []int64) (int, error) {
	// TODO: Remake! This is probably the worst way to do weighted selection!
	if len(items) != len(weights) { // Vibe check
		return -1, errors.New("'items' and 'weights' lengths do not match")
	}
	possibilities := []string{}
	appendXTimes := func(s string, t int64) {
		for t > 0 {
			possibilities = append(possibilities, s)
			t--
		}
	}
	for i, item := range items {
		weight := weights[i]
		appendXTimes(item, weight)
	}
	// select randomly from possibilities
	posIndex := rand.Intn(len(possibilities))
	selection := possibilities[posIndex]
	// return index of selection from items
	for i, item := range items {
		if selection == item {
			return i, nil
		}
	}
	return -1, errors.New("selection not found in 'items'") // This should never, ever happen
}
