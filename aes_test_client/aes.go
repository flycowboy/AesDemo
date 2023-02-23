package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

func KeyEncrypt(keyStr string, cryptoText string) string {
	keyBytes := sha256.Sum256([]byte(keyStr))
	return encrypt(keyBytes[:], cryptoText)
}

func encrypt(key []byte, text string) string {
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext)
}

func KeyDecrypt(keyStr string, cryptoText string) string {
	keyBytes := sha256.Sum256([]byte(keyStr))
	return decrypt(keyBytes[:], cryptoText)
}

func decrypt(key []byte, cryptoText string) string {
	ciphertext, err := base64.StdEncoding.DecodeString(cryptoText)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext)
}
