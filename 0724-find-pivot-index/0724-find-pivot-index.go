
func pivotIndex(nums []int) int {
	var (
		sumBefore   int
		mapForAfter = make(map[int]int)
		tempSum     int
	)
	for i := len(nums) - 1; i >= 0; i-- {
		mapForAfter[i] += tempSum
		tempSum += nums[i]
	}
	for i, num := range nums {
		if sumBefore == mapForAfter[i] {
			return i
		}
		sumBefore += num
	}
	return -1
}
