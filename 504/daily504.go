package main

import "fmt"

// Keep last n orders.
// Create api to append order and to retrieve nth last order
// time and space optimized
func main() {
	o := makeLastOrders(2)
	o.record(7)
	o.record(9)
	o.record(13)
	fmt.Println(o)
	fmt.Println(o.getLast(1))
	fmt.Println(o.getLast(2))
	fmt.Println(o.getLast(3))
}

type LastOrders struct {
	orderIds []int
	next int
	total int
}

func makeLastOrders(n int) LastOrders {
	return LastOrders {
		orderIds: make([]int, n),
	}
}

func (this *LastOrders) record(orderId int) {
	this.orderIds[this.next] = orderId
	this.next = (this.next + 1) % len(this.orderIds)
	this.total++
}

func (this *LastOrders) getLast(i int) int {
	l := len(this.orderIds)
	if i > this.total {
		panic(fmt.Sprintf("Don't have that many orders: %d out of %d", i, this.total))
	}
	if i > l {
		panic(fmt.Sprintf("Don't have that big a buffer: %d out of %d", i, l))
	}
	return this.orderIds[(l + this.next - i) % l]
}