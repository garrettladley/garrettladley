package xerr

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/garrettladley/garrettladley/pkg/xslog"
	"github.com/gofiber/fiber/v2"
)

type APIError struct {
	StatusCode int `json:"statusCode"`
	Message    any `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d %v", e.StatusCode, e.Message)
}

func NewAPIError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}

func BadRequest(err error) APIError {
	return NewAPIError(http.StatusBadRequest, err)
}

func InvalidJSON() APIError {
	return NewAPIError(http.StatusBadRequest, errors.New("invalid JSON request data"))
}

func NotFound(title string, withKey string, withValue any) APIError {
	return NewAPIError(http.StatusNotFound, fmt.Errorf("%s with %s='%s' not found", title, withKey, withValue))
}

func Conflict(title string, withKey string, withValue any) APIError {
	return NewAPIError(http.StatusConflict, fmt.Errorf("conflict: %s with %s='%s' already exists", title, withKey, withValue))
}

func InvalidRequestData(errors map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    errors,
	}
}

func InternalServerError() APIError {
	return NewAPIError(http.StatusInternalServerError, errors.New("internal server error"))
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	var apiErr APIError
	if castedErr, ok := err.(APIError); ok {
		apiErr = castedErr
	} else {
		apiErr = InternalServerError()
	}

	slog.LogAttrs(
		c.Context(),
		slog.LevelError,
		"HTTP API error",
		xslog.Error(err),
		slog.String("method", c.Method()),
		slog.String("path", c.Path()),
	)

	return c.Status(apiErr.StatusCode).JSON(apiErr)
}
