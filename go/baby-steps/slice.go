package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	s := make([]int, 0, 3)

	for {
		var input string
		fmt.Print("Enter integer or X: ")
		fmt.Scan(&input)

		if input == "X" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			continue
		}

		s = append(s, num)
		sort.Ints(s)
		fmt.Println(s)
	}
}
