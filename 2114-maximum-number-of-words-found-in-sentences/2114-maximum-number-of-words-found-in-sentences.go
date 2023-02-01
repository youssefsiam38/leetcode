func mostWordsFound(sentences []string) int {
	max := 0
	for _, sentence := range sentences {
		m := strings.Count(sentence, " ")
		if m > max {
			max = m
		}
	}
	
	return max + 1
}