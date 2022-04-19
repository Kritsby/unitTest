package main

import "fmt"

func main() {
	fmt.Println(max([]int{1, 2, 20, 30, 55, 4, 5}))
}

func max(numbers []int) int {
	var max int

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}

func maxIndex(numbers []int) int {
	var max int
	var index int

	for i, number := range numbers {
		if number > max {
			max = number
			index = i
		}
	}

	return index
}
