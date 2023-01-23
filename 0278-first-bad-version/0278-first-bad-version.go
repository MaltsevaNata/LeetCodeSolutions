/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func checkBadVersion(startN int, endN int) int {
	if !isBadVersion(endN) {
		return -1
	}
	if isBadVersion(startN) {
		return startN
	}
	if endN-startN == 1 && isBadVersion(endN) {
		return endN
	}
	middle := (endN + startN) / 2
	left := checkBadVersion(startN, middle)
	if left < 1 {
		return checkBadVersion(middle, endN)
	}
	return left
}

func firstBadVersion(n int) int {
	return checkBadVersion(1, n)
}