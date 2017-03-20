package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"strconv"
)

func padding(data string) []byte {
	length := aes.BlockSize - (len(data) % aes.BlockSize) + len(data)
	format := "%-" + strconv.Itoa(length) + "s"
	return []byte(fmt.Sprintf(format, data))
}

func main() {
	plainText := padding("example no plain text death.")
	key := padding("example no key death.")

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}

	encrypted := make([]byte, aes.BlockSize+len(plainText))

	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Printf("err: %s\n", err)
	}

	encryptMode := cipher.NewCBCEncrypter(block, iv)
	encryptMode.CryptBlocks(encrypted[aes.BlockSize:], plainText)
	fmt.Printf("encrypted: %v\n", encrypted)

	decrypted := make([]byte, len(encrypted[aes.BlockSize:]))
	decryptMode := cipher.NewCBCDecrypter(block, encrypted[:aes.BlockSize])
	decryptMode.CryptBlocks(decrypted, encrypted[aes.BlockSize:])
	fmt.Printf("decrypted: %s\n", string(decrypted))
}
