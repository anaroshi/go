package main

import (
	"interface/delivery"
	"interface/quick"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	delivery.Start()
	koreaPostsender := &delivery.PostSender{}
	SendBook("어린 왕자", koreaPostsender)
	SendBook("그리스인 조르바", koreaPostsender)

	fedexSender := &delivery.FedexSender{}
	SendBook("어린 왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)

	quickSender := &quick.QuickSender{}
	SendBook("어린 왕자", quickSender)
	SendBook("그리스인 조르바", quickSender)

}