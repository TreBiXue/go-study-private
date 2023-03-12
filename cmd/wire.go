//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-studying/app"
)

func InitServer() *app.Server {
	wire.Build(
		app.GoStudyProviders)
	return &app.Server{}
}
