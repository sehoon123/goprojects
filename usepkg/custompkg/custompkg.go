package custompkg

import "fmt"

type Student struct {
	Name string
	Age int
	Score int
}

var Ratio int
var ttt int

const PI = 3.14
const pI2 = 3.1415926

func PrintCustom() {
	fmt.Println("This is custom package!")
}

func printcustom2() {
	fmt.Println("This is custom package2!")

}