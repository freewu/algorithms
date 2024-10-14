package main

// 2530. Maximal Score After Applying K Operations
// You are given a 0-indexed integer array nums and an integer k. 
// You have a starting score of 0.

// In one operation:
//     choose an index i such that 0 <= i < nums.length,
//     increase your score by nums[i], and
//     replace nums[i] with ceil(nums[i] / 3).

// Return the maximum possible score you can attain after applying exactly k operations.

// The ceiling function ceil(val) is the least integer greater than or equal to val.

// Example 1:
// Input: nums = [10,10,10,10,10], k = 5
// Output: 50
// Explanation: Apply the operation to each array element exactly once. The final score is 10 + 10 + 10 + 10 + 10 = 50.

// Example 2:
// Input: nums = [1,10,3,3,3], k = 3
// Output: 17
// Explanation: You can do the following operations:
// Operation 1: Select i = 1, so nums becomes [1,4,3,3,3]. Your score increases by 10.
// Operation 2: Select i = 1, so nums becomes [1,2,3,3,3]. Your score increases by 4.
// Operation 3: Select i = 2, so nums becomes [1,1,1,3,3]. Your score increases by 3.
// The final score is 10 + 4 + 3 = 17.

// Constraints:
//     1 <= nums.length, k <= 10^5
//     1 <= nums[i] <= 10^9

import "fmt"
import "math"
import "container/heap"

type MaxHeap struct {
    nums []int
}

func (h MaxHeap) Len() int { return len(h.nums) }
func (h MaxHeap) Less(i, j int) bool { return h.nums[i] > h.nums[j] }
func (h MaxHeap) Swap(i, j int) { h.nums[i], h.nums[j] = h.nums[j], h.nums[i] }
func (h *MaxHeap) Push(v interface{}) { h.nums = append(h.nums, v.(int)) }
func (h *MaxHeap) Pop() interface{} {
    v := h.nums[len(h.nums)-1]
    h.nums = h.nums[:len(h.nums)-1]
    return v
}

func maxKelements(nums []int, k int) int64 {
    maxHeap := &MaxHeap{ nums: nums }
    heap.Init(maxHeap)
    res := 0
    for i := 0; i < k; i++ {
        mx := maxHeap.nums[0]
        res += mx
        maxHeap.nums[0] = int(math.Ceil(float64(mx) / 3.0))
        heap.Fix(maxHeap, 0)
    }
    return int64(res)
}


func main() {
    // Example 1:
    // Input: nums = [10,10,10,10,10], k = 5
    // Output: 50
    // Explanation: Apply the operation to each array element exactly once. The final score is 10 + 10 + 10 + 10 + 10 = 50.
    fmt.Println(maxKelements([]int{10,10,10,10,10}, 5)) // 50
    // Example 2:
    // Input: nums = [1,10,3,3,3], k = 3
    // Output: 17
    // Explanation: You can do the following operations:
    // Operation 1: Select i = 1, so nums becomes [1,4,3,3,3]. Your score increases by 10.
    // Operation 2: Select i = 1, so nums becomes [1,2,3,3,3]. Your score increases by 4.
    // Operation 3: Select i = 2, so nums becomes [1,1,1,3,3]. Your score increases by 3.
    // The final score is 10 + 4 + 3 = 17.
    fmt.Println(maxKelements([]int{1,10,3,3,3}, 3)) // 17
}