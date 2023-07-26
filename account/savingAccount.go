package account

import (
	h "bank/client"
	"fmt"
)

type SavingAccount struct {
	Holder                     h.Holder
	Agency, Account, Operation int
	balance                    float64
}

func (s *SavingAccount) Withdraw(value float64) string {

	canWithdraw := value > 0 && value <= s.balance

	if canWithdraw {
		s.balance -= value
		return "The withdraw of " + fmt.Sprintf("%.2f", value) + " was concluded"
	} else {
		return "The account don't have balance enouth. withdraw an value bellow " + fmt.Sprintf("%.2f", s.balance)
	}

}

func (s *SavingAccount) Deposit(value float64) (string, float64) {

	if value > 0 {
		s.balance += value
		return "You deposit ", value
	} else {
		return "Deposit a value bigger than ", 0
	}

}
