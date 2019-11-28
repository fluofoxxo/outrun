package logic

import (
    "log"

    "github.com/fluofoxxo/outrun/consts"
    "github.com/fluofoxxo/outrun/db"
    "github.com/fluofoxxo/outrun/db/dbaccess"
    "github.com/fluofoxxo/outrun/netobj"
)

func FindPlayersByPassword(password string, silent bool) ([]netobj.Player, error) {
    playerIDs := []string{}
    players := []netobj.Player{}
    dbaccess.ForEachKey(consts.DBBucketPlayers, func(k, v []byte) error {
        playerIDs = append(playerIDs, string(k))
        return nil
    })
    for _, pid := range playerIDs {
        player, err := db.GetPlayer(pid)
        if err != nil {
            if !silent {
                return []netobj.Player{}, err
            } else {
                log.Printf("[WARN] (logic.FindPlayersByPassword) Unable to get player '%s': %s", pid, err.Error())
            }
        }
        if player.Password == password {
            players = append(players, player)
        }
    }
    return players, nil
}
