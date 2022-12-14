package stripe

import (
	"stripe-example/pkg/domain/output"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentmethod"
	"github.com/stripe/stripe-go/v74/price"
	"github.com/stripe/stripe-go/v74/product"
)

func (s *Stripe) CreatePaymentMethod() (*stripe.PaymentMethod, error) {
	stripe.Key = s.Key
	params := &stripe.PaymentMethodParams{
		Card: &stripe.PaymentMethodCardParams{
			Number:   stripe.String("4242424242424242"),
			ExpMonth: stripe.Int64(8),
			ExpYear:  stripe.Int64(2022),
			CVC:      stripe.String("314"),
		},
		Type: stripe.String("card"),
	}
	pm, _ := paymentmethod.New(params)

	return pm, nil
}

func (s *Stripe) ListPaymentMethods() ([]*stripe.PaymentMethod, error) {
	stripe.Key = s.Key
	params := &stripe.PaymentMethodListParams{

		// Customer: stripe.String("cus_60uWheDVvFJCig"),
		// Type: stripe.String("card"),
	}
	i := paymentmethod.List(params)
	pm := make([]*stripe.PaymentMethod, 0)
	for i.Next() {
		pm = append(pm, i.PaymentMethod())
	}

	return pm, nil
}

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
