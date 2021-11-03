package main

import (
	"log"

	"bitbucket.org/isbtotogroup/sdsb4d-apifrontend/db"
	"bitbucket.org/isbtotogroup/sdsb4d-apifrontend/routers"
)

func main() {
	db.Init()
	app := routers.Init()
	log.Fatal(app.Listen(":7071"))
}
