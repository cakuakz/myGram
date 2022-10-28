package main

import (
	"final-project/database"
	"final-project/routers"
)

func main() {
	database.StartDB()

	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}