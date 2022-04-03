//go:build wireinject
// +build wireinject

package config

import "github.com/google/wire"

func Build() Config {
	panic(wire.Build(Provider))
}
