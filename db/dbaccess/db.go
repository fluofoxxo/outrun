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
		LockDB()
		defer UnlockDB()
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
	resultChan := make(chan []byte)
	errChan := make(chan error)
	go db.View(func(tx *bolt.Tx) error {
		LockDB()
		defer UnlockDB()
		b := tx.Bucket([]byte(bucket))
		value := b.Get([]byte(key))
		if value == nil {
			go func() {
				resultChan <- []byte{}
				errChan <- errors.New("no value named '" + key + "' in bucket '" + bucket + "'")
			}()
			return nil // don't return anything since we never use this function's returned error
		}
		resultChan <- value
		errChan <- nil
		return nil
	})
	// TODO: handle timeouts, as unoften as they are
	resultRecv := <-resultChan
	err := <-errChan
	result, derr := Decompress(resultRecv) // decompress the result
	if derr != nil {
		if err != nil {
			return result, err
		}
		return result, derr
	}
	return result, err
}

func Delete(bucket, key string) error {
	CheckIfDBSet()
	LockDB()
	UnlockDB()
	return db.View(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Delete([]byte(key))
	})
}

func ForEachLogic(each func(tx *bolt.Tx) error) error {
	CheckIfDBSet()
	LockDB()
	defer UnlockDB()
	err := db.View(each)
	return err
}

func LockDB() {
	// TODO: this whole locking mechanism is a race condition, but goroutines on the database can be extremely dangerous.
	// find a better way to do this!
	for DatabaseIsBusy { // wait until we're not busy anymore
	}
	DatabaseIsBusy = true
}

func UnlockDB() {
	DatabaseIsBusy = false
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
