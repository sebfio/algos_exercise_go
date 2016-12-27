//package main
package bag

import (
//	"fmt"
)

type bag interface {
	Add_item()
	Iterate() []int
	Is_empty() bool
	Get_size() uint32
}

type Node struct {
	data int
	next *Node
}

type Bag struct {
	head *Node
	size int
}

func (b *Bag) Add_item(val int) {
	b.size++

	n := new(Node)
	n.next = b.head
	n.data = val
	b.head = n
}

func (b *Bag) Iterate() []int {
	n := make([]int, b.size)
	curr := b.head

	for i := 0; i < b.size; i++ {
		n[i] = curr.data
		curr = curr.next
	}

	return n
}

func (b Bag) Get_size() int {
	return b.size
}

func (b Bag) Is_empty() bool {
	if b.size == 0 {
		return true
	} else {
		return false
	}
}

func Init() *Bag {
	b := new(Bag)
	b.size = 0
	b.head = nil

	return b
}

// func main() {
// 	b := Init()

// 	fmt.Println(b.size)

// 	b.Add_item(5)
// 	b.Add_item(6)
// 	b.Add_item(10)
// 	fmt.Println(b.Get_size())
// 	fmt.Println(b.Iterate())
// }
