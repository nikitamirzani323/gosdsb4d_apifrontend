package main

import (
	"log"

	"github.com/nikitamirzani323/gosdsb4d_apifrontend/db"
	"github.com/nikitamirzani323/gosdsb4d_apifrontend/routers"
)

func main() {
	db.Init()
	app := routers.Init()
	log.Fatal(app.Listen(":7071"))
}
