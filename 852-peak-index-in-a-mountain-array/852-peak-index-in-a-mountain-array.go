func peakIndexInMountainArray(arr []int) int {
	max_idx := len(arr) - 1
	var res_idx int
	for idx, val := range arr {
		if idx < max_idx {
			if val > arr[idx+1] {
				res_idx = idx
				break
			}
		} else {
			res_idx = idx
			break
		}
	}
	return res_idx
}