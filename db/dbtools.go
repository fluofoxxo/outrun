package db

import (
    "encoding/json"
    "log"
    "math/rand"
    "strconv"
    "time"

    bolt "go.etcd.io/bbolt"

    "github.com/fluofoxxo/outrun/helper"
    "github.com/fluofoxxo/outrun/playerdata"
)

func UIDIsTaken(uid string) bool {
    _, worked, _ := Get("players", uid)
    return worked
}

func CheckIfDBSet() {
    if db == nil {
        bdb, err := bolt.Open("subsonic.db", 0600, &bolt.Options{Timeout: 3 * time.Second})
        if err != nil {
            panic(err)
        }
        db = bdb
    }
}

func PopulateNewAccount(uid, password, key string) error {
    player := pdata.NewPlayer(uid, "", password, key) // second arg is blank username for new user
    jplayerBytes, err := json.Marshal(player)
    if err != nil {
        log.Println("[ERR] (populateNewAccount) Error marshalling: " + err.Error())
        return err
    }
    jplayer := jplayerBytes
    err = Set("players", uid, jplayer)
    return err
}

func AssignSessionID(uid string) (string, error) {
    sid := helper.NewSessionID()
    err := Set("sessions/UIDtoSID", uid, []byte(sid))
    if err != nil {
        return sid, err
    }
    err = Set("sessions/SIDtoUID", sid, []byte(uid))
    if err != nil {
        return sid, err
    }
    return sid, nil
}

func SessionIDToUID(sid string) (string, error) {
    // TODO: IMPLEMENT MAX TIME FOR SESSION (1 HR)
    result, worked, err := Get("sessions/SIDtoUID", sid)
    if !worked {
        log.Println("[ERR] (uidBySessionID) Could not get UID by session " + sid)
        return "", err
    }
    return string(result), nil
}

func SessionIDToPlayer(sid string) (pdata.Player, error) {
    uid, err := SessionIDToUID(sid)
    if err != nil {
        return pdata.BLANK_PLAYER, err
    }
    player, err := GetPlayerByUID(uid)
    return player, err
}

func UIDToSessionID(uid string) (string, error) {
    // TODO: IMPLEMENT MAX TIME FOR SESSION (1 HR)
    result, worked, err := Get("sessions/UIDtoSID", uid)
    if !worked {
        log.Println("[ERR] (SessionIDByUID) Could not get session ID by UID " + uid)
        return "", err
    }
    return string(result), nil
}

func NewAccount() (string, string, string, string) {
    // generate a random uid
    uid := strconv.Itoa(rand.Intn(9999999999-1000000000) + 1000000000)
    // generate a random password (10 characters)
    password := helper.RandomCharacters("abcdefghijklmnopqrstuvwxyz1234567890", 10)
    // generate a random key (10 characters)
    key := helper.RandomCharacters("abcdefghijklmnopqrstuvwxyz1234567890", 10)
    // register the values
    PopulateNewAccount(uid, password, key)
    return uid, "", password, key // second arg is username
}

func GetPlayerByUID(uid string) (pdata.Player, error) {
    result, worked, err := Get("players", uid)
    if worked == false {
        return pdata.BLANK_PLAYER, err
    }
    var userData pdata.Player
    err = json.Unmarshal(result, &userData)
    if err != nil {
        return pdata.BLANK_PLAYER, err
    }
    return userData, nil
}

func SavePlayer(player pdata.Player) error {
    uid := player.UserID
    playerJson, err := json.Marshal(player)
    if err != nil {
        return err
    }
    err = Set("players", uid, playerJson)
    return err
}
