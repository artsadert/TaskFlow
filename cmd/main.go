package main

import (
	"fmt"
	"os"

	"github.com/artsadert/TaskFlow/internal/infrastructure/db/mongodb"
)

func main() {
	err := mongodb.LoadDotenv()
	if err != nil {
		fmt.Printf("error occupied: %s\n", err.Error())
	}
	fmt.Println(os.Getenv("MONGO_USERNAME"))
}
