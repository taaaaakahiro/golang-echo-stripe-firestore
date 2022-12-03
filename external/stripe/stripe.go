package stripe

import (
	"context"
	"stripe-example/pkg/config"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/v74/client"
)

type Stripe struct {
	Key          string
	StripeClient *client.API
}

func NewStripe(ctx context.Context, cfg *config.EnvConfig) (*Stripe, error) {
	stripe.Key = cfg.StripeKey

	return &Stripe{
		Key:          stripe.Key,
		StripeClient: client.New(cfg.StripeKey, nil),
	}, nil
}
