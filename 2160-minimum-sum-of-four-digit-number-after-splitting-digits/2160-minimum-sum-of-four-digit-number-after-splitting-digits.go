func minimumSum(num int) int {
	sorted, _ := sortNumbers(strconv.Itoa(num))
	a, err := strconv.Atoi(fmt.Sprintf("%s%s", sorted[0], sorted[2]))
	if err != nil {
		return 0
	}
	b, err := strconv.Atoi(fmt.Sprintf("%s%s", sorted[1], sorted[3]))
	if err != nil {
		return 0
	}
	return a + b
}

func sortNumbers(in string) ([]string, error) {
	data := strings.Split(in, "")

	var lastErr error
	sort.Slice(data, func(i, j int) bool {
		a, err := strconv.ParseInt(data[i], 10, 64)
		if err != nil {
			lastErr = err
			return false
		}
		b, err := strconv.ParseInt(data[j], 10, 64)
		if err != nil {
			lastErr = err
			return false
		}
		return a < b
	})
	return data, lastErr
}
