package stripe

import (
	"stripe-example/external/stripe/domain/input"

	"github.com/pkg/errors"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

func (s *Stripe) GetCustomer() {

}

func (s *Stripe) CreateCustomer(p *input.Customer) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{
		Name:  &p.Name,
		Email: &p.Email,
		Phone: &p.Phone,
	}
	c, err := customer.New(params)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return c, nil

}
