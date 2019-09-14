package rpcobj

import (
	"strconv"

	"github.com/fluofoxxo/outrun/db"
	"github.com/fluofoxxo/outrun/netobj"
)

type Toolbox struct {
}

func (t *Toolbox) RegisterPlayerWithID(uid string, reply *ToolboxReply) error {
	player := db.NewAccountWithID(uid)
	err := db.SavePlayer(player)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to save player: " + err.Error()
		return err
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

func (t *Toolbox) ResetCampaign(uid string, reply *ToolboxReply) error {
	player, err := db.GetPlayer(uid)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to get player: " + err.Error()
		return err
	}
	player.MileageMapState = netobj.DefaultMileageMapState()
	err = db.SavePlayer(player)
	if err != nil {
		reply.Status = StatusOtherError
		reply.Info = "unable to save player: " + err.Error()
		return err
	}
	reply.Status = StatusOK
	reply.Info = "OK"
	return nil
}

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

type ToolboxReply struct {
	Status uint
	Info   string
}
