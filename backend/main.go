package main

import (
	"www.github.com/backend/dbmodels"
	"www.github.com/backend/server"
	servercom "www.github.com/backend/server_communication"
	serverfiles "www.github.com/backend/server_files"
)

func main() {
	servercom.InitRedis()
	dbmodels.InitDB()
	serverfiles.InitMongoDb()
	server.Init()
}
