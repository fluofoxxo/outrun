package constnetobjs

import (
	"math/rand"
	"strconv"

	"github.com/fluofoxxo/outrun/config/eventconf"
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
	// create ID
	uid := ""
	for i := range make([]byte, 10) {
		if i == 0 { // if first character
			uid += strconv.Itoa(rand.Intn(9) + 1)
		} else {
			uid += strconv.Itoa(rand.Intn(10))
		}
	}
	username := ""
	password := randChar("abcdefghijklmnopqrstuvwxyz1234567890", 10)
	key := randChar("abcdefghijklmnopqrstuvwxyz1234567890", 10)
	playerState := netobj.DefaultPlayerState()
	characterState := netobj.DefaultCharacterState()
	chaoState := GetAllNetChaoList() // needed for chaoRouletteAllowed
	mileageMapState := netobj.DefaultMileageMapState()
	mileageFriends := []netobj.MileageFriend{}
	playerVarious := netobj.DefaultPlayerVarious()
	rouletteInfo := netobj.DefaultRouletteInfo()
	wheelOptions := netobj.DefaultWheelOptions(playerState.NumRouletteTicket, rouletteInfo.RouletteCountInPeriod)
	// TODO: get rid of logic here?
	allowedCharacters := []string{}
	allowedChao := []string{}
	for _, chao := range chaoState {
		if chao.Level < 10 { // not max level
			allowedChao = append(allowedChao, chao.ID)
		}
	}
	for _, character := range characterState {
		if character.Star < 10 { // not max star
			allowedCharacters = append(allowedCharacters, character.ID)
		}
	}
	chaoRouletteGroup := netobj.DefaultChaoRouletteGroup(playerState, allowedCharacters, allowedChao)
	personalEvents := []eventconf.ConfiguredEvent{}
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
		personalEvents,
	)
}() // TODO: Solve duplication requirement with db/assistants.go
