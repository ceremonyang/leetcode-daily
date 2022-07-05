package main

import "fmt"

func maxArea(height []int) int {
	s := 0
	e := len(height)
	max := 0
	sum := 0
	for e < s {
		if height[s] < height[e-1] {
			sum = height[s] * (e - s)
			s++
		} else {
			sum = height[e-1] * (e - s)
			e--
		}
		if max < sum {
			max = sum
		}
	}
	return max
}

func main() {
	height := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(maxArea(height))
}
