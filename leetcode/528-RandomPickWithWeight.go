package main

// 528. Random Pick with Weight
// You are given a 0-indexed array of positive integers w where w[i] describes the weight of the ith index.
// You need to implement the function pickIndex(), which randomly picks an index in the range [0, w.length - 1] (inclusive) and returns it. 
// The probability of picking an index i is w[i] / sum(w).

// For example, 
//     if w = [1, 3], the probability of picking index 0 is 1 / (1 + 3) = 0.25 (i.e., 25%), 
//     and the probability of picking index 1 is 3 / (1 + 3) = 0.75 (i.e., 75%).
 
// Example 1:
// Input
// ["Solution","pickIndex"]
// [[[1]],[]]
// Output
// [null,0]
// Explanation
// Solution solution = new Solution([1]);
// solution.pickIndex(); // return 0. The only option is to return 0 since there is only one element in w.

// Example 2:
// Input
// ["Solution","pickIndex","pickIndex","pickIndex","pickIndex","pickIndex"]
// [[[1,3]],[],[],[],[],[]]
// Output
// [null,1,1,1,1,0]
// Explanation
// Solution solution = new Solution([1, 3]);
// solution.pickIndex(); // return 1. It is returning the second element (index = 1) that has a probability of 3/4.
// solution.pickIndex(); // return 1
// solution.pickIndex(); // return 1
// solution.pickIndex(); // return 1
// solution.pickIndex(); // return 0. It is returning the first element (index = 0) that has a probability of 1/4.
// Since this is a randomization problem, multiple answers are allowed.
// All of the following outputs can be considered correct:
// [null,1,1,1,1,0]
// [null,1,1,1,1,1]
// [null,1,1,1,0,0]
// [null,1,1,1,0,1]
// [null,1,0,1,0,0]
// ......
// and so on.

// Constraints:
//     1 <= w.length <= 10^4
//     1 <= w[i] <= 10^5
//     pickIndex will be called at most 10^4 times.

import "fmt"
import "math/rand"
import "sort"

type Solution struct {
	data []int
}

func Constructor(w []int) Solution {
    data := make([]int, len(w))
    data[0] = w[0]
    for i := 1; i < len(w); i++ {
        data[i] = data[i-1] + w[i]
    }
    return Solution{data}
}

// 一个长度为 n 的构造好的「前缀和」数组可以看是一个基本单位为 1 的 [1,sum[n−1]] 数轴。
// 使用随机函数参数产生 [1,sum[n−1]] 范围内的随机数，通过「二分」前缀和数组即可找到分布
// 位置对应的原始下标值。
func (this *Solution) PickIndex() int {
    weight := rand.Intn(this.data[len(this.data)-1]) + 1
    return sort.SearchInts(this.data, weight)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */

func main() {
    // Solution solution = new Solution([1]);
    obj := Constructor([]int{1})
    // solution.pickIndex(); // return 0. The only option is to return 0 since there is only one element in w.
    fmt.Println(obj.PickIndex()) // 0
    // Example 2:
    // Input
    // ["Solution","pickIndex","pickIndex","pickIndex","pickIndex","pickIndex"]
    // [[[1,3]],[],[],[],[],[]]
    // Output
    // [null,1,1,1,1,0]
    // Explanation
    // Solution solution = new Solution([1, 3]);
    obj1 := Constructor([]int{1,3})
    // solution.pickIndex(); // return 1. It is returning the second element (index = 1) that has a probability of 3/4.
    fmt.Println(obj1.PickIndex()) // 1
    // solution.pickIndex(); // return 1
    fmt.Println(obj1.PickIndex()) // 1
    // solution.pickIndex(); // return 1
    fmt.Println(obj1.PickIndex()) // 1
    // solution.pickIndex(); // return 1
    fmt.Println(obj1.PickIndex()) // 1
    // solution.pickIndex(); // return 0. It is returning the first element (index = 0) that has a probability of 1/4.
    fmt.Println(obj1.PickIndex()) // 0
    // Since this is a randomization problem, multiple answers are allowed.
    // All of the following outputs can be considered correct:
    // [null,1,1,1,1,0]
    // [null,1,1,1,1,1]
    // [null,1,1,1,0,0]
    // [null,1,1,1,0,1]
    // [null,1,0,1,0,0]


}