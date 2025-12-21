package main

// 3781. Maximum Score After Binary Swaps
// You are given an integer array nums of length n and a binary string s of the same length.

// Initially, your score is 0. Each index i where s[i] = '1' contributes nums[i] to the score.

// You may perform any number of operations (including zero). 
// In one operation, you may choose an index i such that 0 <= i < n - 1, where s[i] = '0', and s[i + 1] = '1', and swap these two characters.

// Return an integer denoting the maximum possible score you can achieve.

// Example 1:
// Input: nums = [2,1,5,2,3], s = "01010"
// Output: 7
// Explanation:
// We can perform the following swaps:
// Swap at index i = 0: "01010" changes to "10010"
// Swap at index i = 2: "10010" changes to "10100"
// Positions 0 and 2 contain '1', contributing nums[0] + nums[2] = 2 + 5 = 7. This is maximum score achievable.

// Example 2:
// Input: nums = [4,7,2,9], s = "0000"
// Output: 0
// Explanation:
// There are no '1' characters in s, so no swaps can be performed. The score remains 0.

// Constraints:
//     n == nums.length == s.length
//     1 <= n <= 105
//     1 <= nums[i] <= 109
//     s[i] is either '0' or '1'

import "fmt"
import "sort"
import "slices"
import "container/heap"

type MinHeap struct{ sort.IntSlice }

func (h *MinHeap) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *MinHeap) Pop() (_ any)  { return }

// time: O(n log n), space: O(n)
func maximumScore(nums []int, s string) int64 {
    res, h := 0, MinHeap{}
    for i, v := range slices.Backward(nums) {
        if s[i] == '1' {
            res += v
            heap.Push(&h, v)
        } else if h.Len() > 0 && v > h.IntSlice[0] {
            res += (v - h.IntSlice[0])
            h.IntSlice[0] = v
            heap.Fix(&h, 0)
        }
    }
    return int64(res)
}


type PriorityQueue []int

func (q *PriorityQueue) Len() int           { return len(*q) }
func (q *PriorityQueue) Less(i, j int) bool { return (*q)[j] < (*q)[i] }
func (q *PriorityQueue) Swap(i, j int)      { (*q)[i], (*q)[j] = (*q)[j], (*q)[i] }
func (q *PriorityQueue) Push(x any)         { *q = append(*q, x.(int)) }
func (q *PriorityQueue) Pop() any {
    l := len(*q) - 1
    res := (*q)[l]
    *q = (*q)[:l]
    return res
}

// time: O(n log n), space: O(n)
func maximumScore1(nums []int, s string) int64 {
    res, pq := 0, PriorityQueue{}
    for i, num := range nums {
        heap.Push(&pq, num) 
        if s[i] == '1' {
            res += heap.Pop(&pq).(int)
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [2,1,5,2,3], s = "01010"
    // Output: 7
    // Explanation:
    // We can perform the following swaps:
    // Swap at index i = 0: "01010" changes to "10010"
    // Swap at index i = 2: "10010" changes to "10100"
    // Positions 0 and 2 contain '1', contributing nums[0] + nums[2] = 2 + 5 = 7. This is maximum score achievable.
    fmt.Println(maximumScore([]int{2,1,5,2,3}, "01010")) // 7
    // Example 2:
    // Input: nums = [4,7,2,9], s = "0000"
    // Output: 0
    // Explanation:
    // There are no '1' characters in s, so no swaps can be performed. The score remains 0.
    fmt.Println(maximumScore([]int{4,7,2,9}, "0000")) // 0

    fmt.Println(maximumScore([]int{1,2,3,4,5,6,7,8,9}, "010101010")) // 20
    fmt.Println(maximumScore([]int{1,2,3,4,5,6,7,8,9}, "101010101")) // 25
    fmt.Println(maximumScore([]int{1,2,3,4,5,6,7,8,9}, "000000000")) // 0
    fmt.Println(maximumScore([]int{1,2,3,4,5,6,7,8,9}, "111111111")) // 45
    fmt.Println(maximumScore([]int{1,2,3,4,5,6,7,8,9}, "000011111")) // 35
    fmt.Println(maximumScore([]int{1,2,3,4,5,6,7,8,9}, "111100000")) // 10
    fmt.Println(maximumScore([]int{1,2,3,4,5,6,7,8,9}, "111000111")) // 30
    fmt.Println(maximumScore([]int{1,2,3,4,5,6,7,8,9}, "000111000")) // 15
    fmt.Println(maximumScore([]int{9,8,7,6,5,4,3,2,1}, "010101010")) // 30
    fmt.Println(maximumScore([]int{9,8,7,6,5,4,3,2,1}, "101010101")) // 35
    fmt.Println(maximumScore([]int{9,8,7,6,5,4,3,2,1}, "000000000")) // 0
    fmt.Println(maximumScore([]int{9,8,7,6,5,4,3,2,1}, "111111111")) // 45
    fmt.Println(maximumScore([]int{9,8,7,6,5,4,3,2,1}, "000011111")) // 35
    fmt.Println(maximumScore([]int{9,8,7,6,5,4,3,2,1}, "111100000")) // 30
    fmt.Println(maximumScore([]int{9,8,7,6,5,4,3,2,1}, "111000111")) // 39
    fmt.Println(maximumScore([]int{9,8,7,6,5,4,3,2,1}, "000111000")) // 24

    fmt.Println(maximumScore1([]int{2,1,5,2,3}, "01010")) // 7
    fmt.Println(maximumScore1([]int{4,7,2,9}, "0000")) // 0
    fmt.Println(maximumScore1([]int{1,2,3,4,5,6,7,8,9}, "010101010")) // 20
    fmt.Println(maximumScore1([]int{1,2,3,4,5,6,7,8,9}, "101010101")) // 25
    fmt.Println(maximumScore1([]int{1,2,3,4,5,6,7,8,9}, "000000000")) // 0
    fmt.Println(maximumScore1([]int{1,2,3,4,5,6,7,8,9}, "111111111")) // 45
    fmt.Println(maximumScore1([]int{1,2,3,4,5,6,7,8,9}, "000011111")) // 35
    fmt.Println(maximumScore1([]int{1,2,3,4,5,6,7,8,9}, "111100000")) // 10
    fmt.Println(maximumScore1([]int{1,2,3,4,5,6,7,8,9}, "111000111")) // 30
    fmt.Println(maximumScore1([]int{1,2,3,4,5,6,7,8,9}, "000111000")) // 15
    fmt.Println(maximumScore1([]int{9,8,7,6,5,4,3,2,1}, "010101010")) // 30
    fmt.Println(maximumScore1([]int{9,8,7,6,5,4,3,2,1}, "101010101")) // 35
    fmt.Println(maximumScore1([]int{9,8,7,6,5,4,3,2,1}, "000000000")) // 0
    fmt.Println(maximumScore1([]int{9,8,7,6,5,4,3,2,1}, "111111111")) // 45
    fmt.Println(maximumScore1([]int{9,8,7,6,5,4,3,2,1}, "000011111")) // 35
    fmt.Println(maximumScore1([]int{9,8,7,6,5,4,3,2,1}, "111100000")) // 30
    fmt.Println(maximumScore1([]int{9,8,7,6,5,4,3,2,1}, "111000111")) // 39
    fmt.Println(maximumScore1([]int{9,8,7,6,5,4,3,2,1}, "000111000")) // 24
}