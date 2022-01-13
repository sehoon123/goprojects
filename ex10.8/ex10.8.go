package main

import "fmt"

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "red"
	case Blue:
		return "blue"
	case Green:
		return "green"
	case Yellow:
		return "yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Red
}

func main() {
	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))
}