package helpers


import (
"crypto/sha512"
"encoding/base64"
)

const (
passwordSalt = "!@#$%^&*()"
)

func HashPassword(email, password string) string {

hash := sha512.New()
hash.Write([]byte(passwordSalt))
hash.Write([]byte(email))
hash.Write([]byte(password))

return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}