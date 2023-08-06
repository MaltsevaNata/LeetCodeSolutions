func nextGreaterElements(nums []int) []int {
    result := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        result[i] = -1
    }
    
    stack := []int{}
    for j := 0; j < 2; j++ {
        for i := 0; i < len(nums); i++ {
            last := len(stack) - 1
            for len(stack) > 0 && nums[stack[last]] < nums[i]   {
                item := stack[last]
                stack = stack[:last]
                result[item] = nums[i]
                last--
            }
            stack = append(stack, i)
        }
    }
    return result
}