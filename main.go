package main

import "fmt"

func findSecLargest(arr []int) int {
	if len(arr) < 2 {
		return -1
	}

	largest, secLargest := arr[0], -1

	for i := 1; i < len(arr); i++ {
		if arr[i] > largest {
			secLargest = largest
			largest = arr[i]
		} else if arr[i] > secLargest && arr[i] != largest {
			secLargest = arr[i]
		}
	}

	if secLargest == -1 {
		return -1
	}

	return secLargest
}

func main() {
	arr := []int{20, 35, 40, 15, 30}
	secLargest := findSecLargest(arr)
	if secLargest == -1 {
		fmt.Println("No second largest element found")
	} else {
		fmt.Printf("The second largest element is: %d\n", secLargest)
	}
}
