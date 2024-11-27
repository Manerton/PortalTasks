package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Устаревший и не безопасный и атакам по времени
// Уязвим к коллизиям, переборам
func hashMD5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

func hashSHA256(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	message := "HelloWorld123"
	controlMessage := "HelloWorld123"

	fmt.Println("MSG:", message)
	hashSha256 := hashSHA256(message)
	fmt.Println("SHA256:", hashSha256)

	controlHashSha256 := hashSHA256(controlMessage)
	if hashSha256 != controlHashSha256 {
		fmt.Println("Error")
	} else {
		fmt.Println("Ok SHA256")
	}

	hashMd5 := hashMD5(message)
	fmt.Println("MD5:", hashMd5)

	controlHashMD5 := hashMD5(controlMessage)
	if hashMd5 != controlHashMD5 {
		fmt.Println("Error")
	} else {
		fmt.Println("Ok MD5")
	}

}
