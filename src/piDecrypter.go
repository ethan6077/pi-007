package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/ethan6077/pi-007/src/models"
)

func DecryptPiData(form models.Form, key string) models.Form {
	println("Decrypting Pi Data")

	piFields := [5]string{"phone", "email", "password", "passport", "creditCard"}

	for _, v := range piFields {
		switch v {
		case "phone":
			form.Phone = decriptWithKey(form.Phone, key)
		case "email":
			form.Email = decriptWithKey(form.Email, key)
		}
	}

	return form
}

func decriptWithKey(cryptoText string, key string) string {
	ciphertext, _ := hex.DecodeString(cryptoText)

	block, newCipherErr := aes.NewCipher([]byte(key))
	if newCipherErr != nil {
		panic(newCipherErr)
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
