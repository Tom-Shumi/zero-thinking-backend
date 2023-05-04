package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"zero-thinking-backend/database"
	"zero-thinking-backend/server"

	"zero-thinking-backend/config"
)

func main() {

	env := flag.String("e", "dev", "")
	flag.Parse()

	config.Init(*env)

	err := godotenv.Load(fmt.Sprintf("env/%s.env", *env))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Init()
	defer database.Close()
	if err := server.Init(); err != nil {
		panic(err)
	}
}
