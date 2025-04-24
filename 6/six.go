package six

import (
	"crypto/rand"
	"errors"
)

const (
	useChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func generatePassword(n int) (string, error) {
	if n <= 0 {
		return "", errors.New("n must be greater than zero")
	}

	password := make([]byte, n)
	if _, err := rand.Read(password); err != nil {
		return "", err
	}

	for i := range password {
		password[i] = useChars[int(password[i])%len(useChars)]
	}

	return string(password), nil
}
