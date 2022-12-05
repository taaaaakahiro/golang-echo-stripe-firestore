package firestore

import (
	"context"
	"log"
	"stripe-example/pkg/config"

	"cloud.google.com/go/firestore"
)

type FireStore struct {
	cfg    *config.EnvConfig
	Client *firestore.Client
}

func NewFireStore(ctx context.Context, cfg *config.EnvConfig) (*FireStore, error) {
	client, err := firestore.NewClient(ctx, cfg.ProjectID)
	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
	}

	fs := &FireStore{
		cfg:    cfg,
		Client: client,
	}
	return fs, nil
}
