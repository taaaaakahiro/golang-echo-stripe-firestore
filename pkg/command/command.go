package command

import (
	"context"
	"fmt"
	"log"
	"os"
	firestoreClient "stripe-example/external/firestore"
	stripeClient "stripe-example/external/stripe"
	"stripe-example/pkg/config"
	"stripe-example/pkg/handler"
	"stripe-example/pkg/server"
)

const exitOK = 0

func Run() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
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
	handler, err := handler.NewHandler(stripeService, fsService)
	if err != nil {
		log.Fatal(err)
	}

	// init server
	server, err := server.NewServer(handler)
	if err != nil {
		log.Fatal(err)
	}
	// start
	server.Echo.Logger.Fatal(server.Echo.Start(fmt.Sprintf(":%d", cfg.Port)))

	return exitOK
}
