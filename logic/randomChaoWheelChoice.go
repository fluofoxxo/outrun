package logic

import (
	"math/rand"

	"github.com/fluofoxxo/outrun/consts"
)

func GetRandomChaoWheelCharacter(count int) []string {
	// Make weighted random pull of a character using an... Unconventional method!
	polls := [1000]string{} // TODO: parameterize
	pollIndex := 0
	for pollIndex < 1000 {
		for cid, load := range consts.RandomChaoWheelCharacterPrizes {
			rint := rand.Intn(1000000 + 1)            // + 1 to make inclusive; high number for precision
			chosen := (float64(rint) / 10000) <= load // + 1 to make inclusive
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
	//pollChoice := polls[rand.Intn(len(polls))]
	for count > 0 {
		count--
		pollChoices = append(pollChoices, polls[rand.Intn(len(polls))])
	}
	return pollChoices
}
