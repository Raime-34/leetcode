func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		slices.Sort(stones)
		iX := len(stones) - 1
		iY := iX - 1

		x := stones[iX]
		y := stones[iY]

		stones = stones[:iY]

		if res := x - y; res > 0 {
			stones = append(stones, res)
		}
	}

	if len(stones) == 0 {
		return 0
	}

	return stones[0]
}