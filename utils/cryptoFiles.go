package utils

import (
	"github.com/matthewhartstonge/argon2"
)

func Encrypt(plaintext string) string {
	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(plaintext))
	CheckErrors(err, "Code 2", "Error crypting the text", "Report the error in the github or write the value again")
	return string(encoded)
}

func Decrypt(plaintext, hash string) bool {
	ok, err := argon2.VerifyEncoded([]byte(plaintext), []byte(hash))
	CheckErrors(err, "Code 2", "Error in the verification of the encodes", "Report the error on the github or check the values :D")
	if ok {
		return true
	} else {
		return false
	}
}
