package main

// 2558. Take Gifts From the Richest Pile
// You are given an integer array gifts denoting the number of gifts in various piles. Every second, you do the following:
//     1. Choose the pile with the maximum number of gifts.
//     2. If there is more than one pile with the maximum number of gifts, choose any.
//     3. Leave behind the floor of the square root of the number of gifts in the pile. Take the rest of the gifts.

// Return the number of gifts remaining after k seconds.

// Example 1:
// Input: gifts = [25,64,9,4,100], k = 4
// Output: 29
// Explanation: 
// The gifts are taken in the following way:
// - In the first second, the last pile is chosen and 10 gifts are left behind.
// - Then the second pile is chosen and 8 gifts are left behind.
// - After that the first pile is chosen and 5 gifts are left behind.
// - Finally, the last pile is chosen again and 3 gifts are left behind.
// The final remaining gifts are [5,8,9,4,3], so the total number of gifts remaining is 29.

// Example 2:
// Input: gifts = [1,1,1,1], k = 4
// Output: 4
// Explanation: 
// In this case, regardless which pile you choose, you have to leave behind 1 gift in each pile. 
// That is, you can't take any pile with you. 
// So, the total gifts remaining are 4.

// Constraints:
//     1 <= gifts.length <= 10^3
//     1 <= gifts[i] <= 10^9
//     1 <= k <= 10^3

import "fmt"
import "container/heap"
import "math"

type PriorityQueue []int

func (p PriorityQueue) Len() int { return len(p) }
func (p PriorityQueue) Less(i, j int) bool { return p[i] > p[j] }
func (p PriorityQueue) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p *PriorityQueue) Top() int { return (*p)[0] }
func (p *PriorityQueue) Push(x interface{}) {  *p = append(*p, x.(int)) }
func (p *PriorityQueue) Pop() interface{} {
    v := (*p)[len(*p) - 1]
    *p = (*p)[:len(*p) - 1]
    return v
}

func pickGifts(gifts []int, k int) int64 {
    pq := &PriorityQueue{}
    for _, gift := range gifts {
        heap.Push(pq, gift)
    }
    for k > 0 {
        heap.Push(pq, int(math.Sqrt(float64(heap.Pop(pq).(int)))))
        k--
    }
    res := 0
    for pq.Len() > 0 {
        res += heap.Pop(pq).(int)
    }
    return int64(res)
}

type MaxHeap struct { data []int }
func (h *MaxHeap) Len() int { return len(h.data) }
func (h *MaxHeap) Less(i, j int) bool { return h.data[i] > h.data[j] }
func (h *MaxHeap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *MaxHeap) Top() int { return h.data[0] }
// func (h *MaxHeap) Push(x interface{}) { h.data = append(h.data, x.(int)) }
// func (h *MaxHeap) Pop() interface{} {
//     n := h.Len()
//     v := h.data[n-1]
//     h.data = h.data[:n-1]
//     return v
// }
func (h *MaxHeap) Push(x int) {
    h.data = append(h.data, x)
    h.filterUp(h.Len() - 1)
}
func (h *MaxHeap) Pop() int {
    n := h.Len()
    h.Swap(0, n - 1)
    v := h.data[n - 1]
    h.data = h.data[:n - 1]
    h.filterDown(0)
    return v
}
func (h *MaxHeap) filterUp(t int) {
    for t > 0 {
        p := (t - 1) / 2
        if h.Less(t, p) {
            h.Swap(t, p)
            t = p
        } else {
            break
        }
    }
}
func (h *MaxHeap) filterDown(t int) {
    n, left := h.Len(), 1
    for left < n {
        if left + 1 < n && h.Less(left + 1, left) {
            left += 1
        }
        if !h.Less(left, t) { break }
        h.Swap(left, t)
        t = left
        left = t* 2 + 1
    }
}

func pickGifts1(gifts []int, k int) int64 {
    mxh := &MaxHeap{}
    res, sum := 0, 0
    for _, v := range gifts {
        mxh.Push(v)
        sum += v
    }
    for k > 0 {
        v := mxh.Pop()
        left := int(math.Sqrt(float64(v)))
        mxh.Push(left)
        res += v - left
        k--
    }
    return int64(sum - res)
}

func main() {
    // Example 1:
    // Input: gifts = [25,64,9,4,100], k = 4
    // Output: 29
    // Explanation: 
    // The gifts are taken in the following way:
    // - In the first second, the last pile is chosen and 10 gifts are left behind.
    // - Then the second pile is chosen and 8 gifts are left behind.
    // - After that the first pile is chosen and 5 gifts are left behind.
    // - Finally, the last pile is chosen again and 3 gifts are left behind.
    // The final remaining gifts are [5,8,9,4,3], so the total number of gifts remaining is 29.
    fmt.Println(pickGifts([]int{25,64,9,4,100}, 4)) // 29
    // Example 2:
    // Input: gifts = [1,1,1,1], k = 4
    // Output: 4
    // Explanation: 
    // In this case, regardless which pile you choose, you have to leave behind 1 gift in each pile. 
    // That is, you can't take any pile with you. 
    // So, the total gifts remaining are 4.
    fmt.Println(pickGifts([]int{1,1,1,1}, 4)) // 4

    fmt.Println(pickGifts1([]int{25,64,9,4,100}, 4)) // 29
    fmt.Println(pickGifts1([]int{1,1,1,1}, 4)) // 4
}