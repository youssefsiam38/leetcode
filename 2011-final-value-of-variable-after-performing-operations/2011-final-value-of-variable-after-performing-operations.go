func finalValueAfterOperations(operations []string) int {
	result := 0

	for _, operation := range operations {
		if strings.Contains(operation, "++") {
			result++
		} else if strings.Contains(operation, "--") {
			result--
		}
	}

	return result
}