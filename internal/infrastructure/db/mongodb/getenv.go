package mongodb

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadDotenv() error {
	err := godotenv.Load("./../../../../../.env.test")
	if err != nil {
		//log.Fatal("Could not load dotenv file!!!!!")
		return err
	}

	if os.Getenv("MONGO_USERNAME") == "" {
		log.Fatal("could not find MONGO_USERNAME in .env")
		return fmt.Errorf("could not find MONGODB_USERNAME in .env")
	}

	if os.Getenv("MONGO_PASSWORD") == "" {
		log.Fatal("could not find MONGO_PASSWORD in .env")
		return fmt.Errorf("could not find MONGO_PASSWORD in .env")
	}

	return nil
}
