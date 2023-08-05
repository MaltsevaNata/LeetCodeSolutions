/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 *
 * func (this *ArrayReader) get(index int) int {}
 */

//const notFound = math.Pow(2, 31) - 1

func search(reader ArrayReader, target int) int {
    return searchArr(reader, 0, 10000, target)
}

func searchArr(reader ArrayReader, start, end int, target int) int {
    mid := start + (end - start) / 2
    val := reader.get(mid)
    if val == target {
        return mid
    }
    if val > target {
        if mid == 0 || reader.get(mid-1) < target {
            return -1
        }
        return searchArr(reader, start, mid, target)
    }
    return searchArr(reader, mid+1, end, target)
}