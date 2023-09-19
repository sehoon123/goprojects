package main

import (
	"goprojects/ch20/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	var sender Sender = &koreaPost.PostSender{}
	SendBook("The Go Programming Language", sender)
	SendBook("Harry Potter", sender)
}