func isIsomorphic(s string, t string) bool {
	m := map[string]string{}
	m2 := map[string]string{}
	for i, ch := range s {
		if v, ok := m[getChar(ch)]; ok {
			if v != getChar(int32(t[i])) {
				return false
			}
		}
		if v, ok := m2[getChar(int32(t[i]))]; ok {
			if v != getChar(ch) {
				return false
			}
		}
		m[getChar(ch)] = getChar(int32(t[i]))
		m2[getChar(int32(t[i]))] = getChar(ch)

	}
	return true
}

func getChar(n int32) string {
	return fmt.Sprintf("%c", n)
}
