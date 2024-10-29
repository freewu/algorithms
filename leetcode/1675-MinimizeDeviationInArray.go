package main

// 1675. Minimize Deviation in Array
// You are given an array nums of n positive integers.

// You can perform two types of operations on any element of the array any number of times:
//     If the element is even, divide it by 2.
//         For example, if the array is [1,2,3,4], 
//         then you can do this operation on the last element, and the array will be [1,2,3,2].
//     If the element is odd, multiply it by 2.
//         For example, if the array is [1,2,3,4], 
//         then you can do this operation on the first element, and the array will be [2,2,3,4].

// The deviation of the array is the maximum difference between any two elements in the array.

// Return the minimum deviation the array can have after performing some number of operations.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: 1
// Explanation: You can transform the array to [1,2,3,2], then to [2,2,3,2], then the deviation will be 3 - 2 = 1.

// Example 2:
// Input: nums = [4,1,5,20,3]
// Output: 3
// Explanation: You can transform the array after two operations to [4,2,5,5,3], then the deviation will be 5 - 2 = 3.

// Example 3:
// Input: nums = [2,10,8]
// Output: 3

// Constraints:
//     n == nums.length
//     2 <= n <= 5 * 10^4
//     1 <= nums[i] <= 10^9

import "fmt"
import "container/heap"

type MaxHeap []int
func (h MaxHeap)  Len()  int { return len(h)}
func (h MaxHeap)  Swap(i int, j int) { h[i], h[j] = h[j], h[i]}
func (h MaxHeap)  Less(i int, j int) bool { return h[i] > h[j]}
func (h *MaxHeap) Push(a interface{}) {*h = append(*h, a.(int))}
func (h *MaxHeap) Pop()  interface{} { 
    n := len(*h)
    res := (*h)[n-1]
    *h = (*h)[:n-1]
    return res
}

func minimumDeviation(nums []int) int {
    mx, mn, diff := 0, 1 << 31, 1 << 31
    mh := &MaxHeap{}
    heap.Init(mh)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, v := range nums {
        if v % 2 == 1 {
            v *= 2
        }
        mn = min(mn, v)
        heap.Push(mh, v)
    }
    for {
        mx = heap.Pop(mh).(int)
        if mx % 2 == 1 {
            heap.Push(mh, mx)
            break
        }
        diff = min(diff, mx - mn)
        mx = mx / 2
        mn = min(mn, mx)
        heap.Push(mh, mx)
    }
    return min(diff, mx - mn)
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3,4]
    // Output: 1
    // Explanation: You can transform the array to [1,2,3,2], then to [2,2,3,2], then the deviation will be 3 - 2 = 1.
    fmt.Println(minimumDeviation([]int{1,2,3,4})) // 1
    // Example 2:
    // Input: nums = [4,1,5,20,3]
    // Output: 3
    // Explanation: You can transform the array after two operations to [4,2,5,5,3], then the deviation will be 5 - 2 = 3.
    fmt.Println(minimumDeviation([]int{4,1,5,20,3})) // 3
    // Example 3:
    // Input: nums = [2,10,8]
    // Output: 3
    fmt.Println(minimumDeviation([]int{2,10,8})) // 3
}