func isSubsequence(s string, t string) bool {
	curSubCharInd := 0
    lenS := len(s)
    lenT := len(t)
	for ind := 0; ind < lenT; ind++ {
		if curSubCharInd >= lenS {
			return true
		}
		if (lenS - curSubCharInd) > (lenT - ind) {
			return false
		}
		if t[ind] == s[curSubCharInd] {
			curSubCharInd += 1
		}
	}
	if curSubCharInd == lenS {
		return true
	}
	return false
}