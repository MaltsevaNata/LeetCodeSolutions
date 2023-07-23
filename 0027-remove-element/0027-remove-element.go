func removeElement(nums []int, val int) int {
    i := 0
    k := 0
    
    for i < len(nums){
        if nums[i] == val {
            last := len(nums) - 1
            nums[last], nums[i] = nums[i], nums[last]
            nums = nums[:last]
        } else {
            i++
            k++
        }
    }
    return k
}