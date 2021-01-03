package database

import (
	"crypto/sha256"
	"encoding/hex"
)

//EncryptPassword get hashed password
func EncryptPassword(pw string) string {
	return hash(pw + salt)
}

//IsEqualPassword Check if given passwords are equal
func IsEqualPassword(encrypted, decrypted string) bool {
	return encrypted == EncryptPassword(decrypted)
}

func hash(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	b := hash.Sum(nil)
	return hex.EncodeToString(b)
}
