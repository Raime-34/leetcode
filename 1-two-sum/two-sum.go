func twoSum(nums []int, target int) []int {
	numsAsMap := make(map[int]int, len(nums))

	for i, num := range nums {
		numsAsMap[num] = i
	}

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if index, ok := numsAsMap[diff]; ok && i != index {
			return []int{i, index}
		}
	}

	return []int{}
}