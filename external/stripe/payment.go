package stripe

import (
	"stripe-example/pkg/domain/output"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/price"
	"github.com/stripe/stripe-go/v74/product"
)

func (s *Stripe) CreatePayment(id string) (*output.Payment, error) {
	product_params := &stripe.ProductParams{
		ID:          &id,
		Name:        stripe.String("Starter Subscription"),
		Description: stripe.String("$12/Month subscription"),
	}
	starter_product, err := product.New(product_params)
	if err != nil {
		return nil, err
	}

	price_params := &stripe.PriceParams{
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Product:  stripe.String(starter_product.ID),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		UnitAmount: stripe.Int64(1200),
	}
	starter_price, _ := price.New(price_params)

	payment := output.Payment{ProductID: starter_price.ID, PriceID: starter_price.ID}

	return &payment, nil

}
