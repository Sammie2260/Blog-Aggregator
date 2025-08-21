package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 200 OK response
func Success(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, data)
}

// 500 Internal Server Error response
func Error(c echo.Context, message string) error {
	return c.JSON(http.StatusInternalServerError, map[string]string{"error": message})
}

// 404 Not Found response
func NotFound(c echo.Context, message string) error {
	return c.JSON(http.StatusNotFound, map[string]string{"error": message})
}

// 400 Bad Request response
func BadRequest(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": message})
}

