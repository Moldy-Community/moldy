package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetDotEnv() string {
	err := godotenv.Load()
	CheckErrors(err, "2", "Error in the read the dot env files")

	return os.Getenv("CRYPTO_KEY")
}
