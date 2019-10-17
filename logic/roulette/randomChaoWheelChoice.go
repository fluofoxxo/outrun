package roulette

import (
	"errors"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/fluofoxxo/outrun/config"
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

func GetRandomChaoWheelCharacter(count int) []string {
	// Make weighted random pull of a character using an... Unconventional method!
	polls := [1000]string{} // TODO: parameterize
	pollIndex := 0
	for pollIndex < 1000 {
		for cid, load := range consts.RandomChaoWheelCharacterPrizes {
			rint := rand.Intn(1000000 + 1) // + 1 to make inclusive; high number for precision
			chosen := (float64(rint) / 10000) <= load
			if chosen {
				polls[pollIndex] = cid
				pollIndex++
				if pollIndex >= 1000 { // breached bounds
					break
				}
			}
		}
	}
	pollChoices := []string{}
	for count > 0 {
		count--
		pollChoices = append(pollChoices, polls[rand.Intn(len(polls))])
	}
	return pollChoices
}

func GetRandomChaoWheelChao(rarity int64, count int) ([]string, error) {
	polls := [1000]string{} // TODO: parameterize
	pollIndex := 0
	for pollIndex < 1000 {
		for chid, load := range consts.RandomChaoWheelChaoPrizes {
			rint := rand.Intn(1000000 + 1)                   // + 1 to make inclusive; high number for precision
			chaoRarity, err := strconv.Atoi(string(chid[2])) // third digit is rarity
			if err != nil {
				return []string{}, err
			}
			chosen := (float64(rint)/10000) <= load && int64(chaoRarity) == rarity
			if chosen {
				polls[pollIndex] = chid
				pollIndex++
				if pollIndex >= 1000 { // breached bounds
					break
				}
			}
		}
	}
	pollChoices := []string{}
	for count > 0 {
		count--
		pollChoices = append(pollChoices, polls[rand.Intn(len(polls))])
	}
	return pollChoices, nil
}

func GetRandomChaoRouletteItems(rarities []int64, allowedCharacters, allowedChao []string) ([]string, []int64, error) { // TODO: Possibly rename to GetRandomChaoWheelItems?
	failsafeTimeout := int64(3) // seconds, used to break in case of infinite loop
	debugOriginalTime := time.Now().Unix()

	dprint := func(a ...interface{}) {
		if config.CFile.DebugPrints {
			log.Println(a...)
		}
	}
	allowedRarity := func(x int) bool {
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
	}
	chaoAllowed := func(chid string) bool {
		// Checks if chid is in allowedChao
		for _, chao := range allowedChao {
			if chao == chid {
				return true
			}
		}
		return false
	}
	characterAllowed := func(cid string) bool {
		// Checks if cid is in allowedCharacters
		for _, char := range allowedCharacters {
			if char == cid {
				return true
			}
		}
		return false
	}

	getUnusedChao := func(rarity int64) (string, error) {
		// Implies that there is a Chao of the given rarity allowed! Check with allowedRarity.
		chao, err := GetRandomChaoWheelChao(rarity, 1)
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
	}
	getUnusedCharacter := func() string {
		// Implies that there are characters available! Check with len(allowedCharacters) != 0 or allowedRarity.
		char := GetRandomChaoWheelCharacter(1)[0]
		for !characterAllowed(char) {
			if time.Now().Unix()-failsafeTimeout > debugOriginalTime { // break out of long loop (debug)
				return enums.CTStrTails
			}
			dprint("searching for allowed character ", char)
			char = GetRandomChaoWheelCharacter(1)[0]
		}
		return char
	}

	newRarities := rarities
	items := []string{}
	for i, rarity := range rarities {
		if rarity == 100 { // Character
			if !allowedRarity(100) { // if character not possible
				rarity = 2 // degrade to rarity 2 Chao
			} else {
				char := getUnusedCharacter()
				items = append(items, char)
			}
		}
		/*
			if rarity == 3 { // Rarity 3 Chao
				if !allowedRarity(3) { // if impossible
					rarity = 2 // degrade
				} else {
					chao, err := getUnusedChao(3)
					if err != nil {
						return []string{}, []int64{}, err
					}
					items = append(items, chao)
				}
			}
		*/
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
				// TODO: what do we do now?
				items = append(items, enums.ChaoIDStrHeroChao)
			} else {
				chao, err := getUnusedChao(1)
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
