//go:build wireinject
// +build wireinject

package main

import (
	handler "github.com/asirko/go-template/internal/adapter/handler/http"
	repo "github.com/asirko/go-template/internal/adapter/repository/postgres"
	"github.com/asirko/go-template/internal/core/port"
	"github.com/asirko/go-template/internal/core/service"
	"github.com/google/wire"
)

func InitializeEvent(token port.TokenService, db *repo.DB) (*handler.Router, error) {
	wire.Build(
		repo.NewUserRepository,
		wire.Bind(new(port.UserRepository), new(*repo.UserRepository)),

		service.NewUserService,
		wire.Bind(new(port.UserService), new(*service.UserService)),
		service.NewAuthService,
		wire.Bind(new(port.AuthService), new(*service.AuthService)),

		handler.NewUserHandler,
		handler.NewAuthHandler,
		handler.NewRouter,
	)
	return &handler.Router{}, nil
}
