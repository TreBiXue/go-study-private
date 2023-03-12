package app

import (
	"github.com/google/wire"
	"go-studying/api/handler"
	"go-studying/api/router"
	"go-studying/repo"
	"go-studying/service"
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

var GoStudyProviders = wire.NewSet(
	InitSpannerDB,
	AccountProviders,
	CenterProviders,
	VenderProviders,
	NewServer,
	NewGinEngine,
)
