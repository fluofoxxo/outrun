package dbaccess

import (
	"errors"
	"time"

	bolt "go.etcd.io/bbolt"

	"github.com/fluofoxxo/outrun/consts"
)

var db *bolt.DB
var DatabaseIsBusy = false

func Set(bucket, key string, value []byte) error {
	CheckIfDBSet()
	value = Compress(value) // compress the input first
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(key), value)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func Get(bucket, key string) ([]byte, error) {
	CheckIfDBSet()
	var value []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		value = b.Get([]byte(key))
		if value == nil {
			return errors.New("no value named '" + key + "' in bucket '" + bucket + "'")
		}
		return nil
	})
	result, derr := Decompress(value) // decompress the result
	if derr != nil {
		return result, derr
	}
	return result, err
}

func Delete(bucket, key string) error {
	CheckIfDBSet()
	return db.View(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Delete([]byte(key))
	})
}

func ForEachLogic(each func(tx *bolt.Tx) error) error {
	CheckIfDBSet()
	err := db.View(each)
	return err
}

func CheckIfDBSet() {
	if db == nil {
		bdb, err := bolt.Open(consts.DBFileName, 0600, &bolt.Options{Timeout: 3 * time.Second})
		if err != nil {
			panic(err)
		}
		db = bdb
	}
}
