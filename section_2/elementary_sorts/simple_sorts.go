package simple_sorts

// import (
// 	"fmt"
// )

func insertion(a []int) {
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

func shellsort(a []int) {
	divisor := 2
	k := int(len(a) / divisor)
	
	for k > 0 {
		for j := 0; j + k < len(a); j++ {
			i := int(0)

			for a[j - i] > a[j + k - i] &&  j - i >= 0 {
				// swap the two things
				a[j - i] = a[j - i] ^ a[j + k - i]
				a[j + k - i] = a[j - i] ^ a[j + k - i]
				a[j - i] = a[j - i] ^ a[j + k - i]

				// add delta k to increment
				i += k
	 		}
		}
 		divisor++
 		k = len(a) / divisor
	}
}

func bubblesort(a []int) {
	swap := true
	for swap == true {
		swap = false
		for i := 0; i < len(a) - 1; i++ {
			if a[i] > a[i + 1] {
				a[i] = a[i] ^ a[i + 1]
				a[i + 1] = a[i] ^ a[i + 1]
				a[i] = a[i] ^ a[i + 1]

				swap = true
			}
		}
	}
	
}

// func main() {
// 	fmt.Println("insertion")
// 	arr := []int{1, 3, 2, 14, 6, 1512, 3, 16, 69}
// 	fmt.Println("Unsorted: ", arr)
// 	insertion(arr)
// 	fmt.Println("Sorted: ", arr)

// 	fmt.Println("shellsort")
// 	arr = []int{1, 3, 2, 14, 6, 1512, 3, 16, 69}
// 	fmt.Println("Unsorted: ", arr)
// 	shellsort(arr)
// 	fmt.Println("Sorted: ", arr)

// 	fmt.Println("bubblesort")
// 	arr = []int{1, 3, 2, 14, 6, 1512, 3, 16, 69}
// 	fmt.Println("Unsorted: ", arr)
// 	shellsort(arr)
// 	fmt.Println("Sorted: ", arr)
// }