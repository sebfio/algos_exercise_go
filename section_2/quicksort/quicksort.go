package main

import (
	"fmt"
)

func swap(arr []int, a,b int) {
	arr[a] = arr[a] ^ arr[b]
	arr[b] = arr[a] ^ arr[b]
	arr[a] = arr[a] ^ arr[b]
}

func middle3(arr []int, a, c int) int{
	b := (a + c) / 2
	
	if arr[a] > arr[b] {
		if arr[b] > arr[c] {
			return arr[b]
		} else if arr[c] > arr[a] {
			return arr[a]
		} else {
			return arr[c]
		}
	} else { // arr[b] >= arr[a]
		if arr[a] > arr[c] {
			return arr[a]
		} else if arr[c] > arr[b] {
			return arr[b]
		} else {
			return arr[c]
		}
	}
}

func partition(arr []int, low, high int) (pivot_index int){
	pivot := arr[0]
	if high - low > 1 {
		pivot = middle3(arr, low, high)
	}
	for true {
		for arr[low] < pivot && low < high {
			low++
		}
		for arr[high] > pivot && high > low {
			high--
		}
		if low == high {
			break
		}
		swap(arr, low, high)
	}
	return high
}

func quicksort(arr []int, low, high int) {
	if high - low < 1 {
		return
	}
	index := partition(arr, low, high)
	quicksort(arr, low, index - 1)
	quicksort(arr, index + 1, high)
}

func main () {
	arr := []int{123,513,54523,62,63,72,1,36,4,37,4,74,2,342316,27,3258}
	fmt.Println("Arr: ", arr, "len: ", len(arr))
	quicksort(arr, 0, len(arr) - 1)
	fmt.Println("Sort: ", arr, "len: ", len(arr))
}
