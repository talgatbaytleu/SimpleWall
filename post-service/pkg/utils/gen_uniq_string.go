package utils

import (
	"crypto/rand"
	"fmt"
	"time"
)

func GenUniqueString() string {
	timestamp := time.Now().UnixNano()
	randomBytes := make([]byte, 4)
	_, _ = rand.Read(randomBytes)
	return fmt.Sprintf("/%x-%x", timestamp, randomBytes)
}
