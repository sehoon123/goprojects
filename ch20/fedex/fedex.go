package fedex

import "fmt"

type FedexSender struct {

}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("FedexSender: sending %s\n", parcel)
}