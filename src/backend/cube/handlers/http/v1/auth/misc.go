package auth

import "crypto/sha256"

func getPasswordHash(password string) []byte {
	pwd := sha256.New()
	pwd.Write([]byte(password))
	return pwd.Sum(nil)
}
