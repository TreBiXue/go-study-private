//go:build wireinject
// +build wireinject

package cmd

import (
	"go-studying/app"

	"github.com/google/wire"
)

func InitServer() *app.Server {
	wire.Build(
		app.GoStudyProviders)
	return &app.Server{}
}
