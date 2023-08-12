import "math/rand"

type Solution struct {
    arr []int
}


func Constructor(nums []int) Solution {
    return Solution{
        arr: nums,
    }
}


func (this *Solution) Reset() []int {
    return this.arr
}


func (this *Solution) Shuffle() []int {
    arr := make([]int, len(this.arr))
    copy(arr, this.arr)
    for i := 0; i < len(this.arr); i++ {
        idx := rand.Intn(len(this.arr) - i) + i
        arr[i], arr[idx] = arr[idx], arr[i]
    }
    return arr
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */