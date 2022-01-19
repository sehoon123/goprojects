package main

func main() {
	arr := []int{1, 2}
	arr = append(arr[:1], arr[2:]...)
}
