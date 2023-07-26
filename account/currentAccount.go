package account

import (
	h "bank/client"
	"fmt"
)

type CurrentAccount struct {
	Holder  h.Holder
	Agency  int
	Account int
	balance float64
}

func (c *CurrentAccount) Withdraw(value float64) string {

	canWithdraw := value > 0 && value <= c.balance

	if canWithdraw {
		c.balance -= value
		return "The withdraw of " + fmt.Sprintf("%.2f", value) + " was concluded"
	} else {
		return "The account don't have balance enouth. withdraw an value bellow " + fmt.Sprintf("%.2f", c.balance)
	}

}

func (c *CurrentAccount) Deposit(value float64) (string, float64) {

	if value > 0 {
		c.balance += value
		return "You deposit ", value
	} else {
		return "Deposit a value bigger than ", 0
	}

}

func (c *CurrentAccount) Transfer(value float64, destination *CurrentAccount) bool {

	canTransfer := value <= c.balance && value > 0

	if canTransfer {
		c.balance -= value
		destination.balance += value
		return true
	} else {
		return false
	}
}

func (c *CurrentAccount) GetBalance() float64{
	return c.balance
}
