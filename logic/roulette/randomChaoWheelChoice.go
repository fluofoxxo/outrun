package roulette

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/fluofoxxo/outrun/config"
	"github.com/fluofoxxo/outrun/consts"
)

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

func GetRandomChaoRouletteItems(rarities []int64, allowed []string) ([]string, error) { // TODO: Possibly rename to GetRandomChaoWheelItems?
	isAllowed := func(x string) bool {
		for _, s := range allowed {
			if s == x {
				return true
			}
		}
		return false
	}

	items := []string{}
	for _, rarity := range rarities {
		if rarity == 1 { // Rarity 1 Chao
			chao, err := GetRandomChaoWheelChao(1, 1)
			if err != nil {
				return []string{}, err
			}
			for !isAllowed(chao[0]) { // keep getting chao until we have one that is not max level
				if config.CFile.DebugPrints {
					log.Println("[DEBUG] Rarity 1 Chao Search")
				}
				chao, err = GetRandomChaoWheelChao(1, 1)
				if err != nil {
					return []string{}, err
				}
			}
			items = append(items, chao[0])
		} else if rarity == 2 { // Rarity 2 Chao
			chao, err := GetRandomChaoWheelChao(1, 2)
			if err != nil {
				return []string{}, err
			}
			for !isAllowed(chao[0]) { // keep getting chao until we have one that is not max level
				if config.CFile.DebugPrints {
					log.Println("[DEBUG] Rarity 2 Chao Search")
				}
				chao, err = GetRandomChaoWheelChao(1, 2)
				if err != nil {
					return []string{}, err
				}
			}
			items = append(items, chao[0])
		} else if rarity == 100 { // Character
			char := GetRandomChaoWheelCharacter(1)[0]
			for !isAllowed(char) { // keep getting character until we have one that is not max level
				if config.CFile.DebugPrints {
					log.Println("[DEBUG] Character Search")
				}
				char = GetRandomChaoWheelCharacter(1)[0]
			}
			items = append(items, char)
		} else { // Should never happen!
			return []string{}, fmt.Errorf("invalid rarity '%v'", rarity)
		}
	}
	return items, nil
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
