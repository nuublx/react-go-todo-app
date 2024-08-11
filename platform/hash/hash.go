package hash

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/argon2"
)

// HashPassword hashes the password using Argon2id
func HashPassword(password string) (string, string, error) {
	if len(password) == 0 {
		return "", "", errors.New("password must not be empty")
	}
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", "", err
	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	return encodedSalt, encodedHash, nil
}

// CheckPasswordHash compares a hashed password with its plaintext version using Argon2id
func CheckPasswordHash(password, encodedSalt, expectedHash string) bool {
	if len(password) == 0 || len(encodedSalt) == 0 || len(expectedHash) == 0 {
		return false
	}
	hash := argon2.IDKey([]byte(password), []byte(encodedSalt), 1, 64*1024, 4, 32)

	return subtle.ConstantTimeCompare(hash, []byte(expectedHash)) == 1
}
