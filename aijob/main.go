package main

import (
	"aijob/db"
	"aijob/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":3000"))

}
