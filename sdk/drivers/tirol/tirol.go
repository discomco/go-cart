package tirol

import (
	"github.com/discomco/go-cart/sdk/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Tirol struct {
	*echo.Echo
	cfg config.IHttpConfig
}

func (t *Tirol) GetConfig() config.IHttpConfig {
	return t.cfg
}

func newTirol(cfg config.IHttpConfig) *Tirol {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	t := &Tirol{
		cfg: cfg,
	}
	t.Echo = e
	return t
}

func NewTirol(config config.IAppConfig) ITirol {
	return newTirol(config.GetHttpConfig())
}
