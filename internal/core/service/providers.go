package service

import (
	"github.com/asirko/go-template/internal/core/port"
	"github.com/google/wire"
)

var Providers = wire.NewSet(
	NewUserService,
	wire.Bind(new(port.UserService), new(*UserService)),
	NewAuthService,
	wire.Bind(new(port.AuthService), new(*AuthService)),
)
