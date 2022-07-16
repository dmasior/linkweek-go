//go:build wireinject
// +build wireinject

package legacy

import (
	"database/sql"
	"linkweek-go/config"

	"github.com/google/wire"
)

func Build() *sql.DB {
	panic(wire.Build(Provider, config.ProviderSet))
}
