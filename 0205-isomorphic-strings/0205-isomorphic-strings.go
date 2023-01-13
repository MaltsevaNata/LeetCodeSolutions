func getLettersCode(char int32, ind int, letterCodeMap map[int32]int) int {
	letterCode, ok := letterCodeMap[char]
	if !ok {
		letterCodeMap[char] = ind
		return ind
	}
	return letterCode
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if s == t {
		return true
	}

	letterCodeMapS := make(map[int32]int) // {letter: code}
	letterCodeMapT := make(map[int32]int) // {letter: code}

	for ind, char := range s {
		if getLettersCode(char, ind, letterCodeMapS) != getLettersCode(int32(t[ind]), ind, letterCodeMapT) {
			return false
		}
	}
	return true
}
