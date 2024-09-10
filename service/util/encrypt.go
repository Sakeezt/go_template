package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"strings"
)

//go:generate mockery --name=Encrypts
type Encrypts interface {
	EncryptAES(stringToEncrypt string, passPhrase string) (resultEncode string, err error)
	DecryptAES(stringToDecrypt string, passPhrase string) (resultDecode string, err error)
}
type Encrypt struct{}

func NewEncrypts() (e *Encrypt) {
	return &Encrypt{}
}

func createHash(key string) string {
	harsher := sha256.New()
	harsher.Write([]byte(key))
	return string(harsher.Sum(nil))
}

func (e *Encrypt) EncryptAES(stringToEncrypt string, passPhrase string) (resultEncode string, err error) {
	block, err := aes.NewCipher([]byte(createHash(passPhrase)))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(strings.NewReader(passPhrase), nonce); err != nil {
		return "", err
	}

	byteData := []byte(stringToEncrypt)
	cipherText := gcm.Seal(nonce, nonce, byteData, nil)
	return hex.EncodeToString(cipherText), nil
}

func (e *Encrypt) DecryptAES(stringToDecrypt string, passPhrase string) (resultDecode string, err error) {
	byteData, err := hex.DecodeString(stringToDecrypt)

	block, err := aes.NewCipher([]byte(createHash(passPhrase)))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce := byteData[:nonceSize]
	cipherText := byteData[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
