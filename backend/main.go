package main

import (
	"fmt"

	"www.github.com/backend/dbmodels"
	"www.github.com/backend/server"
)

func main() {
	_, err := dbmodels.ExportAsPdf()
	fmt.Println(err)
	dbmodels.InitDB()
	server.Init()
}
