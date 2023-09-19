package koreaPost

import "fmt"

type PostSender struct {

}

func (k *PostSender) Send(msg string) {
	fmt.Printf("[국제포스트] %s\n", msg)
}