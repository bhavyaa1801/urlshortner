package main

import (
	"crypto/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortCode() string {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	for i := range b {
		b[i] = charset[int(b[i])%len(charset)]
	}

	return string(b)

}


func getUniqueShortCode() string {
    for {
        code := generateShortCode()

        if _, exists := urlStore[code]; !exists {
            return code
        }
    }
}