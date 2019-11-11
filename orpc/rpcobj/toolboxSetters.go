package rpcobj

import (
    "encoding/json"

    "github.com/fluofoxxo/outrun/config/eventconf"
    "github.com/fluofoxxo/outrun/db"
)

func (t *Toolbox) SetRings(args ChangeValueArgs, reply *ToolboxReply) error {
    player, err := db.GetPlayer(args.UID)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    newRings := args.Value.(int64)
    player.PlayerState.NumRings = newRings
    err = db.SavePlayer(player)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to save player: " + err.Error()
    }
    reply.Status = StatusOK
    reply.Info = "OK"
    return nil
}

func (t *Toolbox) SetRedRings(args ChangeValueArgs, reply *ToolboxReply) error {
    player, err := db.GetPlayer(args.UID)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    newRedRings := args.Value.(int64)
    player.PlayerState.NumRedRings = newRedRings
    err = db.SavePlayer(player)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to save player: " + err.Error()
    }
    reply.Status = StatusOK
    reply.Info = "OK"
    return nil
}

func (t *Toolbox) SetBuyRings(args ChangeValueArgs, reply *ToolboxReply) error {
    player, err := db.GetPlayer(args.UID)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    newRings := args.Value.(int64)
    player.PlayerState.NumBuyRings = newRings
    err = db.SavePlayer(player)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to save player: " + err.Error()
    }
    reply.Status = StatusOK
    reply.Info = "OK"
    return nil
}

func (t *Toolbox) SetBuyRedRings(args ChangeValueArgs, reply *ToolboxReply) error {
    player, err := db.GetPlayer(args.UID)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    newRedRings := args.Value.(int64)
    player.PlayerState.NumBuyRedRings = newRedRings
    err = db.SavePlayer(player)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to save player: " + err.Error()
    }
    reply.Status = StatusOK
    reply.Info = "OK"
    return nil
}

func (t *Toolbox) SetMainCharacter(args ChangeValueArgs, reply *ToolboxReply) error {
    player, err := db.GetPlayer(args.UID)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    newMainCharaID := args.Value.(string)
    player.PlayerState.MainCharaID = newMainCharaID
    err = db.SavePlayer(player)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to save player: " + err.Error()
    }
    reply.Status = StatusOK
    reply.Info = "OK"
    return nil
}

func (t *Toolbox) SetSubCharacter(args ChangeValueArgs, reply *ToolboxReply) error {
    player, err := db.GetPlayer(args.UID)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    newSubCharaID := args.Value.(string)
    player.PlayerState.SubCharaID = newSubCharaID
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

func (t *Toolbox) SetUsername(args ChangeValueArgs, reply *ToolboxReply) error {
    player, err := db.GetPlayer(args.UID)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    newUsername := args.Value.(string)
    player.Username = newUsername
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

func (t *Toolbox) SetCampaignEpisode(args ChangeValueArgs, reply *ToolboxReply) error {
    player, err := db.GetPlayer(args.UID)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    player.MileageMapState.Episode = args.Value.(int64)
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

func (t *Toolbox) SetCampaignChapter(args ChangeValueArgs, reply *ToolboxReply) error {
    player, err := db.GetPlayer(args.UID)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    player.MileageMapState.Chapter = args.Value.(int64)
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

func (t *Toolbox) SetCampaignPoint(args ChangeValueArgs, reply *ToolboxReply) error {
    player, err := db.GetPlayer(args.UID)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    player.MileageMapState.Point = args.Value.(int64)
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

func (t *Toolbox) SetRank(args ChangeValueArgs, reply *ToolboxReply) error {
    player, err := db.GetPlayer(args.UID)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    player.PlayerState.Rank = args.Value.(int64)
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

func (t *Toolbox) SetRouletteTickets(args ChangeValueArgs, reply *ToolboxReply) error {
    player, err := db.GetPlayer(args.UID)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    player.PlayerState.NumRouletteTicket = args.Value.(int64)
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

func (t *Toolbox) SetRouletteInfoResetTime(args ChangeValueArgs, reply *ToolboxReply) error {
    player, err := db.GetPlayer(args.UID)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    player.RouletteInfo.RoulettePeriodEnd = args.Value.(int64)
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

func (t *Toolbox) SetPersonalEvents(args ChangeValueArgs, reply *ToolboxReply) error {
    player, err := db.GetPlayer(args.UID)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    newEventList := args.Value.([]eventconf.ConfiguredEvent)
    player.PersonalEvents = newEventList
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

func (t *Toolbox) SetPersonalEventsJSON(args ChangeValueArgs, reply *ToolboxReply) error {
    player, err := db.GetPlayer(args.UID)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }

    var events []eventconf.ConfiguredEvent
    err = json.Unmarshal(args.Value.([]byte), events)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to unmarshal value: " + err.Error()
        return err
    }
    newEventList := events
    player.PersonalEvents = newEventList
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
