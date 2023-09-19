package main

import "fmt"

type User struct {
	Name string
	ID string
	Age int
}

type VIPUser struct {
	User
	VIPLevel int
	Price int
}

func main() {
	user := User{ Name: "Tom", ID: "123", Age: 18 }
	vip := VIPUser{ User: user, VIPLevel: 1, Price: 100 }

	fmt.Println(user)
	fmt.Println(vip.Name)
}
