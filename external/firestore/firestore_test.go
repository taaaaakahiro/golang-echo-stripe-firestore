package firestore

import (
	"context"
	"os"
	"stripe-example/pkg/config"
	"testing"
)

var (
	fsclient *FireStore
)

func TestMain(m *testing.M) {
	// before
	c := context.Background()
	cfg, _ := config.LoadConfig(c)

	fsclient, _ = NewFireStore(c, cfg)

	res := m.Run()
	// after

	os.Exit(res)
}
