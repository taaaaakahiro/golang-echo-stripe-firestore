package stripe

import (
	"context"
	"os"
	"stripe-example/pkg/config"
	"testing"
)

var (
	testStripe *Stripe
)

func TestMain(m *testing.M) {
	// before
	c := context.Background()
	cfg, _ := config.LoadConfig(c)

	testStripe, _ = NewStripe(c, cfg)

	res := m.Run()
	// after

	os.Exit(res)
}
