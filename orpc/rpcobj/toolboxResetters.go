package rpcobj

import (
    "github.com/fluofoxxo/outrun/db"
    "github.com/fluofoxxo/outrun/netobj"
    "github.com/fluofoxxo/outrun/obj"
)

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

func (t *Toolbox) ResetPlayerVarious(uid string, reply *ToolboxReply) error {
    player, err := db.GetPlayer(uid)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    player.PlayerVarious = netobj.DefaultPlayerVarious()
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

func (t *Toolbox) ResetMapInfo(uid string, reply *ToolboxReply) error {
    player, err := db.GetPlayer(uid)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    player.MileageMapState.MapInfo = obj.DefaultMapInfo()
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

func (t *Toolbox) ResetRouletteInfo(uid string, reply *ToolboxReply) error {
    player, err := db.GetPlayer(uid)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    player.RouletteInfo = netobj.DefaultRouletteInfo()
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

func (t *Toolbox) ResetLastWheelOptions(uid string, reply *ToolboxReply) error {
    player, err := db.GetPlayer(uid)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    player.LastWheelOptions = netobj.DefaultWheelOptions(player.PlayerState.NumRouletteTicket, player.RouletteInfo.RouletteCountInPeriod)
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
