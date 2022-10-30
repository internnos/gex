package main

import (
	"fmt"
	"testing"
)

func TestLimit(t *testing.T) {
	l := NewLimit(10000)
	buyOrder := NewOrder(true, 5)
	l.AddOrder(buyOrder)
	fmt.Println(l)
}
