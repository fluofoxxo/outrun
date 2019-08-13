package cryption

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
)

func B64Decode(s string) []byte {
    result, _ := base64.StdEncoding.DecodeString(s)
    return result
}

func B64Encode(b []byte) string {
    result := base64.StdEncoding.EncodeToString(b)
    return result
}

func Decrypt(content, key, iv []byte) []byte {
    //println(content, " ", key, " ", iv)
    //fmt.Println(string(content))
    //fmt.Println(content, " ", key, " ", iv)
    block, err := aes.NewCipher(key) // get cipher block from key
    if err != nil {
        panic(err)
    }
    cbcD := cipher.NewCBCDecrypter(block, iv) // create cbc decrypter
    cbcD.CryptBlocks(content, content)        // decrypt content
    return content
}

func Encrypt(content, key, iv []byte) []byte {
    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }
    content = PKCS5Padding(content, block.BlockSize())
    cbcE := cipher.NewCBCEncrypter(block, iv)
    cbcE.CryptBlocks(content, content)
    return content
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}
