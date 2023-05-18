package helper

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomFileName() string {
	// Generate 16 random bytes
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		panic(err) // Handle error appropriately in your application
	}

	// Convert random bytes to hexadecimal
	randomHex := hex.EncodeToString(randomBytes)

	return randomHex + ".jpg" // Customize the file extension if needed
}
