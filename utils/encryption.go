package utils

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"errors"
)

func Encrypt3Des(key string, payload []byte) (string, error) {
	block, err := des.NewTripleDESCipher([]byte(key))
	if err != nil {
		return "", err
	}

	bs := block.BlockSize() // block size is 8 by default
	payloadBytes := Pkcs5Padding(payload, bs)

	if len(payloadBytes)%bs != 0 {
		return "", errors.New("need a multiple of the blocksize")
	}
	encrypted := make([]byte, len(payloadBytes))
	dst := encrypted

	for len(payloadBytes) > 0 {
		block.Encrypt(dst, payloadBytes[:bs])
		payloadBytes = payloadBytes[bs:]
		dst = dst[bs:]
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func Pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
