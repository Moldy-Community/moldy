package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetDotEnv() string {
	err := godotenv.Load()
	CheckErrors(err, "code 2", "Error reading the enviroment variables", "Report the error in the Github :D")

	return os.Getenv("CRYPTO_KEY")
}
