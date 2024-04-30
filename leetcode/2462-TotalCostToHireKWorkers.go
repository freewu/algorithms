package main

// 2462. Total Cost to Hire K Workers
// You are given a 0-indexed integer array costs where costs[i] is the cost of hiring the ith worker.
// You are also given two integers k and candidates. 
// We want to hire exactly k workers according to the following rules:
//     You will run k sessions and hire exactly one worker in each session.
//     In each hiring session, choose the worker with the lowest cost from either the first candidates workers or the last candidates workers. Break the tie by the smallest index.
//         For example, if costs = [3,2,7,7,1,2] and candidates = 2, then in the first hiring session, we will choose the 4th worker because they have the lowest cost [3,2,7,7,1,2].
//         In the second hiring session, we will choose 1st worker because they have the same lowest cost as 4th worker but they have the smallest index [3,2,7,7,2]. Please note that the indexing may be changed in the process.
//     If there are fewer than candidates workers remaining, choose the worker with the lowest cost among them. Break the tie by the smallest index.
//     A worker can only be chosen once.

// Return the total cost to hire exactly k workers.

// Example 1:
// Input: costs = [17,12,10,2,7,2,11,20,8], k = 3, candidates = 4
// Output: 11
// Explanation: We hire 3 workers in total. The total cost is initially 0.
// - In the first hiring round we choose the worker from [17,12,10,2,7,2,11,20,8]. The lowest cost is 2, and we break the tie by the smallest index, which is 3. The total cost = 0 + 2 = 2.
// - In the second hiring round we choose the worker from [17,12,10,7,2,11,20,8]. The lowest cost is 2 (index 4). The total cost = 2 + 2 = 4.
// - In the third hiring round we choose the worker from [17,12,10,7,11,20,8]. The lowest cost is 7 (index 3). The total cost = 4 + 7 = 11. Notice that the worker with index 3 was common in the first and last four workers.
// The total hiring cost is 11.

// Example 2:
// Input: costs = [1,2,4,1], k = 3, candidates = 3
// Output: 4
// Explanation: We hire 3 workers in total. The total cost is initially 0.
// - In the first hiring round we choose the worker from [1,2,4,1]. The lowest cost is 1, and we break the tie by the smallest index, which is 0. The total cost = 0 + 1 = 1. Notice that workers with index 1 and 2 are common in the first and last 3 workers.
// - In the second hiring round we choose the worker from [2,4,1]. The lowest cost is 1 (index 2). The total cost = 1 + 1 = 2.
// - In the third hiring round there are less than three candidates. We choose the worker from the remaining workers [2,4]. The lowest cost is 2 (index 0). The total cost = 2 + 2 = 4.
// The total hiring cost is 4.
 
// Constraints:
//     1 <= costs.length <= 10^5
//     1 <= costs[i] <= 10^5
//     1 <= k, candidates <= costs.length

import "fmt"
import "container/heap"
import "sort"

type MinHeap []int

func (h MinHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Len() int{
    return len(h)
}

func (h *MinHeap) Push(v any) {
    *h = append(*h, v.(int))
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func (h MinHeap) Peek() any {
    return h[0]
}

func totalCost(costs []int, k int, candidates int) int64 {
    set1, set2, res, i, j := &MinHeap{}, &MinHeap{}, 0, 0, len(costs) - 1
    for i < candidates {
        heap.Push(set1, costs[i])
        i++
    }
    for len(costs) - j <= candidates && j >= i {
        heap.Push(set2, costs[j])
        j--
    }
    for k > 0 && (set1.Len() > 0 || set2.Len() > 0) {
        if set1.Len() <= 0 {
            res += heap.Pop(set2).(int)
            
        } else if set2.Len() <= 0 {
            res += heap.Pop(set1).(int)
        } else {
            if set1.Peek().(int) <= set2.Peek().(int) {
                res += heap.Pop(set1).(int)
                if i <= j {
                    heap.Push(set1, costs[i])
                    i++
                }
            } else {
                res += heap.Pop(set2).(int)
                if j >= i {
                    heap.Push(set2, costs[j])
                    j--
                }
            }
        }
        k--
    }
    return int64(res)
}

type hp struct{ sort.IntSlice }
func(hp) Push(interface{}) {}
func (hp) Pop() (_ interface{}) {return}

func totalCost1(costs []int, k, candidates int) int64 {
    res := 0
    if n := len(costs); candidates * 2 < n {
        pre := hp{costs[:candidates]}
        heap.Init(&pre)
        suf := hp{costs[n-candidates:]}
        heap.Init(suf)
        for i, j := candidates, n - 1 - candidates; k >0 && i <= j; k-- {
            if pre.IntSlice[0] <= suf.IntSlice[0] {
                res += pre.IntSlice[0]
                pre.IntSlice[0] = costs[i]
                heap.Fix(&pre, 0)
                i++
            } else {
                res += suf.IntSlice[0]
                suf.IntSlice[0] = costs[j]
                heap.Fix(&suf, 0)
                j--
            }
        }
        costs = append(pre.IntSlice, suf.IntSlice...)
    }
    sort.Ints(costs)
    for _, c := range costs[:k] {
        res += c
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: costs = [17,12,10,2,7,2,11,20,8], k = 3, candidates = 4
    // Output: 11
    // Explanation: We hire 3 workers in total. The total cost is initially 0.
    // - In the first hiring round we choose the worker from [17,12,10,2,7,2,11,20,8]. The lowest cost is 2, and we break the tie by the smallest index, which is 3. The total cost = 0 + 2 = 2.
    // - In the second hiring round we choose the worker from [17,12,10,7,2,11,20,8]. The lowest cost is 2 (index 4). The total cost = 2 + 2 = 4.
    // - In the third hiring round we choose the worker from [17,12,10,7,11,20,8]. The lowest cost is 7 (index 3). The total cost = 4 + 7 = 11. Notice that the worker with index 3 was common in the first and last four workers.
    // The total hiring cost is 11.
    fmt.Println(totalCost([]int{17,12,10,2,7,2,11,20,8}, 3, 4)) // 11
    // Example 2:
    // Input: costs = [1,2,4,1], k = 3, candidates = 3
    // Output: 4
    // Explanation: We hire 3 workers in total. The total cost is initially 0.
    // - In the first hiring round we choose the worker from [1,2,4,1]. The lowest cost is 1, and we break the tie by the smallest index, which is 0. The total cost = 0 + 1 = 1. Notice that workers with index 1 and 2 are common in the first and last 3 workers.
    // - In the second hiring round we choose the worker from [2,4,1]. The lowest cost is 1 (index 2). The total cost = 1 + 1 = 2.
    // - In the third hiring round there are less than three candidates. We choose the worker from the remaining workers [2,4]. The lowest cost is 2 (index 0). The total cost = 2 + 2 = 4.
    // The total hiring cost is 4.
    fmt.Println(totalCost([]int{1,2,4,1}, 3, 3)) // 4

    fmt.Println(totalCost1([]int{17,12,10,2,7,2,11,20,8}, 3, 4)) // 11
    fmt.Println(totalCost1([]int{1,2,4,1}, 3, 3)) // 4
}