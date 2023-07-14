func countBits(n int) []int {
    ans := make([]int, n+1)
    for i:=0; i<=n; i++ {
        count := 0
        iCopy := i
        for iCopy != 0 {
            count += iCopy & 1
            iCopy = iCopy >> 1
        }
        ans[i] = count
    }
    return ans
}