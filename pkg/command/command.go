package command

import (
	"context"
	"fmt"
	"log"
	firestoreClient "stripe-example/external/firestore"
	stripeClient "stripe-example/external/stripe"
	"stripe-example/pkg/config"
	"stripe-example/pkg/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() {
	run(context.Background())
}

func run(ctx context.Context) {
	// init config
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// init stripe api
	stripeService, err := stripeClient.NewStripe(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	// init firestore api
	fsService, err := firestoreClient.NewFireStore(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	// init handler
	handler := handler.NewHandler(stripeService, fsService)

	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/", handler.Healthz)
	e.POST("/subscription", handler.CreateSubscription)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))

}
