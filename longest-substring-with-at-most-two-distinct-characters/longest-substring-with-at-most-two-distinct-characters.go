func lengthOfLongestSubstringTwoDistinct(s string) int {
    var left, right int
    chars := make(map[rune]struct{})
    max := 0
    
    for right < len(s) {
        _, ok := chars[rune(s[right])]
        if ok || (!ok && len(chars) < 2) {
            chars[rune(s[right])] = struct{}{}
            right++
        } else {
            length := right - left
            if length > max {
                max = length
            }
            left++
            right = left

            for k := range chars {
                delete(chars, k)
            }
        }
    }
    length := right - left
    if length > max {
        max = length
    }
    return max
}