package code

import "fmt"

func isHappy(n int) bool {
	circle := map[int]int{}
	for {
		n = getSum(n)
		fmt.Println(n)
		if n == 1 {
			return true
		}
		if _, ok := circle[n]; ok {
			return false
		}
		circle[n] = 1
	}
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
