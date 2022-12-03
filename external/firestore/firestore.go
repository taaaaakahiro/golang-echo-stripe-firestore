package firestore

import (
	"context"
	"log"
	"stripe-example/pkg/config"
	"stripe-example/pkg/domain/collection"

	"cloud.google.com/go/firestore"
)

type FireStore struct {
	ctx      context.Context
	cfg      *config.EnvConfig
	FsClient *firestore.Client
}

func NewFireStore(ctx context.Context, cfg *config.EnvConfig) (*FireStore, error) {
	fsClient, err := firestore.NewClient(context.Background(), cfg.ProjectID)
	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
	}

	return &FireStore{
		ctx:      ctx,
		cfg:      cfg,
		FsClient: fsClient,
	}, nil

}

func (f *FireStore) CreateSubscription(subs *collection.Subscription) error {
	// DBに保存
	dr := f.FsClient.Collection(f.cfg.Collection).Doc(f.cfg.Document)
	if _, err := dr.Set(f.ctx, subs); err != nil {
		log.Fatalf("Failed to create subscription. err=%v", err)
	}
	return nil

}
