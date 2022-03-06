package delivery

import (
	"context"

	"github.com/bagusays/efishery-technical-test/internal/app"
	"github.com/bagusays/efishery-technical-test/internal/model"
	"github.com/bagusays/efishery-technical-test/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	echoServer      *echo.Echo
	authService     *service.Auth
	resourceService *service.Resource
}

// New - creating new instance for apidb echoServer
func New(container *app.Container) Server {
	echoServer := echo.New()

	echoServer.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderAuthorization,
			echo.HeaderAccessControlAllowOrigin,
			"token",
			"Pv",
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderContentLength,
			echo.HeaderAcceptEncoding,
			echo.HeaderXCSRFToken,
			echo.HeaderXRequestID,
		},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	echoServer.Use(middleware.Recover())

	echoServer.HTTPErrorHandler = ErrorHandler
	echoServer.HideBanner = true

	h := server{
		echoServer:      echoServer,
		authService:     container.AuthService,
		resourceService: container.ResourceService,
	}

	h.routes()

	return &h
}

func (h *server) Start(port string) error {
	return h.echoServer.Start(port)
}

func (h *server) Stop(ctx context.Context) error {
	if err := h.echoServer.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func (h *server) routes() {
	h.echoServer.GET("/api/auth/validate", h.ValidateToken)
	h.echoServer.GET("/api/resources", h.FetchAllResource, AuthorizeFor(model.RoleBasic, model.RoleAdmin))
	h.echoServer.GET("/api/resources/statistics", h.ResourceStatistics, AuthorizeFor(model.RoleAdmin))
}
