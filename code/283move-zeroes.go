/*遍历替换*/
func moveZeroes(nums []int)  {
    j := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[j] = nums[i]
            if i != j {
                nums[i] = 0
            }
            j++
        }
    }
}