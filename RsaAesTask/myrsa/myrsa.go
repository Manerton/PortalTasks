package myrsa

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func GenerateNewKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

func EncryptRSA(publicKey *rsa.PublicKey, data []byte, labelStr string, hmacKey []byte) ([]byte, error) {
	label := []byte(labelStr)
	hash := sha256.New()

	myHmac := hmac.New(sha256.New, hmacKey)
	myHmac.Write(data)
	hmacDigest := myHmac.Sum(nil)

	encryodData, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, data, label)
	if err != nil {
		return nil, err
	}

	encryodDataWithHMAC := append(encryodData, hmacDigest...)
	return encryodDataWithHMAC, nil
}

func DecryptRSA(privateKey *rsa.PrivateKey, data []byte, labelStr string, hmacKey []byte) ([]byte, error) {
	label := []byte(labelStr)
	hash := sha256.New()

	hmacLen := sha256.Size
	encryptedData := data[:len(data)-hmacLen]
	receivedHmac := data[len(data)-hmacLen:]

	decryptData, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, encryptedData, label)
	if err != nil {
		return nil, err
	}

	myHmac := hmac.New(sha256.New, hmacKey)
	myHmac.Write(decryptData)
	hmacDigest := myHmac.Sum(nil)

	if !hmac.Equal(receivedHmac, hmacDigest) {
		return nil, fmt.Errorf("HMAC mismatch: data has been tampered with")
	}

	return decryptData, nil
}
