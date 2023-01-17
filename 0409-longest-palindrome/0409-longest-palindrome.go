func longestPalindrome(s string) (length int) {
	letterCounter := make(map[int32]int)
	for _, char := range s {
		_, ok := letterCounter[char]
		if !ok {
			letterCounter[char] = 1
		} else {
			letterCounter[char] += 1
		}
	}
	for _, count := range letterCounter {
		if count%2 == 0 {
			length += count
		} else {
			if count > 2 {
				length += count - 1
			}
		}
	}
	if length < len(s) {
		length += 1 // middle element can be without pair
	}
	return length
}