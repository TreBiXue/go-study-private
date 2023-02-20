package main

import "go-studying/api/routers"

func InitServer() *Server {
	engine := NewGinEngine()
	db := InitSpannerDB()

	routers.InitAccountRouter(db, engine)

	server := NewServer(engine)
	return server
}

func main() {
	server := InitServer()
	server.Start()
}
