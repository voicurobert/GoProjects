package main

import "fmt"

// Order struct
type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

// Queue array of Order
type Queue []*Order

func (o *Order) new(priority, quantity int, product, customerName string) {
	o.priority = priority
	o.quantity = quantity
	o.product = product
	o.customerName = customerName
}

func (q *Queue) add(order *Order) {
	if len(*q) == 0 {
		*q = append(*q, order)
	} else {
		appended := false
		for i, addedOrder := range *q {
			if order.priority > addedOrder.priority {
				*q = append((*q)[:i], append(Queue{order}, (*q)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*q = append(*q, order)
		}

	}
}

func main() {
	var queue Queue = make(Queue, 0)
	o1 := &Order{
		priority:     2,
		quantity:     20,
		product:      "Computer",
		customerName: "Bla",
	}
	o2 := &Order{}
	o2.new(1, 10, "Monitor", "XXX")

	queue.add(o1)
	queue.add(o2)
	fmt.Println(queue[0])
}
