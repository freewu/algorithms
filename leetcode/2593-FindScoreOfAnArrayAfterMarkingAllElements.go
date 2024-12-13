package main

// 2593. Find Score of an Array After Marking All Elements
// You are given an array nums consisting of positive integers.

// Starting with score = 0, apply the following algorithm:
//     1. Choose the smallest integer of the array that is not marked. 
//        If there is a tie, choose the one with the smallest index.
//     2. Add the value of the chosen integer to score.
//     3. Mark the chosen element and its two adjacent elements if they exist.
//     4. Repeat until all the array elements are marked.

// Return the score you get after applying the above algorithm.

// Example 1:
// Input: nums = [2,1,3,4,5,2]
// Output: 7
// Explanation: We mark the elements as follows:
// - 1 is the smallest unmarked element, so we mark it and its two adjacent elements: [2,1,3,4,5,2].
// - 2 is the smallest unmarked element, so we mark it and its left adjacent element: [2,1,3,4,5,2].
// - 4 is the only remaining unmarked element, so we mark it: [2,1,3,4,5,2].
// Our score is 1 + 2 + 4 = 7.

// Example 2:
// Input: nums = [2,3,5,1,3,2]
// Output: 5
// Explanation: We mark the elements as follows:
// - 1 is the smallest unmarked element, so we mark it and its two adjacent elements: [2,3,5,1,3,2].
// - 2 is the smallest unmarked element, since there are two of them, we choose the left-most one, so we mark the one at index 0 and its right adjacent element: [2,3,5,1,3,2].
// - 2 is the only remaining unmarked element, so we mark it: [2,3,5,1,3,2].
// Our score is 1 + 2 + 2 = 5.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^6

import "fmt"
import "sort"
import "container/heap"

type Pair struct{ x, i int }
type MinHeap []Pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].x < h[j].x || (h[i].x == h[j].x && h[i].i < h[j].i) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v any)        { *h = append(*h, v.(Pair)) }
func (h *MinHeap) Pop() any          { a := *h; v := a[len(a) - 1]; *h = a[:len(a) - 1]; return v }

func findScore(nums []int) int64 {
    mnh := MinHeap{}
    for i, v := range nums {
        heap.Push(&mnh, Pair{ v, i })
    }
    res, n := 0, len(nums)
    visited := make([]bool, n)
    for len(mnh) > 0 {
        p := heap.Pop(&mnh).(Pair)
        res += p.x
        visited[p.i] = true
        for _, j := range []int{ p.i - 1, p.i + 1} {
            if j >= 0 && j < n {
                visited[j] = true
            }
        }
        for len(mnh) > 0 && visited[mnh[0].i] {
            heap.Pop(&mnh)
        }
    }
    return int64(res)
}


func findScore1(nums []int) int64 {
    type Pair struct { x, i int }
    res, n := 0, len(nums)
    sorted, visited := make([]Pair, n), make([]bool, n) // false-sz be default
    for i, v := range nums {
        sorted[i] = Pair{ v, i }
    }
    sort.Slice(sorted, func(i, j int) bool {
        if sorted[i].x == sorted[j].x { return sorted[i].i < sorted[j].i }
        return sorted[i].x < sorted[j].x
    })
    for _, v := range sorted {
        if !visited[v.i] {
            visited[v.i] = true
            res += v.x
            // Mark adjacent elements
            if v.i > 0 { visited[v.i - 1] = true }
            if v.i < n - 1 { visited[v.i + 1] = true }
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [2,1,3,4,5,2]
    // Output: 7
    // Explanation: We mark the elements as follows:
    // - 1 is the smallest unmarked element, so we mark it and its two adjacent elements: [2,1,3,4,5,2].
    // - 2 is the smallest unmarked element, so we mark it and its left adjacent element: [2,1,3,4,5,2].
    // - 4 is the only remaining unmarked element, so we mark it: [2,1,3,4,5,2].
    // Our score is 1 + 2 + 4 = 7.
    fmt.Println(findScore([]int{2,1,3,4,5,2})) // 7
    // Example 2:
    // Input: nums = [2,3,5,1,3,2]
    // Output: 5
    // Explanation: We mark the elements as follows:
    // - 1 is the smallest unmarked element, so we mark it and its two adjacent elements: [2,3,5,1,3,2].
    // - 2 is the smallest unmarked element, since there are two of them, we choose the left-most one, so we mark the one at index 0 and its right adjacent element: [2,3,5,1,3,2].
    // - 2 is the only remaining unmarked element, so we mark it: [2,3,5,1,3,2].
    // Our score is 1 + 2 + 2 = 5.
    fmt.Println(findScore([]int{2,3,5,1,3,2})) // 5

    fmt.Println(findScore1([]int{2,1,3,4,5,2})) // 7
    fmt.Println(findScore1([]int{2,3,5,1,3,2})) // 5
}