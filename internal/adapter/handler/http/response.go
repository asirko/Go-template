package handler

import (
	"net/http"
	"time"

	"github.com/asirko/go-template/internal/core/domain"
	"github.com/asirko/go-template/internal/core/port"
	"github.com/gin-gonic/gin"
)

// errorResponse represents an error response body format
type errorResponse struct {
	Message string `json:"message" example:"Error message"`
	Code    string `json:"code" example:"1"`
}

// newErrorResponse is a helper function to create an error response body
func newErrorResponse(message, code string) errorResponse {
	return errorResponse{
		Message: message,
		Code:    code,
	}
}

// meta represents metadata for a paginated response
type meta struct {
	Total uint64 `json:"total" example:"100"`
	Limit uint64 `json:"limit" example:"10"`
	Skip  uint64 `json:"skip" example:"0"`
}

// newMeta is a helper function to create metadata for a paginated response
func newMeta(total, limit, skip uint64) meta {
	return meta{
		Total: total,
		Limit: limit,
		Skip:  skip,
	}
}

// authResponse represents an authentication response body
type authResponse struct {
	AccessToken string `json:"token" example:"v2.local.Gdh5kiOTyyaQ3_bNykYDeYHO21Jg2..."`
}

// newAuthResponse is a helper function to create a response body for handling authentication data
func newAuthResponse(token string) authResponse {
	return authResponse{
		AccessToken: token,
	}
}

// userResponse represents a user response body
type userResponse struct {
	ID        uint64    `json:"id" example:"1"`
	Name      string    `json:"name" example:"John Doe"`
	Email     string    `json:"email" example:"test@example.com"`
	CreatedAt time.Time `json:"created_at" example:"1970-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"1970-01-01T00:00:00Z"`
}

// newUserResponse is a helper function to create a response body for handling user data
func newUserResponse(user *domain.User) userResponse {
	return userResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// errorStatusMap is a map of defined error messages and their corresponding http status codes
var errorStatusMap = map[error]int{
	port.ErrDataNotFound:               http.StatusNotFound,
	port.ErrConflictingData:            http.StatusConflict,
	port.ErrInvalidCredentials:         http.StatusUnauthorized,
	port.ErrUnauthorized:               http.StatusUnauthorized,
	port.ErrEmptyAuthorizationHeader:   http.StatusUnauthorized,
	port.ErrInvalidAuthorizationHeader: http.StatusUnauthorized,
	port.ErrInvalidAuthorizationType:   http.StatusUnauthorized,
	port.ErrInvalidToken:               http.StatusUnauthorized,
	port.ErrExpiredToken:               http.StatusUnauthorized,
	port.ErrForbidden:                  http.StatusForbidden,
	port.ErrNoUpdatedData:              http.StatusBadRequest,
	port.ErrInsufficientStock:          http.StatusBadRequest,
	port.ErrInsufficientPayment:        http.StatusBadRequest,
}

// validationError sends an error response for some specific request validation error
func validationError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, err)
}

// handleError determines the status code of an error and returns a JSON response with the error message and status code
func handleError(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errRsp := newErrorResponse(err.Error(), "TODO")

	ctx.JSON(statusCode, errRsp)
}

// handleAbort sends an error response and aborts the request with the specified status code and error message
func handleAbort(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	rsp := newErrorResponse(err.Error(), "TODO")
	ctx.AbortWithStatusJSON(statusCode, rsp)
}
