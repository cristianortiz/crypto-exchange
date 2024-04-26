package main

import (
	"time"
)

type Order struct {
	Size      float64
	Bid       bool // to set if the order is bid or ask type, maybe later it's make sense to create a specific bid and ask order type
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
	Asks       []*Limit
	Bids       []*Limit
	AsksLimits map[float64]*Limit
	BidsLimits map[float64]*Limit
}
type Match struct {
	Ask        *Order
	Bid        *Order
	SizeFilled float64
	Price      float64
}

func NewOrderBook() *Orderbook {
	return &Orderbook{
		Asks:       []*Limit{},
		Bids:       []*Limit{},
		AsksLimits: make(map[float64]*Limit),
		BidsLimits: make(map[float64]*Limit),
	}
}

func (ob *Orderbook) PlaceOrder(price float64, o *Order) []Match {
	//1. try to match the orders, implementing matching logic to determine if the size of the order was fullfilment

	//2. add the rest of the order size (if it was not completed) to the books
	if o.Size > 0.0 {
		ob.add(price, o)

	}
	return []Match{}
}
func (ob *Orderbook) add(price float64, o *Order) {

	var limit *Limit
	if o.Bid {
		limit = ob.BidsLimits[price]
	} else {
		limit = ob.AsksLimits[price]
	}
	if limit == nil {
		limit = NewLimit(price)
		limit.AddOrder(o)

		if o.Bid {
			ob.Bids = append(ob.Bids, limit)
			ob.BidsLimits[price] = limit
		} else {
			ob.Asks = append(ob.Asks, limit)
			ob.AsksLimits[price] = limit
		}
	}

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

// func (o *Order) String() string {
// 	return fmt.Sprintf("[size]: %.2f", o.Size)
// }

//	func (l *Limit) String() string {
//		return fmt.Sprintf("[price: %.2f | volume: %.2f]", l.Price, l.TotalVolume)
//	}
func (l *Limit) AddOrder(o *Order) {
	o.Limit = l
	l.Orders = append(l.Orders, o)
	l.TotalVolume += o.Size
}

func (l *Limit) DeleteOrder(o *Order) {
	for i := 0; i < len(l.Orders); i++ {
		if l.Orders[i] == o {
			l.Orders[i] = l.Orders[len(l.Orders)-1]
			l.Orders = l.Orders[:len(l.Orders)-1]
		}
	}
	//for garbage collector, avoid nil pointers
	o.Limit = nil
	l.TotalVolume -= o.Size

	//TODO: resort the whole remaining orders
}
