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
