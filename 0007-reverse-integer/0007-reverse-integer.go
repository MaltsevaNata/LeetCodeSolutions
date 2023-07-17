func reverse(x int) int {
    var min = int(math.Pow(-2.0, 31))
    var max = int(math.Pow(2.0, 31))
    
    nums := []int{}
    var res int
    
    for math.Abs(float64(x)) > 0 {
        nums = append(nums, x%10)
        x /= 10
    }
    for i := 0; i < len(nums); i++ {
        res = res * 10 + nums[i] 
    }
    if res >= max || res < min {
        return 0
    }
    if x < 0 {
        res *= -1
    }
    
    return res
}