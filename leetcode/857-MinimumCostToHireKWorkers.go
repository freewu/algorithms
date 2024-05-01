package main

// 857. Minimum Cost to Hire K Workers
// There are n workers. You are given two integer arrays quality 
// and wage where quality[i] is the quality of the ith worker and wage[i] is the minimum wage expectation for the ith worker.

// We want to hire exactly k workers to form a paid group. 
// To hire a group of k workers, we must pay them according to the following rules:
//     Every worker in the paid group should be paid in the ratio of their quality compared to other workers in the paid group.
//     Every worker in the paid group must be paid at least their minimum wage expectation.

// Given the integer k, return the least amount of money needed to form a paid group satisfying the above conditions. 
// Answers within 10^-5 of the actual answer will be accepted.

// Example 1:
// Input: quality = [10,20,5], wage = [70,50,30], k = 2
// Output: 105.00000
// Explanation: We pay 70 to 0th worker and 35 to 2nd worker.

// Example 2:
// Input: quality = [3,1,10,10,1], wage = [4,8,2,2,7], k = 3
// Output: 30.66667
// Explanation: We pay 4 to 0th worker, 13.33333 to 2nd and 3rd workers separately.

// Constraints:
//     n == quality.length == wage.length
//     1 <= k <= n <= 10^4
//     1 <= quality[i], wage[i] <= 10^4

import "fmt"
import "math"
import "sort"
import "container/heap"

// MaxHeap
type MaxHeap struct{ sort.IntSlice }
func (h MaxHeap) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (MaxHeap) Push(interface{})     {} // 由于没有用到，可以什么都不写
func (MaxHeap) Pop() (_ interface{}) { return }

func mincostToHireWorkers(quality, wage []int, k int) float64 {
    type pair struct{ q, w int }
    qw := make([]pair, len(quality))
    for i, q := range quality {
        qw[i] = pair{q, wage[i]}
    }
    sort.Slice(qw, func(i, j int) bool { a, b := qw[i], qw[j]; return a.w*b.q < b.w*a.q }) // 按照 r 值排序
    h := MaxHeap{make([]int, k)}
    sum := 0
    for i, p := range qw[:k] {
        h.IntSlice[i] = p.q
        sum += p.q
    }
    heap.Init(&h)
    res := float64(sum * qw[k-1].w) / float64(qw[k-1].q) // 选 r 值最小的 k 名工人组成当前的最优解
    for _, p := range qw[k:] {
        if p.q < h.IntSlice[0] { // sum 可以变小，从而可能得到更优的答案
            sum -= h.IntSlice[0] - p.q
            h.IntSlice[0] = p.q
            heap.Fix(&h, 0) // 更新堆顶
            res = math.Min(res, float64(sum * p.w) / float64(p.q))
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: quality = [10,20,5], wage = [70,50,30], k = 2
    // Output: 105.00000
    // Explanation: We pay 70 to 0th worker and 35 to 2nd worker.
    fmt.Println(mincostToHireWorkers([]int{10,20,5},[]int{70,50,30},2)) // 105.00000
    // Example 2:
    // Input: quality = [3,1,10,10,1], wage = [4,8,2,2,7], k = 3
    // Output: 30.66667
    // Explanation: We pay 4 to 0th worker, 13.33333 to 2nd and 3rd workers separately.
    fmt.Println(mincostToHireWorkers([]int{3,1,10,10,1},[]int{4,8,2,2,7},3)) // 30.66667
}