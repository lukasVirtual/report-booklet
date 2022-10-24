package main

import (
	"www.github.com/backend/dbmodels"
	"www.github.com/backend/server"
	serverfiles "www.github.com/backend/server_files"
)

func main() {
	serverfiles.InitRedis()
	dbmodels.InitDB()
	server.Init()
}
