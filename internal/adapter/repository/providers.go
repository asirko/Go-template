package repository

import (
	postgres "github.com/asirko/go-template/internal/adapter/repository/postgres"
	"github.com/asirko/go-template/internal/core/port"
	"github.com/google/wire"
)

var Providers = wire.NewSet(
	postgres.NewUserRepository,
	wire.Bind(new(port.UserRepository), new(*postgres.UserRepository)),
)
