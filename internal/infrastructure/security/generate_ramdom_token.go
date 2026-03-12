package security

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateRandomToken(length int) (string, error) {
	b := make([]byte, length)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(b)[:length]
	fmt.Println("Token :", token)

	return token, nil
}
