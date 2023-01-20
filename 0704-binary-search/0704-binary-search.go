func findInSubArr(arr []int, startIndex int, target int) int {
	length := len(arr)
	switch length {
	case 0:
		return -1
	case 1:
		if arr[0] == target {
			return startIndex
		}
		return -1
	default:
		middle := length / 2
		if target < arr[middle] {
			return findInSubArr(arr[0:middle], startIndex, target)
		} else {
			return findInSubArr(arr[middle:length], startIndex+middle, target)
		}
	}
}

func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 || target < nums[0] || target > nums[length-1] {
		return -1
	}
	return findInSubArr(nums, 0, target)
}