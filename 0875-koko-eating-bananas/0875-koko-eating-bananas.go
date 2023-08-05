func minEatingSpeed(piles []int, h int) int {
    max := 0
    for _, pile := range piles {
        if pile > max {
            max = pile
        }
    }
    return findMinK(piles, 1, max, h)
}

func canEat(piles []int, h, k int) bool {
    total := 0
    for _, pile := range piles {
        total += pile / k
        if pile % k > 0 {
            total ++
        }
    }
    if total <= h {
        return true
    }
    return false
}

func findMinK(piles []int, start, end, h int) int {
    mid := start + (end - start) / 2
    if canEat(piles, h, mid) {
        if mid == 1 || !canEat(piles, h, mid-1) {
            return mid
        }
        return findMinK(piles, start, mid, h)
    }
    return findMinK(piles, mid+1, end, h)
}