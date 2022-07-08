package code

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	arr := [][]string{}
	res := ""

	x, y, count := 0, 0, 0
	for count < len(s) {
		arr1 := make([]string, numRows)
		arr2 := make([]string, numRows)
		for y < numRows && x%(numRows-1) == 0 && count < len(s) {
			arr1[y] = string(s[count])
			count++
			y++
		}
		y--
		arr = append(arr, arr1)
		for y > 0 && x >= 0 && count < len(s) {
			y--
			x++
			arr2[y] = string(s[count])
			count++
		}
		y++
		arr = append(arr, arr2)
	}
	i := 0
	for i < numRows {
		for _, v := range arr {
			res += fmt.Sprint(v[i])
		}
		i++
	}
	return res
}
