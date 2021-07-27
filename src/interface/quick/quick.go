package quick

import "fmt"

type QuickSender struct {}

func(q *QuickSender) Send(parcel string ) {
	fmt.Printf("Quick Service sends %v parcel\n", parcel)
}