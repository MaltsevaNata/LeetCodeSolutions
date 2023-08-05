func searchInsert(nums []int, target int) int {
    return searchArr(nums, 0, len(nums)-1, target)
}

func searchArr(nums []int, startIdx, endIdx int, target int) int {
    mid := startIdx + (endIdx - startIdx) / 2
    if nums[mid] == target {
        return mid
    }
    if nums[mid] > target {
        if mid == 0 || nums[mid-1] < target {
            return mid
        }
        return searchArr(nums, startIdx, mid, target)
    }
    if nums[mid] < target {
        if mid == len(nums) - 1 || nums[mid+1] > target {
            return mid + 1
        }
        return searchArr(nums, mid+1, endIdx, target)
    }
    return -1
}