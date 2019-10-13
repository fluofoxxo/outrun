package constnetobjs

import (
	"math/rand"
	"strconv"

	"github.com/fluofoxxo/outrun/netobj"
)

var BlankPlayer = func() netobj.Player {
	randChar := func(charset string, length int64) string {
		runes := []rune(charset)
		final := make([]rune, 10)
		for i := range final {
			final[i] = runes[rand.Intn(len(runes))]
		}
		return string(final)
	}
	uid := strconv.Itoa(rand.Intn(9999999999-1000000000) + 1000000000) // WARN: This large of an int blocks Outrun from being compiled for 32 bit archs. May be useful to solve.
	username := ""
	password := randChar("abcdefghijklmnopqrstuvwxyz1234567890", 10)
	key := randChar("abcdefghijklmnopqrstuvwxyz1234567890", 10)
	playerState := netobj.DefaultPlayerState()
	characterState := netobj.DefaultCharacterState()
	chaoState := []netobj.Chao{} // no chao for now
	mileageMapState := netobj.DefaultMileageMapState()
	mileageFriends := []netobj.MileageFriend{}
	playerVarious := netobj.DefaultPlayerVarious()
	rouletteInfo := netobj.DefaultRouletteInfo()
	wheelOptions := netobj.DefaultWheelOptions(playerState.NumRouletteTicket, rouletteInfo.RouletteCountInPeriod)
	chaoRouletteGroup := netobj.DefaultChaoRouletteGroup(playerState, []string{}) // TODO: fix having no exceptions. This can be problematic if characterState or chaoState is disturbed!
	return netobj.NewPlayer(
		uid,
		username,
		password,
		key,
		playerState,
		characterState,
		chaoState,
		mileageMapState,
		mileageFriends,
		playerVarious,
		wheelOptions,
		rouletteInfo,
		chaoRouletteGroup,
	)
}() // TODO: Solve duplication requirement with db/assistants.go
