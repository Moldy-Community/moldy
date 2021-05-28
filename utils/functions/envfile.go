package functions

import (
	"os"

	"github.com/joho/godotenv"
)

func GetDotEnv(key string) string {
	err := godotenv.Load()
	CheckErrors(err, "code 2", "Error reading the enviroment variables", "Report the error in the Github :p")

	return os.Getenv(key)
}
