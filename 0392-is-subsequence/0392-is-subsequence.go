func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	curSubCharInd := 0
	for ind := 0; ind < len(t); ind++ {
		if curSubCharInd >= len(s) {
			return true
		}
		if (len(s) - curSubCharInd) > (len(t) - ind) {
			return false
		}
		if t[ind] == s[curSubCharInd] {
			curSubCharInd += 1
		}
	}
	if curSubCharInd == len(s) {
		return true
	}
	return false
}