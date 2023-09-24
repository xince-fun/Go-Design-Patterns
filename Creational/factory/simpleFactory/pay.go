package factory

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return &CashPM{}, nil
	case DebitCard:
		return &DebitCardPM{}, nil
	default:
		return nil, fmt.Errorf("payment method %d not recognized", m)
	}
}

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f payed using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f payed using debit card\n", amount)
}

type NewDebitCardPM struct{}

func (d *NewDebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f payed using debit card (new)\n", amount)
}
