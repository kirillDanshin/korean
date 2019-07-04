package hash

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Password(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func RandToken() string {
	return RandStr(32)
}

func RandStr(byteLen int) string {
	b := make([]byte, byteLen)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}
