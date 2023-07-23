func removeElement(nums []int, val int) int {
    if len(nums) == 0 {
        return 0
    }
    i := 0
    length := len(nums)
    for i < length {
        if nums[i] == val {
            nums[length-1], nums[i] = nums[i], nums[length-1]
            length-- 
        } else {
            i++
        }
    }
    return length
}