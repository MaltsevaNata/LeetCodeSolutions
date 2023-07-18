func romanToInt(roman string) int { // XCIV
	dict := map[rune]int{
		'I': 1,
		'V': 5, 
        'X': 10,
        'L': 50, 
        'C': 100, 
        'D': 500, 
        'M': 1000,
        }
    var prev int // 94
    var res int // 10
    romanRune := []rune(roman)
    for i := len(romanRune) - 1; i >= 0; i-- {
        num := dict[romanRune[i]] // 10
        if prev != 0 && prev > num {
        res -= num
    } else {
        res += num 
    }
    prev = num
    }
    return res
}
