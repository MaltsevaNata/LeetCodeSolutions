func longestOnes(nums []int, k int) int {
    var left, right int
    max := 0
    kCopy := k
    
    for right < len(nums) {
        if nums[left] == 0 {
            left++
            right++
            continue
        }
        if nums[right] == 1 {
            right++
            continue
        } 
        if k > 0 {
            k--
            right++
            continue
        } 
        
        diff := right - left
        if diff > max {
            max = diff
        }
        k = kCopy
        left++
        right = left
    }
    for k > 0 && left >= 0 {
        left--
        k--
    }
    if left < 0 {
        left = 0
    }
    diff := right - left
    if diff > max {
        max = diff
    }
    return max
}