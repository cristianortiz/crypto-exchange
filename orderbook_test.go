package main

import (
	"fmt"
	"testing"
)

func TestLimit(t *testing.T) {
	l := NewLimit(10000)
	buyOrderA := NewOrder(true, 5)
	buyOrderB := NewOrder(true, 8)
	buyOrderC := NewOrder(true, 10)

	l.AddOrder(buyOrderA)
	l.AddOrder(buyOrderB)
	l.AddOrder(buyOrderC)

	fmt.Println(l)
	fmt.Println("deleting Order: ", buyOrderB.String())
	l.DeleteOrder(buyOrderB)
	fmt.Println(l)
}
func TestOrderbook(t *testing.T) {
	ob := NewOrderBook()
	buyOrder := NewOrder(true, 10)

	ob.PlaceOrder(18000, buyOrder)
	fmt.Printf("%+v", ob.Bids[0])

}
