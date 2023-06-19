func rotate(nums []int, k int)  {
    	c := make([]int, len(nums))
	copy(c, nums)
	length := len(nums)
	for i := 0; i < length; i++ {
		nums[(i+k)%length] = c[i]
	}
}