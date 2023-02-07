func runningSum(nums []int) []int {
	var (
		sum int
		ans []int
	)
	for _, num := range nums {
		sum += num
		ans = append(ans, sum)
	}
	return ans
}
