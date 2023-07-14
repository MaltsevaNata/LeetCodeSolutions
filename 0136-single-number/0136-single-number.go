func singleNumber(nums []int) int {
    total := 0
    for _, num := range nums {
        total = total ^ num
    }
    return total
}