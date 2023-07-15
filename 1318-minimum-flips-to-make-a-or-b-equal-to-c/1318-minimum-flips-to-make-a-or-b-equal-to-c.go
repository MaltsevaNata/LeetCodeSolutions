func countOnes(num int) int {
    var count int
    for num != 0 {
        count += num & 1
        num = num >> 1
    }
    return count
}

func minFlips(a int, b int, c int) int {
    or := a | b
    if or == c {
        return 0
    }
    diff := or ^ c
    
    addBits := diff & a & b
    return countOnes(diff) + countOnes(addBits)
}