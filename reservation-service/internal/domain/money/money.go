package money

type Money int64

func FromCents(amount int64) (Money, error) {
	if err := validateMoney(amount); err != nil {
		return 0, err
	}

	return Money(amount), nil
}

func FromDecimal(amount float64) (Money, error) {
	amountInCents := int64(amount * 100)
	if err := validateMoney(amountInCents); err != nil {
		return 0, err
	}

	return Money(amountInCents), nil
}

func (m Money) AmountInCents() int64 {
	return int64(m)
}

func validateMoney(amount int64) error {
	if amount < 0 {
		return ErrInvalidMoneyAmount
	}

	return nil
}
