// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"go-studying/api/handler"
	"go-studying/api/router"
	"go-studying/app"
	"go-studying/repo"
	"go-studying/service"
	"go-studying/pkg/middleware"
)

// Injectors from wire.go:

func InitServer() *app.Server {
	engine := app.NewGinEngine()
	client := app.InitSpannerDB()
	iAccountRepo := repo.NewSpannerAccountRepository(client)
	iAccountService := service.NewAccountService(iAccountRepo)
	accountHandler := handler.NewAccountHandlers(iAccountService)
	accountRouter := router.NewAccountRouter(accountHandler)
	iCenterRepo := repo.NewSpannerCenterRepository(client)
	iCenterService := service.NewCenterService(iCenterRepo)
	centerHandler := handler.NewCenterHandlers(iCenterService)
	centerRouter := router.NewCenterRouter(centerHandler)
	iVenderRepo := repo.NewSpannerVenderRepository(client)
	iVenderService := service.NewVenderService(iVenderRepo)
	venderHandler := handler.NewVenderHandlers(iVenderService)
	venderRouter := router.NewVenderRouter(venderHandler)
	iProductRepo := repo.NewSpannerProductRepository(client)
	iNyukaRepo := repo.NewSpannerNyukaRepository(client)
	iNyukaService := service.NewNyukaService(iProductRepo, iNyukaRepo, iVenderRepo)
	nyukaHandler := handler.NewNyukaHandlers(iNyukaService)
	nyukaRouter := router.NewNyukaRouter(nyukaHandler)

	iAccessLogRepo := repo.NewSpannerAccesslogRepository(client)
	engine.Use(middleware.RecordUaAndTime(iAccessLogRepo))

	server := app.NewServer(engine, accountRouter, centerRouter, venderRouter, nyukaRouter)
	return server
}
