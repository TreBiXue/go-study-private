package main

import (
	"go-studying/api/router"
	"go-studying/app"
)

func InitServer() *app.Server {
	engine := app.NewGinEngine()
	db := app.InitSpannerDB()

	router.InitAccountRouter(db, engine)

	server := app.NewServer(engine)
	return server
}

func main() {
	server := InitServer()
	server.Start()

}
