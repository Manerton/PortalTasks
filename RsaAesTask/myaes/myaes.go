package myaes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func EncryptAES(key []byte, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// Обеспечение защиты шифротекста от модификации за счёт GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	encryptedData := aesGCM.Seal(nonce, nonce, data, nil)
	return encryptedData, nil
}

func DecryptAES(key []byte, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := data[:aesGCM.NonceSize()]
	encryptedData := data[aesGCM.NonceSize():]
	decryptedData, err := aesGCM.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return nil, err
	}
	return decryptedData, nil

}
