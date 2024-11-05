package customecho

import (
	"context"
	"fmt"
	"github.com/duongbuidinh600/core-package/core/constants"
	"github.com/duongbuidinh600/core-package/core/http/customecho/config"
	"github.com/duongbuidinh600/core-package/core/http/customecho/contracts"
	ipratelimit "github.com/duongbuidinh600/core-package/core/http/customecho/middlewares/ip_ratelimit"
	"github.com/duongbuidinh600/core-package/core/http/customecho/middlewares/log"
	problemdetail "github.com/duongbuidinh600/core-package/core/http/customecho/middlewares/problem_detail"
	"github.com/labstack/echo/v4/middleware"

	handlers "github.com/duongbuidinh600/core-package/core/http/customecho/handlers"

	"github.com/duongbuidinh600/core-package/core/logger"
	"github.com/labstack/echo/v4"
	"strings"
)

type echoHttpServer struct {
	echo         *echo.Echo
	config       *config.EchoHttpOptions
	log          logger.Logger
	routeBuilder *contracts.RouteBuilder
}

func NewEchoHttpServer(
	config *config.EchoHttpOptions,
	logger logger.Logger,
) contracts.EchoHttpServer {
	e := echo.New()
	e.HideBanner = true

	return &echoHttpServer{
		echo:         e,
		config:       config,
		log:          logger,
		routeBuilder: contracts.NewRouteBuilder(e),
	}
}

func (s *echoHttpServer) RunHttpServer(
	configEcho ...func(echo *echo.Echo),
) error {
	s.echo.Server.ReadTimeout = constants.ReadTimeout
	s.echo.Server.WriteTimeout = constants.WriteTimeout
	s.echo.Server.MaxHeaderBytes = constants.MaxHeaderBytes

	if len(configEcho) > 0 {
		ehcoFunc := configEcho[0]
		if ehcoFunc != nil {
			configEcho[0](s.echo)
		}
	}
	return s.echo.Start(s.config.Port)
}

func (s *echoHttpServer) Logger() logger.Logger {
	return s.log
}

func (s *echoHttpServer) Cfg() *config.EchoHttpOptions {
	return s.config
}

func (s *echoHttpServer) RouteBuilder() *contracts.RouteBuilder {
	return s.routeBuilder
}

func (s *echoHttpServer) ConfigGroup(
	groupName string,
	groupFunc func(group *echo.Group),
) {
	groupFunc(s.echo.Group(groupName))
}

func (s *echoHttpServer) AddMiddlewares(middlewares ...echo.MiddlewareFunc) {
	if len(middlewares) > 0 {
		s.echo.Use(middlewares...)
	}
}

func (s *echoHttpServer) GracefulShutdown(ctx context.Context) error {
	err := s.echo.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *echoHttpServer) ApplyVersioningFromHeader() {
	s.echo.Pre(apiVersion)
}

func (s *echoHttpServer) GetEchoInstance() *echo.Echo {
	return s.echo
}

// APIVersion Header Based Versioning
func apiVersion(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		headers := req.Header

		apiVersion := headers.Get("version")

		req.URL.Path = fmt.Sprintf("/%s%s", apiVersion, req.URL.Path)

		return next(c)
	}
}

func (s *echoHttpServer) SetupDefaultMiddlewares() {
	skipper := func(c echo.Context) bool {
		return strings.Contains(c.Request().URL.Path, "swagger") ||
			strings.Contains(c.Request().URL.Path, "metrics") ||
			strings.Contains(c.Request().URL.Path, "health") ||
			strings.Contains(c.Request().URL.Path, "favicon.ico")
	}

	s.echo.HTTPErrorHandler = func(err error, c echo.Context) {
		// bypass skip endpoints and its error
		if skipper(c) {
			return
		}

		handlers.ProblemDetailErrorHandlerFunc(err, c, s.log)
	}

	s.echo.Use(
		log.EchoLogger(
			s.log,
			log.WithSkipper(skipper),
		),
	)

	s.echo.Use(middleware.BodyLimit(constants.BodyLimit))
	s.echo.Use(ipratelimit.IPRateLimit())
	s.echo.Use(middleware.RequestID())
	s.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level:   constants.GzipLevel,
		Skipper: skipper,
	}))
	// should be last middleware
	s.echo.Use(problemdetail.ProblemDetail(problemdetail.WithSkipper(skipper)))
}
