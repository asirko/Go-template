package main

import (
	"context"
	"log/slog"
	"os"

	_ "github.com/asirko/go-template/docs"
	repo "github.com/asirko/go-template/internal/adapter/repository/postgres"
	token "github.com/asirko/go-template/internal/adapter/token/paseto"
	"github.com/joho/godotenv"
)

func init() {
	// Init logger
	var logHandler *slog.JSONHandler

	env := os.Getenv("APP_ENV")
	if env == "production" {
		logHandler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})

		// Load .env file
		err := godotenv.Load()
		if err != nil {
			slog.Error("Error loading .env file", "error", err)
			os.Exit(1)
		}
	}

	logger := slog.New(logHandler)
	slog.SetDefault(logger)
}

// @title						Go Template
// @version					1.0
// @description				This is a simple RESTful Service API written in Go using Gin web framework and PostgreSQL database.
//
// @contact.name				Alexander Sirko
// @contact.url				https://github.com/asirko/go-template
// @contact.email				sirko.alexandre@gmail.com
//
// @license.name				MIT
// @license.url				https://github.com/asirko/go-template/blob/main/LICENSE
//
// @BasePath					/v1
// @schemes					http https
//
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and the access token.
func main() {
	appName := os.Getenv("APP_NAME")
	env := os.Getenv("APP_ENV")
	dbConn := os.Getenv("DB_CONNECTION")
	tokenSymmetricKey := os.Getenv("TOKEN_SYMMETRIC_KEY")
	httpUrl := os.Getenv("HTTP_URL")
	httpPort := os.Getenv("HTTP_PORT")
	listenAddr := httpUrl + ":" + httpPort

	slog.Info("Starting the application toto", "app", appName, "env", env)

	// Init database
	ctx := context.Background()
	db, err := repo.NewDB(ctx)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("Successfully connected to the database", "db", dbConn)

	// Init tokenValue service
	tokenValue, err := token.NewToken(tokenSymmetricKey)
	if err != nil {
		slog.Error("Error initializing tokenValue service", "error", err)
		os.Exit(1)
	}

	router, err := InitializeEvent(tokenValue, db)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
