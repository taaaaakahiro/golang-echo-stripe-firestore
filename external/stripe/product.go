package stripe

import (
	"github.com/pkg/errors"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/product"
)

func (s *Stripe) GetProduct(productID string) (*stripe.Product, error) {
	p, err := product.Get(productID, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return p, nil
}

func (s *Stripe) CreteaProduct(title string) (*stripe.Product, *stripe.Price, error) {
	// product
	productParams := &stripe.ProductParams{
		Name:                stripe.String(title),
		StatementDescriptor: stripe.String("Chompy"), // 明細書に記載する文字列. 5 ~ 22文字でアルファベットと数字のみなので注意. https://stripe.com/docs/statement-descriptors
	}
	// productParams.AddMetadata("subscription_id", sub.ID)
	// productParams.AddMetadata("plan_id", plan.ID)
	product, _ := s.Client.Products.New(productParams)

	// price
	priceParams := &stripe.PriceParams{
		Currency: stripe.String(string(stripe.CurrencyJPY)), // 通貨の設定, JPYを設定する
		Product:  stripe.String(product.ID),                 // 上記で作成したProductのIDを設定する
		Recurring: &stripe.PriceRecurringParams{ // サブスク期間の設定
			Interval:      stripe.String("day"), // 日毎
			IntervalCount: stripe.Int64(30),     // 30日
		},
		UnitAmount: stripe.Int64(3000), // 料金, 3000円
	}
	// priceParams.AddMetadata("subscription_id", sub.ID)
	// priceParams.AddMetadata("plan_id", plan.ID)
	price, _ := s.Client.Prices.New(priceParams)

	return product, price, nil
}
