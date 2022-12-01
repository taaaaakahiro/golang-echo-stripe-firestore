package command

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"stripe-example/pkg/config"

	"cloud.google.com/go/firestore"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/price"
	"github.com/stripe/stripe-go/v74/product"
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
	fmt.Println(cli.Data())
	http.HandleFunc("/", stripeHealthz)
	log.Printf("Start REST API Server PORT:%d", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))

}

func stripeHealthz(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("SECRET_KEY")

	product_params := &stripe.ProductParams{
		Name:        stripe.String("Starter Subscription"),
		Description: stripe.String("$12/Month subscription"),
	}
	starter_product, _ := product.New(product_params)

	price_params := &stripe.PriceParams{
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Product:  stripe.String(starter_product.ID),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		UnitAmount: stripe.Int64(1200),
	}
	starter_price, _ := price.New(price_params)

	fmt.Println("Success! Here is your starter subscription product id: " + starter_product.ID)
	fmt.Println("Success! Here is your starter subscription price id: " + starter_price.ID)

}
