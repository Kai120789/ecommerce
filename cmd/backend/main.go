package main

import (
	"shop/internal/db/connect"
	"shop/internal/db/create"
)

func main() {

	create.CreateTables()
	defer connect.CloseDB()
}
