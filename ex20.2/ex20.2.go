package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
}

func (s *Student) String() string {
	return s.Name
}


type User struct {
	Name string
	Age int
}

func (u *User) String() string {
	return u.Name
}

func toString(s Stringer) string {
	return s.String()
}

func main() {
	var sehoon Stringer = &Student{Name: "Sehoon"}
	var user Stringer = &User{Name: "hoon", Age: 30}

	fmt.Println(toString(sehoon))
	fmt.Println(toString(user))


}