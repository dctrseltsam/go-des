package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

// https://stackoverflow.com/questions/41579325/golang-how-do-i-decrypt-with-des-cbc-and-pkcs7

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func DesEncryption(key, iv, plainText []byte) ([]byte, error) {
	block, err := des.NewCipher(key)

	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData := PKCS5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return cryted, nil

}

func DesDecryption(key, iv, cipherText []byte) ([]byte, error) {

	block, err := des.NewCipher(key)

	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func main() {
	originalText := "laberlaberlaber"
	fmt.Println(originalText)
	mytext := []byte(originalText)

	key := []byte{0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC}
	iv := []byte{0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC, 0xBC}

	cryptoText, _ := DesEncryption(key, iv, mytext)
	fmt.Printf("% x\n", string(cryptoText))
	decryptedText, _ := DesDecryption(key, iv, cryptoText)
	fmt.Println(string(decryptedText))
}
