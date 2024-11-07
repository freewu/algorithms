package main

// 3264. Final Array State After K Multiplication Operations I
// You are given an integer array nums, an integer k, and an integer multiplier.

// You need to perform k operations on nums. In each operation:
//     1. Find the minimum value x in nums. If there are multiple occurrences of the minimum value, 
//        select the one that appears first.
//     2. Replace the selected minimum value x with x * multiplier.

// Return an integer array denoting the final state of nums after performing all k operations.

// Example 1:
// Input: nums = [2,1,3,5,6], k = 5, multiplier = 2
// Output: [8,4,6,5,6]
// Explanation:
// Operation	Result
// After operation 1	[2, 2, 3, 5, 6]
// After operation 2	[4, 2, 3, 5, 6]
// After operation 3	[4, 4, 3, 5, 6]
// After operation 4	[4, 4, 6, 5, 6]
// After operation 5	[8, 4, 6, 5, 6]

// Example 2:
// Input: nums = [1,2], k = 3, multiplier = 4
// Output: [16,8]
// Explanation:
// Operation	Result
// After operation 1	[4, 2]
// After operation 2	[4, 8]
// After operation 3	[16, 8]

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 100
//     1 <= k <= 10
//     1 <= multiplier <= 5

import "fmt"
import "container/heap"

func getFinalState(nums []int, k int, multiplier int) []int {
    for i := 0; i < k; i++ {
        value, index := nums[0], 0
        for j := 1; j < len(nums); j++ {
            if value > nums[j] {
                index, value = j, nums[j]
            }
        }
        nums[index] = value * multiplier
    }
    return nums
}

type Pair struct {
    index int
    value int
}
type MinHeap []Pair

func (p MinHeap) Len() int            { return len(p) }
func (p MinHeap) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p MinHeap) Less(i, j int) bool  { 
    if p[i].value == p[j].value { return p[i].index < p[j].index }
    return p[i].value < p[j].value
}
func (p *MinHeap) Push(v interface{}) { *p = append(*p, v.(Pair)) }
func (p *MinHeap) Pop() interface{} {
    a := *p
    res := a[len(a) - 1]
    *p = a[:len(a) - 1]
    return res
}

func getFinalState1(nums []int, k int, multiplier int) []int {
    mnh := MinHeap{}
    heap.Init(&mnh)
    for i := range nums {
        heap.Push(&mnh, Pair{i, nums[i]})
    }
    for k > 0 {
        t := heap.Pop(&mnh).(Pair)
        t.value *= multiplier
        heap.Push(&mnh, t)
        k--
    }
    res := make([]int, len(nums))
    for i := range mnh {
        res[mnh[i].index] = mnh[i].value
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,1,3,5,6], k = 5, multiplier = 2
    // Output: [8,4,6,5,6]
    // Explanation:
    // Operation	Result
    // After operation 1	[2, 2, 3, 5, 6]
    // After operation 2	[4, 2, 3, 5, 6]
    // After operation 3	[4, 4, 3, 5, 6]
    // After operation 4	[4, 4, 6, 5, 6]
    // After operation 5	[8, 4, 6, 5, 6]
    fmt.Println(getFinalState([]int{2,1,3,5,6}, 5, 2)) // [8,4,6,5,6]
    // Example 2:
    // Input: nums = [1,2], k = 3, multiplier = 4
    // Output: [16,8]
    // Explanation:
    // Operation	Result
    // After operation 1	[4, 2]
    // After operation 2	[4, 8]
    // After operation 3	[16, 8]
    fmt.Println(getFinalState([]int{1,2}, 3, 4)) // [16,8]

    fmt.Println(getFinalState1([]int{2,1,3,5,6}, 5, 2)) // [8,4,6,5,6]
    fmt.Println(getFinalState1([]int{1,2}, 3, 4)) // [16,8]
}