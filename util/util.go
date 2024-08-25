package util

import (
	"crypto/sha256"
	"github.com/itchyny/base58-go"
	"log"
)

func Sha256Of(str string) []byte {
	sha := sha256.New()
	sha.Write([]byte(str))
	return sha.Sum(nil)
}

func Base58Encoded(bytes []byte) (string, error) {
	encoding := base58.BitcoinEncoding
	encode, err := encoding.Encode(bytes)
	if err != nil {
		return "", err
	}
	return string(encode), nil
}

func PrintError(err error) {
	if err != nil {
		log.Println("error occurred: ", err)
	}
}
