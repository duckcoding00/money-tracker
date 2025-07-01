package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	now = time.Now().Format("20060102")
)

func genRandomAlphaNumeric(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
func GenID(types string, id int) string {
	randomCode := genRandomAlphaNumeric(4)
	return fmt.Sprintf("%s_%d_%s_%s", types, id, now, randomCode)
}
