package main

//#include "general_functions.h"
//int max(int a, int b);
//int min(int a, int b);
//void swap(int arr[], int a, int b);
import "C"

import (
	"fmt"
	"math/rand"
)

type priority_queue struct {
	arr []int
	size int
}

type PriorityQueue interface {
	init()
	insert()
	remove()
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func swap(arr []int, index_a, index_b int) {
	arr[index_a] = arr[index_a] ^ arr[index_b]
	arr[index_b] = arr[index_a] ^ arr[index_b]
	arr[index_a] = arr[index_a] ^ arr[index_b]
}

func (pq *priority_queue) swim(i int) {
	for i > 1 {
		if pq.arr[i/2] < pq.arr[i] {
			break
		}
		swap(pq.arr, i, i/2)
		i /= 2
	}
}

func (pq *priority_queue) sink(i int) {
	// get max child
	for i*2 <= pq.size {
		child := i*2
		if child + 1 <= pq.size && pq.arr[child] > pq.arr[child + 1] {
			child = child + 1
		}

		if pq.arr[child] < pq.arr[i] {
			swap(pq.arr, child, i)
		} else {
			break
		}
		i = child
	}
}

func (pq *priority_queue) insert(n int) {
	pq.size++
	pq.arr[pq.size] = n
	pq.swim(pq.size)
}

func (pq *priority_queue) remove(i int) (n int) {
	n = pq.arr[i]
	if i < pq.size {
		swap(pq.arr, i, pq.size);
		pq.size--
		pq.sink(i)
	}
	return n
}

func (pq *priority_queue) init() {
	pq.size = 0
	pq.arr = make([]int, 100)
}

func main() {
	minq := priority_queue{}

	minq.init()

	for i := 0; i < 12; i++ {
		minq.insert(rand.Intn(1000))
	}
	for i := 0; i < 12; i++ {
		fmt.Println(minq.remove(1))
	}
}
