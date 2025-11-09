package main

import "fmt"

func main() {
	integers := make([]int, 0, 10)

	fmt.Println("Enter up to 10 integers (press Enter after every number or Ctrl+D) :")

	for len(integers) < 10 {
		var num int
		if _, err := fmt.Scan(&num); err != nil {
			break
		}
		integers = append(integers, num)
	}

	BubbleSort(integers)
	fmt.Println(integers)
}

func BubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func Swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}
