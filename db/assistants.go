package db

import (
	"strconv"
	"strings"
	"time"

	"github.com/fluofoxxo/outrun/consts"
	"github.com/fluofoxxo/outrun/db/dbaccess"
	"github.com/fluofoxxo/outrun/netobj"

	bolt "go.etcd.io/bbolt"
)

func NewAccount() netobj.Player {
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
}

func ParseSIDEntry(sidResult []byte) (string, int64) {
	split := strings.Split(string(sidResult), "/")
	uid := split[0]
	timeAssigned, _ := strconv.Atoi(split[1])
	return uid, int64(timeAssigned)
}

func IsValidSessionTime(sessionTime int64) bool {
	timeNow := time.Now().Unix()
	if sessionTime+consts.DBSessionExpiryTime < timeNow {
		return false
	}
	return true
}

func IsValidSessionID(sid []byte) (bool, error) {
	sidResult, err := dbaccess.Get(consts.DBBucketSessionIDs, string(sid))
	if err != nil {
		return false, err
	}
	_, sessionTime := ParseSIDEntry(sidResult)
	return IsValidSessionTime(sessionTime), err
}

func PurgeSessionID(sid string) error {
	err := dbaccess.Delete(consts.DBBucketSessionIDs, sid)
	return err
}

func PurgeAllExpiredSessionIDs() {
	keysToPurge := [][]byte{}
	each := func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(consts.DBBucketSessionIDs))
		err2 := bucket.ForEach(func(k, v []byte) error { // for each value in the session bucket
			_, sessionTime := ParseSIDEntry(v) // get time the session was created
			if !IsValidSessionTime(sessionTime) {
				keysToPurge = append(keysToPurge, k)
			}
			return nil
		})
		return err2
	}
	dbaccess.ForEachLogic(each) // do the logic above
	for _, key := range keysToPurge {
		PurgeSessionID(string(key))
	}
}
