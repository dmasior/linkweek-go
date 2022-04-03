//go:build wireinject
// +build wireinject

package db

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"linkweek-go/config"
)

func Build() *gorm.DB {
	panic(wire.Build(Provider, config.ProviderSet))
}
