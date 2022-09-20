package main

import (
	"www.github.com/backend/dbmodels"
	"www.github.com/backend/server"
)

func main() {
	dbmodels.ExportAsPdf()
	dbmodels.InitDB()
	server.Init()
}
