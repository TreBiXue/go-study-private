package app

import (
	"go-studying/api/handler"
	"go-studying/api/router"
	"go-studying/repo"
	"go-studying/service"

	"github.com/google/wire"
)

// AccountProviders ログイン
var AccountProviders = wire.NewSet(
	repo.NewSpannerAccountRepository,
	service.NewAccountService,
	handler.NewAccountHandlers,
	router.NewAccountRouter,
)

// CenterProviders センター
var CenterProviders = wire.NewSet(
	repo.NewSpannerCenterRepository,
	service.NewCenterService,
	handler.NewCenterHandlers,
	router.NewCenterRouter,
)

// VenderProviders ベンダー
var VenderProviders = wire.NewSet(
	repo.NewSpannerVenderRepository,
	service.NewVenderService,
	handler.NewVenderHandlers,
	router.NewVenderRouter,
)

// NyukaProviders ベンダー
var NyukaProviders = wire.NewSet(
	repo.NewSpannerNyukaRepository,
	repo.NewSpannerProductRepository,
	// repo.NewSpannerVenderRepository,
	service.NewNyukaService,
	handler.NewNyukaHandlers,
	router.NewNyukaRouter,
)

var GoStudyProviders = wire.NewSet(
	InitSpannerDB,
	AccountProviders,
	CenterProviders,
	VenderProviders,
	NyukaProviders,
	NewServer,
	NewGinEngine,
)
