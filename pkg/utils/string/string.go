package string_util

import (
	"time"

	"math/rand"
)

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

func GeneratePassword(length int) string {
	if length < 8 || length > 32 {
		length = 8
	}

	const (
		lowercase    = "abcdefghijklmnopqrstuvwxyz"
		uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		digits       = "0123456789"
		specialChars = "@$!%*?&"
	)

	allChars := lowercase + uppercase + digits + specialChars
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, length)

	// Ensure the password contains at least one of each required character type
	result[0] = lowercase[random.Intn(len(lowercase))]
	result[1] = uppercase[random.Intn(len(uppercase))]
	result[2] = digits[random.Intn(len(digits))]
	result[3] = specialChars[random.Intn(len(specialChars))]

	// Fill the rest of the password with random characters from allChars
	for i := 4; i < length; i++ {
		result[i] = allChars[random.Intn(len(allChars))]
	}

	// Shuffle the result to ensure the first four characters are not predictable
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return string(result)
}
