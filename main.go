package main

import (
	"flag"

	"zero-thinking-backend/server"

	"zero-thinking-backend/config"
)

func main() {

	env := flag.String("e", "dev", "")
	flag.Parse()

	config.Init(*env)
	// database.Init(false)
	// defer database.Close()
	if err := server.Init(); err != nil {
		panic(err)
	}
}
