func peakIndexInMountainArray(arr []int) int {
	for idx, val := range arr {
        if val > arr[idx+1] {
            return idx
        }
	}
    return 0
}