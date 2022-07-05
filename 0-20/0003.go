
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	start := 0
	diff := 0
	for i := 1; i < len(s); i++ {
		tmp := start
		for i > tmp {
			if s[i] == s[tmp] {
				start = tmp + 1
				break
			}
			tmp++
		}
		if diff < i-start {
			diff = i - start
		}
	}
	return diff + 1
}

