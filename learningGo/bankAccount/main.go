package main

import (
	"fmt"

	"github.com/joshrhee/bankAccount/accounts"
)

func main() {
	account := accounts.NewAccount("Josh")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	//We need to check error manually since there is no exception lol
	if err != nil {
		//We can just fmt.println(err) and it will not complete the program
		fmt.Println("Println: ", err)

		// //Calls the println and kills the program
		// log.Fatalln("Fatalln: ", err)
	}
	//In Go, if I try to print an object, it will try to call String()
	fmt.Println(account)
}
