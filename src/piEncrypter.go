package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"

	"github.com/ethan6077/pi-007/src/models"
)

func EncryptPiData(form models.Form) models.Form {
	println("Encrypting Pi Data")

	piFields := [5]string{"phone", "email", "password", "passport", "creditCard"}

	for _, v := range piFields {
		switch v {
		case "phone":
			form.Phone = encryptWithKey(form.Phone)
		case "email":
			form.Email = encryptWithKey(form.Email)
		}
	}

	return form
}

func encryptWithKey(plainText string) string {

	block, newCipherErr := aes.NewCipher([]byte(PI_KEY))
	if newCipherErr != nil {
		panic(newCipherErr)
	}

	plainTextInBytes := []byte(plainText)

	ciphertext := make([]byte, aes.BlockSize+len(plainTextInBytes))
	iv := ciphertext[:aes.BlockSize]
	if _, ioReadErr := io.ReadFull(rand.Reader, iv); ioReadErr != nil {
		panic(ioReadErr)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainTextInBytes)

	return hex.EncodeToString(ciphertext)
}
