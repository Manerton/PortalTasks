package main

import (
	"fmt"
	"main/RsaAesTask/myaes"
	"main/RsaAesTask/myrsa"
)

func testRSA() {
	message := []byte("Hello world! RSA")
	hmacKey := []byte("hmacsecretkey")
	fmt.Println("Оригинальный msg:", string(message))

	// Ключ можно закодировать в AES
	privateKey, err := myrsa.GenerateNewKey()
	if err != nil {
		return
	}

	label := "testLabel"
	encryptedMSG, err := myrsa.EncryptRSA(&privateKey.PublicKey, message, label, hmacKey)
	if err != nil {
		return
	}
	fmt.Println("Зашифрованное msg:")
	fmt.Println(string(encryptedMSG))

	decryptedMSG, err := myrsa.DecryptRSA(privateKey, encryptedMSG, label, hmacKey)
	if err != nil {
		return
	}
	fmt.Println("Расшифрованное msg:", string(decryptedMSG))
}

func testAES() {
	// Ключ можно закодировать в RSA
	key := []byte("0123456789abcdef0123456789abcdef")
	message := []byte("Hello world! AES")
	fmt.Println("Оригинальный msg:", string(message))

	encryptedMSG, err := myaes.EncryptAES(key, message)
	if err != nil {
		return
	}
	fmt.Println("Зашифрованное msg:")
	fmt.Println(string(encryptedMSG))

	decryptedMSG, err := myaes.DecryptAES(key, encryptedMSG)
	if err != nil {
		return
	}
	fmt.Println("Расшифрованное msg:", string(decryptedMSG))
}

func main() {
	fmt.Println("-----TEST RSA-----")
	testRSA()
	fmt.Println("-----TEST AES-----")
	testAES()
}
