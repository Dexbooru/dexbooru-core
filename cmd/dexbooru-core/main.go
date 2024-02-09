package main

import (
	"fmt"
	"os"

	"github.com/Dexbooru/dexbooru-core/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	serverPort := os.Getenv("SERVER_PORT")
	serverAddr := fmt.Sprintf(":%s", serverPort)

	app := routes.CreateApplicationRouter()	
	app.Run(serverAddr)
}
