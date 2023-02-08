func isSubsequence(s string, t string) bool {
	if len(s) < 1 {
		return true
	}
	sPointer := 0
	for i := range t {
		if s[sPointer] == t[i] {
			if sPointer >= len(s)-1 {
				return true
			} else {
				sPointer++
			}
		}
	}
	return false
}
