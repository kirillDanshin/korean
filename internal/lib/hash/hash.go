package hash

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Password hashing.
func Password(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash - check if the password matched the hash.
func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// RandToken - generate rand token.
func RandToken() string {
	return RandStr(32)
}

// RandStr - generate hash by len.
func RandStr(byteLen int) string {
	b := make([]byte, byteLen)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}
