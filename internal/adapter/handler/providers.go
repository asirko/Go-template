package handler

import (
	http "github.com/asirko/go-template/internal/adapter/handler/http"
	"github.com/google/wire"
)

var Providers = wire.NewSet(
	http.NewUserHandler,
	http.NewAuthHandler,
	http.NewRouter,
)
