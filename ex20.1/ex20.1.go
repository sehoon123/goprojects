package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age int
}

func (s Student) String() string {
	return fmt.Sprintf("%s: %d", s.Name, s.Age)
}

func (s Student) GetAge() int {
	return s.Age
}

func main() {
	student := Student{"Tom", 20}
	var stringer Stringer

	stringer = student
	fmt.Printf("%s\n", stringer.String()) // ,stringer.GetAge())
	fmt.Printf("%s\n", student.String())

}