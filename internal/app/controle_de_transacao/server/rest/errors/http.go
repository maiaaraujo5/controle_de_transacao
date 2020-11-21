package errors

import (
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle_de_transacao/internal/app/controle_de_transacao/errors"
	"net/http"
)

type httpError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func NewErrorResponse(c echo.Context, status int, message string) error {
	return c.JSON(status, newHttpError(status, message))
}

func ToErrorResponse(c echo.Context, err error) error {
	switch {
	case errors.IsNotFound(err):
		return c.JSON(http.StatusNotFound, newHttpError(http.StatusNotFound, err.Error()))
	case errors.IsBadRequest(err):
		return c.JSON(http.StatusBadRequest, newHttpError(http.StatusBadRequest, err.Error()))
	case errors.IsAlreadyExists(err):
		return c.JSON(http.StatusConflict, newHttpError(http.StatusConflict, err.Error()))
	default:
		return c.JSON(http.StatusInternalServerError, newHttpError(http.StatusInternalServerError, "internal server error"))
	}
}

func newHttpError(statusCode int, message string) *httpError {
	return &httpError{
		StatusCode: statusCode,
		Message:    message,
	}
}
