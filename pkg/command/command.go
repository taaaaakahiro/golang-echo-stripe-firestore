package command

import (
	"context"
	"fmt"
	"log"
	"stripe-example/pkg/config"
	"stripe-example/pkg/handler"

	"cloud.google.com/go/firestore"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() {
	run(context.Background())
}

func run(ctx context.Context) {
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	client, err := firestore.NewClient(ctx, cfg.ProjectID)
	if err != nil {
		log.Fatal(err)
	}
	cli, err := client.Collection(cfg.Collection).Doc(cfg.Document).Get(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// init handler
	handler := handler.NewHandler()

	fmt.Println(cli.Data())
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/", handler.StripeHealthz)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Port)))

}
