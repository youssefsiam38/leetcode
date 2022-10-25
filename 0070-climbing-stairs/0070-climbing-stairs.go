func climbStairs(n int) int {
	mem := map[int]int{}
	return step(n, &mem)
}

func step(current int, mem *map[int]int) (r int) {
	if nn, ok := (*mem)[current]; ok {
		return nn
	}
	if current == 1 {
		return 1
	}
	if current == 2 {
		return 2
	}
	(*mem)[current] = step(current-1, mem) + step(current-2, mem)
	return (*mem)[current]
}
