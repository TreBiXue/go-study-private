//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-studying/api/handler"
	"go-studying/api/router"
	"go-studying/app"
	"go-studying/repo"
	"go-studying/service"
)

func InitServer() *app.Server {
	wire.Build(
		app.InitSpannerDB,
		repo.NewSpannerAccountRepository,
		service.NewAccountService,
		handler.NewAccountHandlers,
		router.NewAccountRouter,

		repo.NewSpannerCenterRepository,
		service.NewCenterService,
		handler.NewCenterHandlers,
		router.NewCenterRouter,

		repo.NewSpannerVenderRepository,
		service.NewVenderService,
		handler.NewVenderHandlers,
		router.NewVenderRouter,

		app.NewServer,
		app.NewGinEngine)
	return &app.Server{}
}
