package main

import "sync"

//make chatbot
//use goroutine
//use channel
func chatbot(in chan string, out chan string) {
	var wg sync.WaitGroup
	for {
		wg.Add(1)
		msg := <-in
		if msg == "exit" {
			break
		}
		out <- msg
		wg.Done()
	}

}
func main() {
	var wg sync.WaitGroup

	in := make(chan string)
	out := make(chan string)
	go chatbot(in, out)
	for {
		msg := <-out
		if msg == "exit" {
			break
		}
		in <- msg
	}
	wg.Wait()
}
