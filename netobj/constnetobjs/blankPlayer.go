package constnetobjs

import (
    "math/rand"
    "strconv"
    "time"

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
    uid := strconv.Itoa(rand.Intn(9999999999-1000000000) + 1000000000)
    username := ""
    password := randChar("abcdefghijklmnopqrstuvwxyz1234567890", 10)
    key := randChar("abcdefghijklmnopqrstuvwxyz1234567890", 10)
    lastLogin := time.Now().Unix()
    playerState := netobj.DefaultPlayerState()
    characterState := netobj.DefaultCharacterState()
    chaoState := []netobj.Chao{} // no chao for now
    mileageMapState := netobj.DefaultMileageMapState()
    return netobj.Player{
        uid,
        username,
        password,
        key,
        lastLogin,
        playerState,
        characterState,
        chaoState,
        mileageMapState,
    }
}() // copied from db.NewPlayer to avoid import loop
