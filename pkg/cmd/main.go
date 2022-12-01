package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"stripe-example/pkg/config"

	"cloud.google.com/go/firestore"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/price"
	"github.com/stripe/stripe-go/v74/product"
)

var (
	fsClient *firestore.DocumentSnapshot
)

func main() {
	ctx := context.Background()
	conf := config.NewFSConfig()
	client, err := firestore.NewClient(ctx, conf.ProjectID)
	if err != nil {
		log.Fatal(err)
	}
	cli, err := client.Collection(conf.Collection).Doc(conf.Document).Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cli.Data())
	fsClient = cli
	Init()

}

func Init() {
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
