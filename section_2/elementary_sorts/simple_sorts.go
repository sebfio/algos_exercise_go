package main


import (
	"fmt"
)

func (a []int) insertion() {
	length := len(a)
	
	for i := 1; i < length; i++ {
		for j := i; j < length; j++ {
			if a[i - 1] > a[j] {
				temp := a[i - 1]
				a[i - 1] = a[j]
				a[j] = temp
			}
		}
	}
}

func (a []int) shellsort() {
	
}
