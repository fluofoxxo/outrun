package cryption

import (
    "log"
    "net/http"
    "regexp"
)

var EncryptionIv = []byte("")
var EncryptionKey = []byte("Ec7bLaTdSuXuf5pW")

func CleanBytes(b []byte) []byte {
    re := regexp.MustCompile("[^\u0020-\u007f]+")
    return []byte(re.ReplaceAllLiteralString(string(b), ""))
}

func GetReceivedMessage(r *http.Request) []byte {
    err := r.ParseForm()
    if err != nil {
        log.Println("[ERR] Error in parsing form: " + err.Error())
    }
    param := r.Form.Get("param")
    iv := r.Form.Get("key")
    secure := r.Form.Get("secure")
    if secure != "1" {
        log.Println("[WARN] The secure flag from the client was not 1. Something's up.")
    }
    EncryptionIv = []byte(iv)
    paramUnB64 := B64Decode(param)
    decrypted := Decrypt(paramUnB64, EncryptionKey, EncryptionIv)
    decrypted = CleanBytes(decrypted)
    return decrypted
}
