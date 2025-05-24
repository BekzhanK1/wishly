package main

import (
	"github.com/BekzhanK1/wishly/internal/server"
)

func main() {
	server.LoadEnv()
	db := server.ConnectToDatabase()
	r := server.SetupRouter(db)

	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
