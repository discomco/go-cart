package middlewares

import "github.com/labstack/echo/v4"

type IMiddlewareManager interface {
	RequestLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}
