package accounts

import (
	"errors"
	"fmt"
)

//We need to write      go mod init     in our directory at the first time to create .mod file
//Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

//NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account //Returning a new a new object with same value
	// If we didn't use pointer, It will try to copy every time, so program would be slow
}

//Deposit x amount on your account
//a is a receiver
// For doing this, our Account object would have Deposit method
//This method is for editing private property
//This method means using an Account object who calls Deposit(), not copying Account
//If you want to change actual object's value, put *
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//Balance of your account
func (a Account) Balance() int {
	return a.balance
}

//Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	//erorr type has two value, error and nil
	return nil
}

//ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\n Has: ", a.Balance())
}
