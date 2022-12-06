package command

import (
	"context"
	"fmt"
	"log"
	firestoreClient "stripe-example/external/firestore"
	stripeClient "stripe-example/external/stripe"
	"stripe-example/pkg/config"
	"stripe-example/pkg/handler"
	"stripe-example/pkg/server"
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

}
