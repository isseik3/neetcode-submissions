func twoSum(nums []int, target int) []int {
    x := make(map[int]int, len(nums))
    for i := 0; i < len(nums); i++ {
        if first, ok := x[nums[i]]; ok {
            return []int{first, i} 
        }
        x[target-nums[i]] = i
    }
    return []int{}
}
