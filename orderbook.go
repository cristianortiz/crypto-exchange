package main

import "time"

type Order struct {
	Size float64
	Bid  bool // to set if the order is bid or ask type, maybe later it's make sense to create a specific bid and ask order type

	Limit     *Limit
	Timestamp int64 //unix nano
}

// group of orders at a certain price level
type Limit struct {
	Price       float64
	Orders      []*Order
	TotalVolume float64
}
type Orderbook struct {
	Asks []*Limit
	Bids []*Limit
}

func NewLimit(price float64) *Limit {
	return &Limit{
		Price:  price,
		Orders: []*Order{},
	}
}

func NewOrder(bid bool, size float64) *Order {
	return &Order{
		Size:      size,
		Bid:       bid,
		Timestamp: time.Now().UnixNano(),
	}
}
func (l *Limit) AddOrder(o *Order) {
	o.Limit = l
	l.Orders = append(l.Orders, o)
	l.TotalVolume += o.Size
}
