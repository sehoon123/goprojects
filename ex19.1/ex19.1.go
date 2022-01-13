package main

import "fmt"

type account struct {
	balance   int
	firstname string
	lastname  string
}

func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

func main() {
	var mainA = &account{balance: 1000, firstname: "John", lastname: "Doe"}
	mainA.withdrawPointer(500)
	fmt.Println(mainA.balance)

	mainA.withdrawValue(500)
	fmt.Println(mainA.balance)

}
