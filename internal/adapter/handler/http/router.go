package handler

import (
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"os"
	"strings"
	"time"

	"github.com/asirko/go-template/internal/core/port"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Router is a wrapper for HTTP router
type Router struct {
	*gin.Engine
}

// NewRouter creates a new HTTP router
func NewRouter(
	token port.TokenService,
	userHandler UserHandler,
	authHandler AuthHandler,
) (*Router, error) {
	// Disable debug mode and write logs to file in production
	env := os.Getenv("APP_ENV")
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)

		logFile, _ := os.Create("gin.log")
		gin.DefaultWriter = io.Writer(logFile)
	}

	// CORS
	config := cors.DefaultConfig()
	allowedOrigins := os.Getenv("HTTP_ALLOWED_ORIGINS")
	originsList := strings.Split(allowedOrigins, ",")
	config.AllowOrigins = originsList

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(customLogger), gin.Recovery(), cors.New(config))

	// Custom validators
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		if err := v.RegisterValidation("user_role", userRoleValidator); err != nil {
			return nil, err
		}
	}

	// Routes management
	{
		// Swagger
		router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		v1 := router.Group("/v1")
		{
			auth := v1.Group("/auth")
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", authHandler.Login)
		}
		{
			user := v1.Group("/users").Use(authMiddleware(token))
			user.GET("/", userHandler.ListUsers)
			user.GET("/:id", userHandler.GetUser)
			{
				admin := user.Use(adminMiddleware())
				admin.PUT("/:id", userHandler.UpdateUser)
				admin.DELETE("/:id", userHandler.DeleteUser)
			}
		}
	}

	return &Router{
		router,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}

// customLogger is a custom Gin logger
func customLogger(param gin.LogFormatterParams) string {
	return fmt.Sprintf("[%s] - %s \"%s %s %s %d %s [%s]\"\n",
		param.TimeStamp.Format(time.RFC1123),
		param.ClientIP,
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency.Round(time.Millisecond),
		param.Request.UserAgent(),
	)
}
