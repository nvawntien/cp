type Solution struct {
    freq map[int][]int
}


func Constructor(nums []int) Solution {
    freq := make(map[int][]int)

    for i, num := range nums {
        freq[num] = append(freq[num], i)    
    }

    return Solution{
        freq: freq,
    }    
}


func (this *Solution) Pick(target int) int {
    arr := this.freq[target]
    n := len(arr)
    i := rand.Intn(n)
    return arr[i]
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */