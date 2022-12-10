package stripe

import (
	"stripe-example/external/stripe/domain/input"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stripe/stripe-go/customer"
)

func TestStripe_CreateCustomer(t *testing.T) {

	t.Run("ok", func(t *testing.T) {
		param := &input.Customer{Name: "user1", Email: "test@gmail.com", Phone: "012-345-6789"}
		got, err := testStripe.CreateCustomer(param)
		assert.NoError(t, err)
		assert.NotEmpty(t, got)

		assert.Equal(t, "user1", got.Name)
		assert.Equal(t, "test@gmail.com", got.Email)
		assert.Equal(t, "012-345-6789", got.Phone)

		t.Cleanup(func() {
			_, err := customer.Del(got.ID, nil)
			assert.NoError(t, err)
		})
	})

}
