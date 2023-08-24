package dotenv

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvironmentVariable(key string) string {
	err := godotenv.Load()

	if err != nil {
		panic("Error when loading .env file")
	}

	return os.Getenv(key)
}
