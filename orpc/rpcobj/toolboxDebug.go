package rpcobj

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/fluofoxxo/outrun/config/gameconf"
	"github.com/fluofoxxo/outrun/consts"
	"github.com/fluofoxxo/outrun/db"
	"github.com/fluofoxxo/outrun/db/dbaccess"
	"github.com/fluofoxxo/outrun/logic"
	"github.com/fluofoxxo/outrun/netobj"
	"github.com/fluofoxxo/outrun/netobj/constnetobjs"
	"github.com/fluofoxxo/outrun/obj"
	"github.com/fluofoxxo/outrun/obj/constobjs"
)

func (t *Toolbox) Debug_GetCampaignStatus(uid string, reply *ToolboxReply) error {
	player, err := db.GetPlayer(uid)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to get player: " + err.Error()
		return err
	}
	reply.Status = StatusOK
	reply.Info = strconv.Itoa(int(player.MileageMapState.Chapter)) + "," + strconv.Itoa(int(player.MileageMapState.Episode)) + "," + strconv.Itoa(int(player.MileageMapState.Point))
	return nil
}

func (t *Toolbox) Debug_GetAllPlayerIDs(nothing bool, reply *ToolboxReply) error {
	playerIDs := []string{}
	dbaccess.ForEachKey(consts.DBBucketPlayers, func(k, v []byte) error {
		playerIDs = append(playerIDs, string(k))
		return nil
	})
	final := strings.Join(playerIDs, ",")
	reply.Status = StatusOK
	reply.Info = final
	return nil
}

func (t *Toolbox) Debug_ResetPlayer(uid string, reply *ToolboxReply) error {
	_ = db.NewAccountWithID(uid)
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_GetRouletteInfo(uid string, reply *ToolboxReply) error {
	player, err := db.GetPlayer(uid)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to get player: " + err.Error()
		return err
	}
	rouletteInfo := player.RouletteInfo
	jri, err := json.Marshal(rouletteInfo)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to marshal RouletteInfo: " + err.Error()
		return err
	}
	reply.Status = StatusOK
	reply.Info = string(jri)
	return nil
}

func (t *Toolbox) Debug_ResetChaoRouletteGroup(uid string, reply *ToolboxReply) error {
	player, err := db.GetPlayer(uid)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to get player: " + err.Error()
		return err
	}
	chaoRouletteGroup := netobj.DefaultChaoRouletteGroup(player.PlayerState, player.GetAllNonMaxedCharacters(), player.GetAllNonMaxedChao())
	player.ChaoRouletteGroup = chaoRouletteGroup
	err = db.SavePlayer(player)
	if err != nil {
		reply.Status = StatusOK
		reply.Info = "OK"
		return err
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_ResetCharactersAndCompensate(uid string, reply *ToolboxReply) error {
	player, err := db.GetPlayer(uid)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to get player: " + err.Error()
		return err
	}
	toAdd := int64(0)
	for _, char := range player.CharacterState {
		toAdd += char.Level * 15
	}
	player.PlayerState.NumRedRings += toAdd
	player.CharacterState = netobj.DefaultCharacterState()
	err = db.SavePlayer(player)
	if err != nil {
		reply.Status = StatusOK
		reply.Info = "OK"
		return err
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_ResetChao(uid string, reply *ToolboxReply) error {
	player, err := db.GetPlayer(uid)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to get player: " + err.Error()
		return err
	}
	player.ChaoState = constnetobjs.NetChaoList
	err = db.SavePlayer(player)
	if err != nil {
		reply.Status = StatusOK
		reply.Info = "OK"
		return err
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_MigrateUser(uidToUID string, reply *ToolboxReply) error {
	uidSrc := strings.Split(uidToUID, "->")
	if len(uidSrc) != 2 {
		reply.Status = StatusOtherError
		reply.Info = "improperly formatted string (Example: 1234567890->1987654321)"
	}

	fromUID := uidSrc[0]
	toUID := uidSrc[1]
	oldPlayer, err := db.GetPlayer(fromUID)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = err.Error()
		return err
	}
	currentPlayer, err := db.GetPlayer(toUID)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = err.Error()
		return err
	}
	currentPlayer.PlayerState = oldPlayer.PlayerState
	currentPlayer.Username = oldPlayer.Username
	currentPlayer.LastLogin = oldPlayer.LastLogin
	currentPlayer.CharacterState = oldPlayer.CharacterState
	currentPlayer.ChaoState = oldPlayer.ChaoState
	currentPlayer.MileageMapState = oldPlayer.MileageMapState
	currentPlayer.MileageFriends = oldPlayer.MileageFriends
	currentPlayer.PlayerVarious = oldPlayer.PlayerVarious
	currentPlayer.LastWheelOptions = oldPlayer.LastWheelOptions
	currentPlayer.ChaoRouletteGroup = oldPlayer.ChaoRouletteGroup
	currentPlayer.RouletteInfo = oldPlayer.RouletteInfo

	err = db.SavePlayer(currentPlayer)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = err.Error()
		return err
	}

	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_UsernameSearch(username string, reply *ToolboxReply) error {
	playerIDs := []string{}
	dbaccess.ForEachKey(consts.DBBucketPlayers, func(k, v []byte) error {
		playerIDs = append(playerIDs, string(k))
		return nil
	})
	sameUsernames := []string{}
	for _, uid := range playerIDs {
		player, err := db.GetPlayer(uid)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = "error getting ID " + uid + ": " + err.Error()
			return err
		}
		if player.Username == username {
			sameUsernames = append(sameUsernames, player.ID)
		}
	}
	if len(sameUsernames) == 0 {
		reply.Status = StatusOtherError
		reply.Info = "unable to find ID for username " + username
		return nil
	}
	reply.Status = StatusOK
	reply.Info = strings.Join(sameUsernames, ",")
	return nil
}

func (t *Toolbox) Debug_RawPlayer(uid string, reply *ToolboxReply) error {
	playerSrc, err := dbaccess.Get(consts.DBBucketPlayers, uid)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = err.Error()
		return err
	}
	reply.Status = StatusOK
	reply.Info = string(playerSrc)
	return nil
}

func (t *Toolbox) Debug_ResetCharacterState(uid string, reply *ToolboxReply) error {
	player, err := db.GetPlayer(uid)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to get player: " + err.Error()
		return err
	}
	player.CharacterState = netobj.DefaultCharacterState()
	err = db.SavePlayer(player)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = err.Error()
		return err
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_MatchPlayersToGameConf(uids string, reply *ToolboxReply) error {
	allUIDs := strings.Split(uids, ",")
	for _, uid := range allUIDs {
		player, err := db.GetPlayer(uid)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = fmt.Sprintf("unable to get player %s: ", uid) + err.Error()
			return err
		}
		player.CharacterState = netobj.DefaultCharacterState() // already uses AllCharactersUnlocked
		player.ChaoState = constnetobjs.DefaultChaoState()     // already uses AllChaoUnlocked
		player.PlayerState.MainCharaID = gameconf.CFile.DefaultMainCharacter
		player.PlayerState.SubChaoID = gameconf.CFile.DefaultSubChao
		player.PlayerState.MainChaoID = gameconf.CFile.DefaultMainChao
		player.PlayerState.SubCharaID = gameconf.CFile.DefaultSubCharacter
		player.PlayerState.NumRings = gameconf.CFile.StartingRings
		player.PlayerState.NumRedRings = gameconf.CFile.StartingRedRings
		player.PlayerState.Energy = gameconf.CFile.StartingEnergy
		err = db.SavePlayer(player)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = fmt.Sprintf("error saving player %s: ", uid) + err.Error()
			return err
		}
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_PrepTag1p0(uids string, reply *ToolboxReply) error {
	allUIDs := strings.Split(uids, ",")
	sqrt := func(n int64) int64 {
		fn := float64(n)
		result := math.Sqrt(fn)
		return int64(result)
	}

	for _, uid := range allUIDs {
		player, err := db.GetPlayer(uid)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = fmt.Sprintf("unable to get player %s: ", uid) + err.Error()
			return err
		}
		// conditions for exemption
		if player.MileageMapState.Episode >= 25 { // player is exempt from reset
			continue
		}
		player.CharacterState = netobj.DefaultCharacterState() // already uses AllCharactersUnlocked
		player.ChaoState = constnetobjs.DefaultChaoState()     // already uses AllChaoUnlocked
		player.PlayerState.MainCharaID = gameconf.CFile.DefaultMainCharacter
		player.PlayerState.SubChaoID = gameconf.CFile.DefaultSubChao
		player.PlayerState.MainChaoID = gameconf.CFile.DefaultMainChao
		player.PlayerState.SubCharaID = gameconf.CFile.DefaultSubCharacter
		player.PlayerState.NumRings = sqrt(player.PlayerState.NumRings) * 3
		player.PlayerState.NumRedRings = sqrt(player.PlayerState.NumRedRings)
		player.PlayerState.Energy = gameconf.CFile.StartingEnergy
		player.PlayerState.Items = constobjs.DefaultPlayerStateItems
		player.PlayerState.Rank = 0 // for some reason, this gets incremented 1 by the game

		player.MileageMapState = netobj.DefaultMileageMapState() // reset campaign

		err = db.SavePlayer(player)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = fmt.Sprintf("error saving player %s: ", uid) + err.Error()
			return err
		}
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_PlayersByPassword(password string, reply *ToolboxReply) error {
	foundPlayers, err := logic.FindPlayersByPassword(password, false)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "error finding players by password: " + err.Error()
		return err
	}
	playerIDs := []string{}
	for _, player := range foundPlayers {
		playerIDs = append(playerIDs, player.ID)
	}
	final := strings.Join(playerIDs, ",")
	reply.Status = StatusOK
	reply.Info = final
	return nil
}

func (t *Toolbox) Debug_ResetPlayersRank(uids string, reply *ToolboxReply) error {
	allUIDs := strings.Split(uids, ",")
	for _, uid := range allUIDs {
		player, err := db.GetPlayer(uid)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = fmt.Sprintf("unable to get player %s: ", uid) + err.Error()
			return err
		}
		player.PlayerState.Rank = 0 // for some reason, this gets incremented 1 by the game
		err = db.SavePlayer(player)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = fmt.Sprintf("error saving player %s: ", uid) + err.Error()
			return err
		}
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_FixWerehogRedRings(uids string, reply *ToolboxReply) error {
	wh := constobjs.CharacterWerehog
	whid := wh.ID
	whrr := wh.PriceRedRings
	allUIDs := strings.Split(uids, ",")
	for _, uid := range allUIDs {
		player, err := db.GetPlayer(uid)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = fmt.Sprintf("unable to get player %s: ", uid) + err.Error()
			return err
		}
		i := player.IndexOfChara(whid)
		if i == -1 {
			reply.Status = StatusOK
			reply.Info = "index not found!"
			return fmt.Errorf("index not found!")
		}
		player.CharacterState[i].Character.PriceRedRings = whrr // TODO: check if needed
		player.CharacterState[i].PriceRedRings = whrr
		err = db.SavePlayer(player)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = fmt.Sprintf("error saving player %s: ", uid) + err.Error()
			return err
		}
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) Debug_FixCharacterPrices(uids string, reply *ToolboxReply) error {
	// TODO: This function possibly needs to be adjusted if event characters are ever added.
	cmap := map[string]obj.Character{
		"300000": constobjs.CharacterSonic,
		"300001": constobjs.CharacterTails,
		"300002": constobjs.CharacterKnuckles,
		"300003": constobjs.CharacterAmy,
		"300004": constobjs.CharacterShadow,
		"300005": constobjs.CharacterBlaze,
		"300006": constobjs.CharacterRouge,
		"300007": constobjs.CharacterOmega,
		"300008": constobjs.CharacterBig,
		"300009": constobjs.CharacterCream,
		"300010": constobjs.CharacterEspio,
		"300011": constobjs.CharacterCharmy,
		"300012": constobjs.CharacterVector,
		"300013": constobjs.CharacterSilver,
		"300014": constobjs.CharacterMetalSonic,
		"300015": constobjs.CharacterClassicSonic,
		"300016": constobjs.CharacterWerehog,
		"300017": constobjs.CharacterSticks,
		"300018": constobjs.CharacterTikal,
		"300019": constobjs.CharacterMephiles,
		"300020": constobjs.CharacterPSISilver,
		"301000": constobjs.CharacterAmitieAmy,
		"301001": constobjs.CharacterGothicAmy,
		"301002": constobjs.CharacterHalloweenShadow,
		"301003": constobjs.CharacterHalloweenRouge,
		"301004": constobjs.CharacterHalloweenOmega,
		"301005": constobjs.CharacterXMasSonic,
		"301006": constobjs.CharacterXMasTails,
		"301007": constobjs.CharacterXMasKnuckles,
	}
	allUIDs := strings.Split(uids, ",")
	for _, uid := range allUIDs {
		log.Printf("[RPC-DEBUG] Resetting character prices for player %s\n", uid)
		player, err := db.GetPlayer(uid)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = fmt.Sprintf("unable to get player %s: ", uid) + err.Error()
			return err
		}
		for i, netchar := range player.CharacterState {
			cid := netchar.ID
			char, ok := cmap[cid]
			if !ok {
				reply.Status = StatusOtherError
				reply.Info = "character with ID '" + cid + "' was not found in CharacterState for player ID '" + uid + "'"
				return fmt.Errorf(reply.Info)
			}
			defaultCost := char.Cost
			defaultNumRedRings := char.NumRedRings
			defaultPrice := char.Price
			defaultPriceRedRings := char.PriceRedRings
			player.CharacterState[i].Character.Cost = defaultCost // TODO: check if needed
			player.CharacterState[i].Cost = defaultCost
			player.CharacterState[i].Character.NumRedRings = defaultNumRedRings // TODO: check if needed
			player.CharacterState[i].NumRedRings = defaultNumRedRings
			player.CharacterState[i].Character.Price = defaultPrice // TODO: check if needed
			player.CharacterState[i].Price = defaultPrice
			player.CharacterState[i].Character.PriceRedRings = defaultPriceRedRings // TODO: check if needed
			player.CharacterState[i].PriceRedRings = defaultPriceRedRings
		}
		err = db.SavePlayer(player)
		if err != nil {
			reply.Status = StatusOtherError
			reply.Info = fmt.Sprintf("error saving player %s: ", uid) + err.Error()
			return err
		}
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}
