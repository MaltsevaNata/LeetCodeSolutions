func nextGreaterElement(nums1 []int, nums2 []int) []int {
    greater := make(map[int]int)
    for _, val := range nums2 {
        greater[val] = -1
    }
    
    stack := []int{}
    
    for i := len(nums2) - 1; i >= 0; i-- {
        last := len(stack)-1
        for len(stack) > 0 && nums2[stack[last]] < nums2[i] {
            stack = stack[:last]
            last--
        } 
        if len(stack) > 0 {
            greater[nums2[i]] = nums2[stack[last]]
        }
        stack = append(stack, i)
    }
    
    result := make([]int, len(nums1))
    for i, val := range nums1 {
        result[i] = greater[val]
    }
    return result
}