package main

// 786. K-th Smallest Prime Fraction
// You are given a sorted integer array arr containing 1 and prime numbers, 
// where all the integers of arr are unique. You are also given an integer k.

// For every i and j where 0 <= i < j < arr.length, we consider the fraction arr[i] / arr[j].

// Return the kth smallest fraction considered. Return your answer as an array of integers of size 2, 
// where answer[0] == arr[i] and answer[1] == arr[j].

// Example 1:
// Input: arr = [1,2,3,5], k = 3
// Output: [2,5]
// Explanation: The fractions to be considered in sorted order are:
// 1/5, 1/3, 2/5, 1/2, 3/5, and 2/3.
// The third fraction is 2/5.

// Example 2:
// Input: arr = [1,7], k = 1
// Output: [1,7]

// Constraints:
//     2 <= arr.length <= 1000
//     1 <= arr[i] <= 3 * 10^4
//     arr[0] == 1
//     arr[i] is a prime number for i > 0.
//     All the numbers of arr are unique and sorted in strictly increasing order.
//     1 <= k <= arr.length * (arr.length - 1) / 2
 
// Follow up: Can you solve the problem with better than O(n2) complexity?

import "fmt"
import "sort"
import "container/heap"

// 暴力解法
func kthSmallestPrimeFraction(arr []int, k int) []int {
    type pair struct{
        x int
        y int
    }
    var all []pair
    for i, x := range arr {
        for _, y := range arr[i+1:] {
            all = append(all, pair{x, y})
        }
    }
    sort.Slice(all, func(i, j int) bool {
        a, b := all[i], all[j]
        return a.x*b.y < a.y*b.x
    })
    return []int{all[k-1].x, all[k-1].y}
}

// MinHeap
type frac struct{ x, y, i, j int }
type MinHeap []frac
func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].x*h[j].y < h[i].y*h[j].x }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(v interface{}) { *h = append(*h, v.(frac)) }
func (h *MinHeap) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func kthSmallestPrimeFraction1(arr []int, k int) []int {
    n := len(arr)
    h := make(MinHeap, n - 1)
    for j := 1; j < n; j++ {
        h[j-1] = frac{arr[0], arr[j], 0, j}
    }
    heap.Init(&h)
    for loop := k - 1; loop > 0; loop-- {
        f := heap.Pop(&h).(frac)
        if f.i+1 < f.j {
            heap.Push(&h, frac{arr[f.i+1], f.y, f.i + 1, f.j})
        }
    }
    return []int{h[0].x, h[0].y}
}

func kthSmallestPrimeFraction2(arr []int, k int) []int {
    n := len(arr)
    left, right := 0., 1.
    for {
        mid := (left + right) / 2
        i, count := -1, 0
        x, y := 0, 1 // 记录最大的分数
        for j := 1; j < n; j++ {
            for float64(arr[i+1]) / float64(arr[j]) < mid {
                i++
                if arr[i]*y > arr[j]*x {
                    x, y = arr[i], arr[j]
                }
            }
            count += i + 1
        }
        if count == k {
            return []int{x, y}
        }
        if count < k {
            left = mid
        } else {
            right = mid
        }
    }
}

func main() {
    // Example 1:
    // Input: arr = [1,2,3,5], k = 3
    // Output: [2,5]
    // Explanation: The fractions to be considered in sorted order are:
    // 1/5, 1/3, 2/5, 1/2, 3/5, and 2/3.
    // The third fraction is 2/5.
    fmt.Println(kthSmallestPrimeFraction([]int{1,2,3,5}, 3)) // [2,5]
    // Example 2:
    // Input: arr = [1,7], k = 1
    // Output: [1,7]
    fmt.Println(kthSmallestPrimeFraction([]int{1,7}, 1)) // [1,7]

    fmt.Println(kthSmallestPrimeFraction1([]int{1,2,3,5}, 3)) // [2,5]
    fmt.Println(kthSmallestPrimeFraction1([]int{1,7}, 1)) // [1,7]

    fmt.Println(kthSmallestPrimeFraction2([]int{1,2,3,5}, 3)) // [2,5]
    fmt.Println(kthSmallestPrimeFraction2([]int{1,7}, 1)) // [1,7]
}