package db

import (
    "errors"
    bolt "go.etcd.io/bbolt"
    "log"
    "time"

    "github.com/fluofoxxo/outrun/helper"
)

var db *bolt.DB

func Set(bucket, key string, value []byte) error {
    CheckIfDBSet()
    value = helper.Compress(value)
    err := db.Update(func(tx *bolt.Tx) error {
        bucket, err := tx.CreateBucketIfNotExists([]byte(bucket))
        if err != nil {
            log.Println("[ERR] (Set) Failure creating bucket: " + err.Error())
            return err
        }
        err = bucket.Put([]byte(key), value)
        if err != nil {
            log.Println("[ERR] (Set) Failure putting in bucket: " + err.Error())
            return err
        }
        return nil
    })
    // error handling is already done in-func
    if err != nil {
        return err // maybe return nil?
    }
    return nil
}
func Get(bucket, key string) ([]byte, bool, error) {
    CheckIfDBSet()
    resultChan := make(chan []byte)
    errChan := make(chan error)
    go db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(bucket))
        v := b.Get([]byte(key))
        if v == nil {
            log.Println("[ERR] (Get) No value named '" + key + "' in bucket '" + bucket + "'.")
            err := errors.New("no value named " + key + " in bucket " + bucket)
            go func() {
                resultChan <- []byte{}
                errChan <- err
            }()
            return err
        }
        resultChan <- v
        errChan <- nil
        return nil
    })
    var result []byte
    var err error
    outRes := "'" + bucket + "/" + key + "'"
    select { // get result
    case result = <-resultChan:
        log.Println("[OUT] (Get) Fetched " + outRes + " result")
    case <-time.After(3 * time.Second):
        log.Println("[OUT] (Get) Fetch " + outRes + " timeout!")
        return []byte{}, false, nil
    }
    select { // get error
    case err = <-errChan:
        log.Println("[OUT] (Get) Got error from " + outRes)
    case <-time.After(3 * time.Second):
        log.Println("[OUT] (Get) Error " + outRes + " timeout!")
        return []byte{}, false, nil
    }
    if err != nil {
        return []byte{}, false, err
    }

    result, err = helper.Decompress(result)
    if err != nil {
        log.Println("[ERR] (Get) Error in decompression: " + err.Error())
        return []byte{}, false, err
    }

    return result, true, nil
}
