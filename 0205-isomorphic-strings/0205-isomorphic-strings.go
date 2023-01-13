func getLettersCode(str string) (code []int) {
	letterCodeMap := make(map[string]int) // {letter: code}
	for ind, char := range str {
		letter := string(char)
		letterCode, ok := letterCodeMap[letter]
		if !ok {
			letterCodeMap[letter] = ind
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