package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Port Sets the port from an environment file and returns it as a string.
func Port() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading the env file.")
	}

	port := os.Getenv("PORT")
	if port != "8150" {
		port = "8150"
	}
	return ":" + port
}
