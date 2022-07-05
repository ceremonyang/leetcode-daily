package main

import "fmt"

func maxArea(height []int) int {
	s := 0
	e := len(height)
	max := 0
	sum := 0
	for {
		if height[s] < height[e-1] {
			sum = height[s] * (e - s - 1)
			s++
		} else {
			sum = height[e-1] * (e - s - 1)
			e--
		}
		if max < sum {
			max = sum
		}
		if e <= s {
			break
		}
	}
	return max
}

func main() {
	height := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(maxArea(height))
}
