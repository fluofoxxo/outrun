package rpcobj

import (
    "strconv"

    "github.com/fluofoxxo/outrun/db"
)

func (t *Toolbox) GetUsername(uid string, reply *ToolboxReply) error {
    player, err := db.GetPlayer(uid)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    reply.Status = StatusOK
    reply.Info = player.Username
    return nil
}

func (t *Toolbox) GetRouletteTickets(uid string, reply *ToolboxReply) error {
    player, err := db.GetPlayer(uid)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    reply.Status = StatusOK
    reply.Info = strconv.Itoa(int(player.PlayerState.NumRouletteTicket))
    return nil
}

func (t *Toolbox) GetLastLogin(uid string, reply *ToolboxReply) error {
    player, err := db.GetPlayer(uid)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Info = "unable to get player: " + err.Error()
        return err
    }
    reply.Status = StatusOK
    reply.Info = strconv.Itoa(int(player.LastLogin))
    return nil
}

func (t *Toolbox) GetPersonalEvents(args ChangeValueArgs, reply *ToolboxValueReply) error {
    player, err := db.GetPlayer(args.UID)
    if err != nil {
        reply.Status = StatusOtherError
        reply.Result = "unable to get player: " + err.Error()
        return err
    }
    reply.Status = StatusOK
    reply.Result = player.PersonalEvents
    return nil
}
