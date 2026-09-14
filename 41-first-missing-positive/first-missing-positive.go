func firstMissingPositive(nums []int) int {
	min := len(nums) + 1

	numsAsMap := make(map[int]struct{})
	for _, num := range nums {
		numsAsMap[num] = struct{}{}
	}

	for i := 1; i < len(nums)+1; i++ {
		if _, ok := numsAsMap[i]; !ok {
			min = i
			break
		}
	}

	return min
}