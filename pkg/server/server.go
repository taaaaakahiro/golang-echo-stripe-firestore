package server

import (
	"stripe-example/pkg/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	Echo    *echo.Echo
	handler *handler.Handler
}

func NewServer(handler *handler.Handler) (*Server, error) {
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// endpoint
	e.GET("/", handler.Product.Healthz)

	s := e.Group("subscript")
	s.POST("/subs", handler.Product.CreateSubscription)

	c := e.Group("customer")
	c.GET("", handler.Customer.ListCustomer)

	return &Server{
		Echo:    e,
		handler: handler,
	}, nil
}
