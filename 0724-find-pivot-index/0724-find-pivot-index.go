func pivotIndex(nums []int) int {
	leftSum := 0
	rightSum := 0
	for ind := 0; ind < len(nums); ind++ {
		rightSum += nums[ind]
	}
	for ind := 0; ind < len(nums); ind++ {
		rightSum -= nums[ind]
		if leftSum == rightSum {
			return ind
		}
		leftSum += nums[ind]
	}
	return -1
}