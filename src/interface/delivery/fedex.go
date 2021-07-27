package delivery

import "fmt"

type FedexSender struct {}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex sends %v parcel\n", parcel)
}

func Start() {
	fmt.Println("Good Lluck!")
}