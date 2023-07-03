func removeDuplicates(nums []int) int {
	returnedArrayPointer := 0
	originalNums := make([]int, 0, len(nums))
	originalNums = nums[:]
	bufferedNums := map[int]bool{}

	for i := 0; i < len(originalNums); i++ {
		_, ok := bufferedNums[originalNums[i]]
		if ok {
			continue
		}
		nums[returnedArrayPointer] = originalNums[i]
		bufferedNums[originalNums[i]] = true
		returnedArrayPointer++
	}
	fmt.Println(nums)
	return returnedArrayPointer
}
