package stripe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripe_CreatePaymentMethod(t *testing.T) {
	//TODO: add test
	t.Run("ok", func(t *testing.T) {
		_, err := testStripe.CreatePaymentMethod()
		assert.NoError(t, err)

	})

}
func TestStripe_ListPaymentMethods(t *testing.T) {
	//TODO: add test
	t.Run("ok", func(t *testing.T) {
		_, err := testStripe.ListPaymentMethods()
		assert.NoError(t, err)

	})

}
