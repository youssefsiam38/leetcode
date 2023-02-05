func smallestEvenMultiple(n int) int {
	ans := n
	for {
		if ans%2 == 0 && ans%n == 0 {
			return ans
		}
		ans++
	}
}
