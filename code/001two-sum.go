/*暴力解法, O(n2)*/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums) - 1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if target - nums[j] == nums[i] {
				return []int{i, j}
			}
		}
	}

	return nil
}

/*hash法, O(n)*/
func twoSum(nums []int, target int) []int {
	hashTbl := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		j, ok := hashTbl[target-nums[i]]
		if ok {
			return []int{i, j}
		} else {
			hashTbl[nums[i]] = i
		}
	}
	
	return nil
}
