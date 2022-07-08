package code

import (
	"regexp"
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	r, _ := regexp.Compile(`^[+-]?\d+`)
	res := r.FindString(s)
	if res != "" {
		n, _ := strconv.Atoi(res)
		if n > 2147483647 {
			return 2147483647
		}
		if n < -2147483648 {
			return -2147483648
		}
		return n
	}
	return 0
}
