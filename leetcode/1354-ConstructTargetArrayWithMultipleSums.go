package main

// 1354. Construct Target Array With Multiple Sums
// You are given an array target of n integers. 
// From a starting array arr consisting of n 1's, you may perform the following procedure :
//     let x be the sum of all elements currently in your array.
//     choose index i, such that 0 <= i < n and set the value of arr at index i to x.
//     You may repeat this procedure as many times as needed.

// Return true if it is possible to construct the target array from arr, otherwise, return false.

// Example 1:
// Input: target = [9,3,5]
// Output: true
// Explanation: Start with arr = [1, 1, 1] 
// [1, 1, 1], sum = 3 choose index 1
// [1, 3, 1], sum = 5 choose index 2
// [1, 3, 5], sum = 9 choose index 0
// [9, 3, 5] Done

// Example 2:
// Input: target = [1,1,1,2]
// Output: false
// Explanation: Impossible to create target array from [1,1,1,1].

// Example 3:
// Input: target = [8,5]
// Output: true

// Constraints:
//     n == target.length
//     1 <= n <= 5 * 10^4
//     1 <= target[i] <= 10^9

import "fmt"
import "container/heap"

func isPossible(t []int) bool {
    target := MaxHeap(t)
    heap.Init(&target)
    pop := heap.Pop(&target).(int)
    sum := 0
    for _, v := range target {
        sum += v
    }
    for pop != 1 && sum != 1{
        if pop < sum || sum == 0 || pop % sum == 0 {
            return false
        }
        pop %= sum
        heap.Push(&target, pop)
        sum += pop
        pop = heap.Pop(&target).(int)
        sum -= pop
    }
    return true
}

type MaxHeap []int
func (h MaxHeap)  Len() int { return len(h) }
func (h MaxHeap)  Less(a, b int) bool{ return h[a] > h[b] }
func (h MaxHeap)  Swap(a, b int) { h[a], h[b] = h[b], h[a] }
func (h *MaxHeap) Push(a interface{}) { *h = append(*h, a.(int)) }
func (h *MaxHeap) Pop() interface{} {
    l := len(*h)
    res := (*h)[l-1]
    *h = (*h)[:l-1]
    return res
}

func main() {
    // Example 1:
    // Input: target = [9,3,5]
    // Output: true
    // Explanation: Start with arr = [1, 1, 1] 
    // [1, 1, 1], sum = 3 choose index 1
    // [1, 3, 1], sum = 5 choose index 2
    // [1, 3, 5], sum = 9 choose index 0
    // [9, 3, 5] Done
    fmt.Println(isPossible([]int{9,3,5})) // true
    // Example 2:
    // Input: target = [1,1,1,2]
    // Output: false
    // Explanation: Impossible to create target array from [1,1,1,1].
    fmt.Println(isPossible([]int{1,1,1,2})) // false
    // Example 3:
    // Input: target = [8,5]
    // Output: true
    fmt.Println(isPossible([]int{8,5})) // true
}