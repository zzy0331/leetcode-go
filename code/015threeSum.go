/*暴力求解, O(n3)*/
func threeSum(nums []int) [][]int {
	
	var answer [][]int

	for i := 0; i < len(nums) - 2; i++ {
		for j := i + 1; j < len(nums) - 1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] + nums[j] + nums[k] == 0 {
					temp := []int{nums[i], nums[j], nums[k]}
					answer = append(answer, temp)
				}
			}
		}
	}

	return answer
}

/*hash表来记录O(n2), a, b, a + b = -c
func threeSum(nums []int) [][]int {
	var answer [][]int
	
	hashTbl := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hashTbl[-nums[i]] = i
	}

	for j := 0; j < len(nums) - 1; j++ {
		for k := j + 1; k < len(nums); k++ {
			_, ok := hashTbl[nums[j]+nums[k]]
			if ok {
				temp := []int{nums[j], nums[k], -(nums[k]+nums[j])}
				answer = append(answer, temp)
			}
		}
	}

	return answer
}

/*先排序，再双指针向内夹逼办法, O(n2)*/
func threeSum(nums []int) [][]int {
	var answer [][]int

	/*从小到大排序*/
	sort.Ints(nums)
	
	for k := 0; k < len(nums) - 2; k++ {
		/*如果k所在的值比0大，说明k以后都比0大，加起来不可能为0*/
		if nums[k] > 0 {
			break
		}

		/*跳过重复的值以便计算到重复的集合*/
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i := k + 1
		j := len(nums) - 1
		for i < j {
			s := nums[i] + nums[j] + nums[k]
			if s > 0 {
				j--
			} else if s < 0 {
				i++
			} else {
				answer = append(answer, []int{nums[i], nums[j], nums[k]})
				i++
				j--

				for i < j && nums[i-1] == nums[i] {
					i++
				}

				for i < j && nums[j+1] == nums[j] {
					j--
				}
			}
		}
	}

	return answer
}
