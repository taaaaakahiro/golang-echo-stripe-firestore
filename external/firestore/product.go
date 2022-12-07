package firestore

import (
	"context"
	"log"
	"stripe-example/pkg/domain/collection"

	"google.golang.org/api/iterator"
)

func (f *FireStore) CreateSubscription(ctx context.Context, subs *collection.Subscription) error {
	// DBに保存
	dr := f.Client.Collection(f.cfg.Collection).Doc(f.cfg.Document)
	defer f.Client.Close()
	if _, err := dr.Set(ctx, subs); err != nil {
		log.Fatalf("Failed to create subscription. err=%v", err)
	}
	return nil

}

func (f *FireStore) ListSubscription(ctx context.Context) map[string]interface{} {
	iter := f.Client.Collection(f.cfg.Collection).Documents(ctx)
	defer f.Client.Close()
	var m map[string]interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		m = doc.Data()
	}

	return m

}
