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

func generateBlock(key string) (block cipher.Block) {
	pk := padding(key)
	block, err := aes.NewCipher(pk)
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}

	return block
}

func encrypt(block cipher.Block, plainText string) (encypted []byte) {
	data := padding(plainText)
	encrypted := make([]byte, aes.BlockSize+len(data))

	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Printf("err: %s\n", err)
	}

	encryptMode := cipher.NewCBCEncrypter(block, iv)
	encryptMode.CryptBlocks(encrypted[aes.BlockSize:], data)

	return encrypted
}

func decrypte(block cipher.Block, encrypted []byte) (decrypted []byte) {
	decrypted = make([]byte, len(encrypted[aes.BlockSize:]))
	decryptMode := cipher.NewCBCDecrypter(block, encrypted[:aes.BlockSize])
	decryptMode.CryptBlocks(decrypted, encrypted[aes.BlockSize:])

	return decrypted
}

func main() {
	plainText := "example no plain text death."
	key := "example no key death."

	block := generateBlock(key)

	encrypted := encrypt(block, plainText)
	fmt.Printf("encrypted: %v\n", encrypted)

	decrypted := decrypte(block, encrypted)
	fmt.Printf("decrypted: %s\n", string(decrypted))
}
