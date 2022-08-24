package tirol

import (
	context2 "context"
	"io/fs"
	"net"
	"net/http"

	"github.com/discomco/go-cart/config"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
)

type ITirol interface {
	GetConfig() config.IHttpConfig
	NewContext(r *http.Request, w http.ResponseWriter) echo.Context
	Router() *echo.Router
	Routers() map[string]*echo.Router
	DefaultHTTPErrorHandler(err error, c echo.Context)
	Pre(middleware ...echo.MiddlewareFunc)
	Use(middleware ...echo.MiddlewareFunc)
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	Any(path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route
	Match(methods []string, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) []*echo.Route
	File(path, file string, m ...echo.MiddlewareFunc) *echo.Route
	Add(method, path string, handler echo.HandlerFunc, middleware ...echo.MiddlewareFunc) *echo.Route
	Host(name string, m ...echo.MiddlewareFunc) (g *echo.Group)
	Group(prefix string, m ...echo.MiddlewareFunc) (g *echo.Group)
	URI(handler echo.HandlerFunc, params ...interface{}) string
	URL(h echo.HandlerFunc, params ...interface{}) string
	Reverse(name string, params ...interface{}) string
	Routes() []*echo.Route
	AcquireContext() echo.Context
	ReleaseContext(c echo.Context)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	Start(address string) error
	StartTLS(address string, certFile, keyFile interface{}) (err error)
	StartAutoTLS(address string) error
	StartServer(s *http.Server) (err error)
	ListenerAddr() net.Addr
	TLSListenerAddr() net.Addr
	StartH2CServer(address string, h2s *http2.Server) error
	Close() error
	Shutdown(ctx context2.Context) error
	Static(prefix, root string) *echo.Route
	StaticFS(pathPrefix string, filesystem fs.FS) *echo.Route
	FileFS(path, file string, filesystem fs.FS, m ...echo.MiddlewareFunc) *echo.Route
}
