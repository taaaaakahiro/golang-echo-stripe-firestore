package firestore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// todo: 途中
func TestFirestore_ListSubscription(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got := fsclient.ListSubscription(context.Background())
		assert.NotEmpty(t, got)
	})
}
