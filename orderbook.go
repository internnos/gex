package main

import "time"

type Order struct {
	Size      float64
	Bids      bool
	Limit     *Limit
	Timestamp int64
}

func NewOrder(bid bool, size float64) *Order {
	return &Order{
		Size:      size,
		Bids:      bid,
		Timestamp: time.Now().UnixNano(),
	}
}

type Limit struct {
	Price       float64
	Orders      []*Order
	TotalVolume float64
}

func NewLimit(price float64) *Limit {
	return &Limit{
		Price:  price,
		Orders: []*Order{},
	}
}

func (l *Limit) AddOrder(o *Order) {
	o.Limit = l
	l.Orders = append(l.Orders, o)
	l.TotalVolume += o.Size
}

type Orderbook struct {
	Asks []*Limit
	Bids []*Limit
}
