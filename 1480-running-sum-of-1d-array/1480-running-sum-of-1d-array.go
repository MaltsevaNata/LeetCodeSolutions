func runningSum(nums []int) []int {
	for ind := 1; ind < len(nums); ind++ {
		nums[ind] += nums[ind-1]
	}
	return nums
}