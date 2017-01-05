package main

import (
	"fmt"
)

func build_up(left, right []int) (result []int){
	left_index, right_index := 0, 0

	result = make([]int, len(left) + len(right))

	for i := 0; i < len(left) + len(right); i++ {
		if left_index >= len(left) {
			result[i] = right[right_index]
			right_index++
		} else if right_index >= len(right) {
			result[i] = left[left_index]
			left_index++
		} else if left[left_index] < right[right_index] {
			result[i] = left[left_index]
			left_index++
		} else {
			result[i] = right[right_index]
			right_index++
		}
	}
	return result
}

func mergesort(arr []int, ic chan []int) {
	if len(arr) <= 1 {
		ic <- arr
		return
	}

	leftc  := make(chan []int)
	rightc := make(chan []int)

	mid := len(arr) / 2

	go mergesort(arr[:mid], leftc)
	go mergesort(arr[mid:], rightc)

	left_arr  := <-leftc
	right_arr := <-rightc

	close(leftc)
	close(rightc)

	ic <- build_up(left_arr, right_arr)
	return
}

func main() {
	arr := []int{123,52,3,42,3536,235,73,243,53,64234,6,34524,63,1,1,4235,362,3}

	fmt.Println("Unsorted: ", arr)

	arrc := make(chan []int)

	go mergesort(arr, arrc)

	sorted := <-arrc
	close(arrc)
	fmt.Println("Sorted: ", sorted)
}
