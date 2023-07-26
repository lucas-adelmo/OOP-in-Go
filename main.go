package main

import (
	a "bank/account"
	h "bank/client"
	"fmt"
)

func PayBill(account bankAccount, value float64) string {
	return account.Withdraw(value)
}

type bankAccount interface {
	Withdraw(value float64) string //Any account which has the method Withdraw, sign the contract bankAccount.
}

func main() {

	howInstanciate()

	// get multiples returns from the function:
	accountX := a.CurrentAccount{Holder: h.Holder{Name: "Xman"}, Agency: 4635, Account: 468957}
	message, value := accountX.Deposit(1100)

	fmt.Println(message, value)
	fmt.Println(accountX)

	//Testing operations:
	accountAlpha := a.CurrentAccount{Holder: h.Holder{
		Name: "Alpha",
		Cpf:  "123.456.450-89",
	}}

	accountBeta := a.SavingAccount{Holder: h.Holder{
		Name: "Beta",
		Cpf:  "123.458.465-00",
	}}

	accountAlpha.Deposit(1000)
	accountBeta.Deposit(500)

	PayBill(&accountAlpha, 950)
	PayBill(&accountBeta, 50)
	// accountAlpha.Transfer(500, &accountBeta)

	// Status
	fmt.Println(accountAlpha)
	fmt.Println(accountBeta)

}

func howInstanciate() {
	//using a composite literal initialization
	var account1 a.CurrentAccount = a.CurrentAccount{Holder: h.Holder{
		Name: "Lucas",
		Cpf:  "123.456.45-55",
	}, Agency: 123, Account: 465798}

	// using the short variable declaration syntax (:=). The Go compiler will infer the type of account2 from the
	// right-hand side, which is a composite literal representing a a.CurrentAccount struct with the provided values.
	account2 := a.CurrentAccount{Holder: h.Holder{
		Name: "Lucas Santos",
		Cpf:  "123.456.45-55",
	}, Agency: 123, Account: 465798}

	fmt.Println(&account1 == &account2) //false -> different adress memory alocation
	fmt.Println(account1 == account2)   //true -> equal content

	//The return type of new(a.CurrentAccount) is *a.CurrentAccount, which means it returns a pointer to a a.CurrentAccount struct.
	//Therefore, the variable that will store this value needs to be of the pointer type, i.e., *a.CurrentAccount.
	var account3 *a.CurrentAccount = new(a.CurrentAccount)
	account3.Holder.Name = "Bolsonaro"

	var account4 *a.CurrentAccount = new(a.CurrentAccount)
	account4.Holder.Name = "Bolsonaro"

	fmt.Println(account4 == account3)   //false -> Because we set the type as a pointer (*a.CurrentAccount) we are comparing adress memory alocation
	fmt.Println(*account4 == *account3) //true -> When you have a pointer variable and want to access the value it points to, you use the * . In Go, the * operator is used as both a pointer type specifier and a pointer dereference operator.

}
