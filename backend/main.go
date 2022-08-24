package main

import (
	"www.github.com/backend/dbmodels"
	"www.github.com/backend/mongodb"
	"www.github.com/backend/server"
)

func main() {
	mongodb.InitMongoDb()
	dbmodels.InitDB()
	server.Init()
}
