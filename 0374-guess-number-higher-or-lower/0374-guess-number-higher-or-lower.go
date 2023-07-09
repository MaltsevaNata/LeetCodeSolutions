/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int { // 10
    head := 0
    tail := n
   
    for {
        n = (tail + head) / 2
        switch guess(n) {
        case 0:
            return n
        case -1:
            tail = n - 1
            break
        case 1:
            head = n + 1
            break
        }
    }
}