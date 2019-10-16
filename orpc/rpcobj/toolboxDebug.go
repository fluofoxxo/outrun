package rpcobj

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/fluofoxxo/outrun/consts"
	"github.com/fluofoxxo/outrun/db"
	"github.com/fluofoxxo/outrun/db/dbaccess"
	"github.com/fluofoxxo/outrun/netobj"
	"github.com/fluofoxxo/outrun/netobj/constnetobjs"
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

	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}
