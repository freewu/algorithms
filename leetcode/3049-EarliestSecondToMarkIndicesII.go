package main

// 3049. Earliest Second to Mark Indices II
// You are given two 1-indexed integer arrays, nums and, changeIndices, having lengths n and m, respectively.

// Initially, all indices in nums are unmarked. 
// Your task is to mark all indices in nums.

// In each second, s, in order from 1 to m (inclusive), you can perform one of the following operations:
//     1. Choose an index i in the range [1, n] and decrement nums[i] by 1.
//     2. Set nums[changeIndices[s]] to any non-negative value.
//     3. Choose an index i in the range [1, n], where nums[i] is equal to 0, and mark index i.
//     4. Do nothing.

// Return an integer denoting the earliest second in the range [1, m] when all indices in nums can be marked by choosing operations optimally, or -1 if it is impossible.

// Example 1:
// Input: nums = [3,2,3], changeIndices = [1,3,2,2,2,2,3]
// Output: 6
// Explanation: In this example, we have 7 seconds. The following operations can be performed to mark all indices:
// Second 1: Set nums[changeIndices[1]] to 0. nums becomes [0,2,3].
// Second 2: Set nums[changeIndices[2]] to 0. nums becomes [0,2,0].
// Second 3: Set nums[changeIndices[3]] to 0. nums becomes [0,0,0].
// Second 4: Mark index 1, since nums[1] is equal to 0.
// Second 5: Mark index 2, since nums[2] is equal to 0.
// Second 6: Mark index 3, since nums[3] is equal to 0.
// Now all indices have been marked.
// It can be shown that it is not possible to mark all indices earlier than the 6th second.
// Hence, the answer is 6.

// Example 2:
// Input: nums = [0,0,1,2], changeIndices = [1,2,1,2,1,2,1,2]
// Output: 7
// Explanation: In this example, we have 8 seconds. The following operations can be performed to mark all indices:
// Second 1: Mark index 1, since nums[1] is equal to 0.
// Second 2: Mark index 2, since nums[2] is equal to 0.
// Second 3: Decrement index 4 by one. nums becomes [0,0,1,1].
// Second 4: Decrement index 4 by one. nums becomes [0,0,1,0].
// Second 5: Decrement index 3 by one. nums becomes [0,0,0,0].
// Second 6: Mark index 3, since nums[3] is equal to 0.
// Second 7: Mark index 4, since nums[4] is equal to 0.
// Now all indices have been marked.
// It can be shown that it is not possible to mark all indices earlier than the 7th second.
// Hence, the answer is 7.

// Example 3:
// Input: nums = [1,2,3], changeIndices = [1,2,3]
// Output: -1
// Explanation: In this example, it can be shown that it is impossible to mark all indices, as we don't have enough seconds. 
// Hence, the answer is -1.

// Constraints:
//     1 <= n == nums.length <= 5000
//     0 <= nums[i] <= 10^9
//     1 <= m == changeIndices.length <= 5000
//     1 <= changeIndices[i] <= n

import "fmt"
import "container/heap"

type PriorityQueue []int

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Top() interface{}   { return pq[0] }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i] < pq[j] }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(int))}
func (pq *PriorityQueue) Pop() interface{} {
    var item interface{}
    n := len(*pq)
    *pq, item = (*pq)[:n-1], (*pq)[n-1]
    return item
}

func earliestSecondToMarkIndices(nums []int, changeIndices []int) int {
    ops, n, m := 0, len(nums), len(changeIndices)
    for i := 0; i < n; i++ {
        ops += nums[i] + 1
    }
    canMarkAll := func(nums, changeIndices []int, end int, ops int) bool {
        seen, taken := make(map[int]struct{}), make(map[int]bool)
        for i := 0; i < end; i++ {
            if _, ok := seen[changeIndices[i]]; !ok && nums[changeIndices[i]-1] > 0 {
                seen[changeIndices[i]] = struct{}{}
                taken[i] = true
            }
        }
        var pq PriorityQueue
        var free int
        for i := end-1; i >= 0; i-- {
            if !taken[i] {
                free++
            } else {
                heap.Push(&pq, nums[changeIndices[i]-1])
                if free > 0 {
                    free--
                } else {
                    free++
                    heap.Pop(&pq)
                }
            }
        }
        for pq.Len() > 0 {
            ops -= pq.Top().(int)
            heap.Pop(&pq)
            ops--
        }
        return ops <= free
    }
    res, left, right := -1, 1, m
    for left <= right {
        mid := left + (right - left) /2 
        if canMarkAll(nums, changeIndices, mid, ops) {
            res, right = mid, mid - 1
        } else {
            left = mid + 1
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [3,2,3], changeIndices = [1,3,2,2,2,2,3]
    // Output: 6
    // Explanation: In this example, we have 7 seconds. The following operations can be performed to mark all indices:
    // Second 1: Set nums[changeIndices[1]] to 0. nums becomes [0,2,3].
    // Second 2: Set nums[changeIndices[2]] to 0. nums becomes [0,2,0].
    // Second 3: Set nums[changeIndices[3]] to 0. nums becomes [0,0,0].
    // Second 4: Mark index 1, since nums[1] is equal to 0.
    // Second 5: Mark index 2, since nums[2] is equal to 0.
    // Second 6: Mark index 3, since nums[3] is equal to 0.
    // Now all indices have been marked.
    // It can be shown that it is not possible to mark all indices earlier than the 6th second.
    // Hence, the answer is 6.
    fmt.Println(earliestSecondToMarkIndices([]int{3,2,3}, []int{1,3,2,2,2,2,3})) // 6
    // Example 2:
    // Input: nums = [0,0,1,2], changeIndices = [1,2,1,2,1,2,1,2]
    // Output: 7
    // Explanation: In this example, we have 8 seconds. The following operations can be performed to mark all indices:
    // Second 1: Mark index 1, since nums[1] is equal to 0.
    // Second 2: Mark index 2, since nums[2] is equal to 0.
    // Second 3: Decrement index 4 by one. nums becomes [0,0,1,1].
    // Second 4: Decrement index 4 by one. nums becomes [0,0,1,0].
    // Second 5: Decrement index 3 by one. nums becomes [0,0,0,0].
    // Second 6: Mark index 3, since nums[3] is equal to 0.
    // Second 7: Mark index 4, since nums[4] is equal to 0.
    // Now all indices have been marked.
    // It can be shown that it is not possible to mark all indices earlier than the 7th second.
    // Hence, the answer is 7.
    fmt.Println(earliestSecondToMarkIndices([]int{0,0,1,2}, []int{1,2,1,2,1,2,1,2})) // 7
    // Example 3:
    // Input: nums = [1,2,3], changeIndices = [1,2,3]
    // Output: -1
    // Explanation: In this example, it can be shown that it is impossible to mark all indices, as we don't have enough seconds. 
    // Hence, the answer is -1.
    fmt.Println(earliestSecondToMarkIndices([]int{1,2,3}, []int{1,2,3})) // -1
}