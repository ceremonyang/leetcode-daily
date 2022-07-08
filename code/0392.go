package code

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
		if i == len(s) || j == len(t) {
			break
		}
	}
	if i == len(s) && j <= len(t) {
		return true
	}
	return false
}
