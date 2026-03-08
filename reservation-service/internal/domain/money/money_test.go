package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoney_FromCents(t *testing.T) {
	t.Run("should return error when the amount is less than 0", func(t *testing.T) {
		money, err := FromCents(-1)
		assert.Equal(t, Money(0), money)
		assert.Equal(t, ErrInvalidMoneyAmount, err)
	})

	t.Run("should return money when the amount is valid", func(t *testing.T) {
		money, err := FromCents(150)
		assert.Equal(t, Money(150), money)
		assert.Nil(t, err)
	})
}

func TestMoney_FromDecimal(t *testing.T) {
	t.Run("should return error when the amount is less than 0", func(t *testing.T) {
		money, err := FromDecimal(-10.0)
		assert.Equal(t, Money(0), money)
		assert.Equal(t, ErrInvalidMoneyAmount, err)
	})

	t.Run("should return money when the amount is valid", func(t *testing.T) {
		money, err := FromDecimal(15.50)
		assert.Equal(t, Money(1550), money)
		assert.Nil(t, err)
	})
}
