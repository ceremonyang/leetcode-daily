package code

func isIsomorphic(s string, t string) bool {
	m := make(map[byte]byte)
	n := make(map[byte]byte)
	for i := range s {
		if v, ok := m[s[i]]; ok {
			if v != t[i] {
				return false
			}
		} else {
			m[s[i]] = t[i]
			n[t[i]] = s[i]
		}
	}
	return len(m) == len(n)
}
