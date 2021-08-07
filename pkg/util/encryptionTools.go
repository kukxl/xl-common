package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

type encryptionTools struct {}

var EncryptionTools = new(encryptionTools)

func (*encryptionTools) Md5Encode(content string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(content)))
}

func (*encryptionTools) Sha256Encode(content string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(content)))
}

func (*encryptionTools) Base64Encode(content string) string {
	return base64.StdEncoding.EncodeToString([]byte(content))
}

func (*encryptionTools) Base64Decode(encryptedData string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", bytes), nil
}

