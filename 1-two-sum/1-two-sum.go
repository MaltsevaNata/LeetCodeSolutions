func twoSum(nums []int, target int) []int {
	var found bool
	var idx1, idx2 int
	for i, i_num := range nums {
		for k, k_num := range nums[(i + 1):] {
			if (i_num + k_num) == target {
				found = true
				idx1, idx2 = i, (k + i + 1)
				break
			}
		}
		if found {
			break
		}
	}
	return []int{idx1, idx2}
}