package main

// LCR 071. 按权重随机选择
// 给定一个正整数数组 w ，其中 w[i] 代表下标 i 的权重（下标从 0 开始），
// 请写一个函数 pickIndex ，它可以随机地获取下标 i，选取下标 i 的概率与 w[i] 成正比。

// 例如，对于 w = [1, 3]，挑选下标 0 的概率为 1 / (1 + 3) = 0.25 （即，25%），而选取下标 1 的概率为 3 / (1 + 3) = 0.75（即，75%）。
// 也就是说，选取下标 i 的概率为 w[i] / sum(w) 。


// 示例 1：
// 输入：
// inputs = ["Solution","pickIndex"]
// inputs = [[[1]],[]]
// 输出：
// [null,0]
// 解释：
// Solution solution = new Solution([1]);
// solution.pickIndex(); // 返回 0，因为数组中只有一个元素，所以唯一的选择是返回下标 0。

// 示例 2：
// 输入：
// inputs = ["Solution","pickIndex","pickIndex","pickIndex","pickIndex","pickIndex"]
// inputs = [[[1,3]],[],[],[],[],[]]
// 输出：
// [null,1,1,1,1,0]
// 解释：
// Solution solution = new Solution([1, 3]);
// solution.pickIndex(); // 返回 1，返回下标 1，返回该下标概率为 3/4 。
// solution.pickIndex(); // 返回 1
// solution.pickIndex(); // 返回 1
// solution.pickIndex(); // 返回 1
// solution.pickIndex(); // 返回 0，返回下标 0，返回该下标概率为 1/4 。
// 由于这是一个随机问题，允许多个答案，因此下列输出都可以被认为是正确的:
// [null,1,1,1,1,0]
// [null,1,1,1,1,1]
// [null,1,1,1,0,0]
// [null,1,1,1,0,1]
// [null,1,0,1,0,0]
// ......
// 诸若此类。
 
// 提示：
//     1 <= w.length <= 10000
//     1 <= w[i] <= 10^5
//     pickIndex 将被调用不超过 10000 次

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