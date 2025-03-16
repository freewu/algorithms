package main

// 3478. Choose K Elements With Maximum Sum
// You are given two integer arrays, nums1 and nums2, both of length n, along with a positive integer k.

// For each index i from 0 to n - 1, perform the following:
//     1. Find all indices j where nums1[j] is less than nums1[i].
//     2. Choose at most k values of nums2[j] at these indices to maximize the total sum.

// Return an array answer of size n, where answer[i] represents the result for the corresponding index i.

// Example 1:
// Input: nums1 = [4,2,1,5,3], nums2 = [10,20,30,40,50], k = 2
// Output: [80,30,0,80,50]
// Explanation:
// For i = 0: Select the 2 largest values from nums2 at indices [1, 2, 4] where nums1[j] < nums1[0], resulting in 50 + 30 = 80.
// For i = 1: Select the 2 largest values from nums2 at index [2] where nums1[j] < nums1[1], resulting in 30.
// For i = 2: No indices satisfy nums1[j] < nums1[2], resulting in 0.
// For i = 3: Select the 2 largest values from nums2 at indices [0, 1, 2, 4] where nums1[j] < nums1[3], resulting in 50 + 30 = 80.
// For i = 4: Select the 2 largest values from nums2 at indices [1, 2] where nums1[j] < nums1[4], resulting in 30 + 20 = 50.

// Example 2:
// Input: nums1 = [2,2,2,2], nums2 = [3,1,2,3], k = 1
// Output: [0,0,0,0]
// Explanation:
// Since all elements in nums1 are equal, no indices satisfy the condition nums1[j] < nums1[i] for any i, resulting in 0 for all positions.

// Constraints:
//     n == nums1.length == nums2.length
//     1 <= n <= 10^5
//     1 <= nums1[i], nums2[i] <= 10^6
//     1 <= k <= n

import "fmt"
import "container/heap"
import "sort"
import "slices"

type MinHeap []int
func (h MinHeap)  Len() int           { return len(h) }
func (h MinHeap)  Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap)  Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int))}
func (h *MinHeap) Pop() interface{}   {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[: n-1]
    return x
}

func findMaxSum(nums1 []int, nums2 []int, k int) []int64 {
    type Item struct { Index, V1, V2 int }
    n := len(nums1)
    items := make([]Item, n)
    for i := 0; i < n; i++ {
        items[i] = Item { Index: i, V1: nums1[i],  V2: nums2[i], }
    }
    sort.Slice(items, func(i, j int) bool {
        if items[i].V1 == items[j].V1 { return items[i].V2 > items[j].V2 }
        return items[i].V1 < items[j].V1
    })
    res, sum := make([]int64, n), int64(0)
    h := &MinHeap{}
    heap.Init(h)
    for i := 0; i < n; i++ {
        if i >= 1 && items[i].V1 == items[i-1].V1 {
            res[items[i].Index] = res[items[i - 1].Index]
        } else {
            res[items[i].Index] = sum
        }
        heap.Push(h, items[i].V2)
        sum += int64(items[i].V2)
        if len(*h) > k {
            item := heap.Pop(h).(int)
            sum -= int64(item)
        }
    }
    return res
}

type hp struct{ sort.IntSlice }
func (hp) Push(any)     {}
func (hp) Pop() (_ any) { return }

func findMaxSum1(nums1 []int, nums2 []int, k int) []int64 {
    n := len(nums1)
    type Tuple struct { x, y, i int }
    arr := make([]Tuple, n)
    for i, v := range nums1 {
        arr[i] = Tuple{ v, nums2[i], i}
    }
    slices.SortFunc(arr, func(a, b Tuple) int { 
        return a.x - b.x
    })
    res, sum := make([]int64, n), 0
    h := hp{make([]int, k)}
    for i, v := range arr {
        if i > 0 && v.x == arr[i-1].x {
            res[v.i] = res[arr[i-1].i]
        } else {
            res[v.i] = int64(sum)
        }
        y := v.y
        if i < k {
            sum += y
            h.IntSlice[i] = y
            continue
        }
        if i == k {
            heap.Init(&h)
        }
        if y > h.IntSlice[0] {
            sum += y - h.IntSlice[0]
            h.IntSlice[0] = y
            heap.Fix(&h, 0)
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums1 = [4,2,1,5,3], nums2 = [10,20,30,40,50], k = 2
    // Output: [80,30,0,80,50]
    // Explanation:
    // For i = 0: Select the 2 largest values from nums2 at indices [1, 2, 4] where nums1[j] < nums1[0], resulting in 50 + 30 = 80.
    // For i = 1: Select the 2 largest values from nums2 at index [2] where nums1[j] < nums1[1], resulting in 30.
    // For i = 2: No indices satisfy nums1[j] < nums1[2], resulting in 0.
    // For i = 3: Select the 2 largest values from nums2 at indices [0, 1, 2, 4] where nums1[j] < nums1[3], resulting in 50 + 30 = 80.
    // For i = 4: Select the 2 largest values from nums2 at indices [1, 2] where nums1[j] < nums1[4], resulting in 30 + 20 = 50.
    fmt.Println(findMaxSum([]int{4,2,1,5,3}, []int{10,20,30,40,50}, 2)) // [80,30,0,80,50]
    // Example 2:
    // Input: nums1 = [2,2,2,2], nums2 = [3,1,2,3], k = 1
    // Output: [0,0,0,0]
    // Explanation:
    // Since all elements in nums1 are equal, no indices satisfy the condition nums1[j] < nums1[i] for any i, resulting in 0 for all positions.
    fmt.Println(findMaxSum([]int{2,2,2,2}, []int{3,1,2,3}, 1)) // [0,0,0,0]

    fmt.Println(findMaxSum([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 2)) // [0 9 17 17 17 17 17 17 17]
    fmt.Println(findMaxSum([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 2)) // [0 1 3 5 7 9 11 13 15]
    fmt.Println(findMaxSum([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 2)) // [15 13 11 9 7 5 3 1 0]
    fmt.Println(findMaxSum([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 2)) // [17 17 17 17 17 17 17 9 0]

    fmt.Println(findMaxSum1([]int{4,2,1,5,3}, []int{10,20,30,40,50}, 2)) // [80,30,0,80,50]
    fmt.Println(findMaxSum1([]int{2,2,2,2}, []int{3,1,2,3}, 1)) // [0,0,0,0]
    fmt.Println(findMaxSum1([]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}, 2)) // [0 9 17 17 17 17 17 17 17]
    fmt.Println(findMaxSum1([]int{1,2,3,4,5,6,7,8,9}, []int{1,2,3,4,5,6,7,8,9}, 2)) // [0 1 3 5 7 9 11 13 15]
    fmt.Println(findMaxSum1([]int{9,8,7,6,5,4,3,2,1}, []int{9,8,7,6,5,4,3,2,1}, 2)) // [15 13 11 9 7 5 3 1 0]
    fmt.Println(findMaxSum1([]int{9,8,7,6,5,4,3,2,1}, []int{1,2,3,4,5,6,7,8,9}, 2)) // [17 17 17 17 17 17 17 9 0]
}