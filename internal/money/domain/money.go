package domain

type Money int64

func NewMoney(amount int64) (Money, error) {
	if amount < 0 {
		return 0, ErrInvalidMoneyAmount
	}

	return Money(amount), nil
}

func FromDecimal(amount float64) (Money, error) {
	if amount < 0 {
		return 0, ErrInvalidMoneyAmount
	}

	moneyAmount := int64(amount * 100)
	return Money(moneyAmount), nil
}

func (m Money) AmountInCents() int64 {
	return int64(m)
}
