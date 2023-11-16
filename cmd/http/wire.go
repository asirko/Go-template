//go:build wireinject
// +build wireinject

package main

import (
	handler "github.com/asirko/go-template/internal/adapter/handler"
	http "github.com/asirko/go-template/internal/adapter/handler/http"
	repo "github.com/asirko/go-template/internal/adapter/repository"
	db "github.com/asirko/go-template/internal/adapter/repository/postgres"
	"github.com/asirko/go-template/internal/core/port"
	"github.com/asirko/go-template/internal/core/service"
	"github.com/google/wire"
)

func InitializeEvent(token port.TokenService, db *db.DB) (*http.Router, error) {
	wire.Build(
		repo.Providers,
		service.Providers,
		handler.Providers,
	)
	return &http.Router{}, nil
}
