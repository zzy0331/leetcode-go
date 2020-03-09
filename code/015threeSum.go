/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (25.54%)
 * Likes:    1838
 * Dislikes: 0
 * Total Accepted:    161K
 * Total Submissions: 628.4K
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？找出所有满足条件且不重复的三元组。
 * 
 * 注意：答案中不可以包含重复的三元组。
 * 
 * 
 * 
 * 示例：
 * 
 * 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 * 
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2] 
 * ]
 * 
 * 
 */
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

/*hash表来记录, a, b, a + b = -c, O(n2)*/
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
