package helper

import (
    "bytes"
    "compress/zlib"
    "encoding/json"
    "io/ioutil"
    "log"
    "math/rand"
    "net/http"
    "strconv"
    "strings"
    "time"

    "github.com/fluofoxxo/outrun/cryption"
)

const DB_DELIMITER = "<|>" // more complex = harder to trick, but more space in db
const DB_NODE_UIDS = "uids"

func NewSessionID() string {
    return "SUBSONIC_" + strconv.Itoa(int(time.Now().Unix()))
}

func Compress(s []byte) []byte {
    var b bytes.Buffer
    w := zlib.NewWriter(&b)
    w.Write(s)
    w.Close()
    return b.Bytes()
}

func Decompress(s []byte) ([]byte, error) {
    b := bytes.NewBuffer(s)
    r, err := zlib.NewReader(b)
    if err != nil {
        return []byte{}, err
    }
    result, err := ioutil.ReadAll(r)
    r.Close()
    if err != nil {
        return []byte{}, err
    }
    return result, nil
}

func JoinDelims(ss ...string) string {
    return Join(DB_DELIMITER, ss...)
}

func Join(delim string, ss ...string) string {
    return strings.Join(ss, delim)
}

func RandomCharacters(charset string, length int64) string {
    runes := []rune(charset)
    final := make([]rune, 10)
    for i := range final {
        final[i] = runes[rand.Intn(len(runes))]
    }
    return string(final)
}

func Respond(jstr []byte, w http.ResponseWriter) {
    responseJ := map[string]string{}
    exiv := "ITS IS A REAL IV"
    responseJ["secure"] = "1"
    responseJ["key"] = exiv
    enc := cryption.Encrypt(jstr, cryption.EncryptionKey, []byte(exiv))
    encb64 := cryption.B64Encode(enc)
    responseJ["param"] = encb64
    var result []byte
    result, err := json.Marshal(responseJ)
    if err != nil {
        log.Println("[ERR] (Respond) Error marshalling: " + err.Error())
        w.Write([]byte(`{"what": "sorry!"}`))
        panic(err) // prints traceback so we know
        return
    }
    w.Write(result)
}
