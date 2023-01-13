func getLettersCode(str string) (code []int) {
	letterCodeMap := make(map[int32]int) // {letter: code}
	for ind, char := range str {
		letterCode, ok := letterCodeMap[char]
		if !ok {
			letterCodeMap[char] = ind
			letterCode = ind
		}
		code = append(code, letterCode)
	}

	return code
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if s == t {
		return true
	}
	sCode := getLettersCode(s)
	tCode := getLettersCode(t)
	if len(sCode) != len(tCode) {
		return false
	}
	for ind := 0; ind < len(sCode); ind++ {
		if sCode[ind] != tCode[ind] {
			return false
		}
	}
	return true
}