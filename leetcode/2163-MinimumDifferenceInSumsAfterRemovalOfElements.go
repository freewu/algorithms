package main

// 2163. Minimum Difference in Sums After Removal of Elements
// You are given a 0-indexed integer array nums consisting of 3 * n elements.

// You are allowed to remove any subsequence of elements of size exactly n from nums. 
// The remaining 2 * n elements will be divided into two equal parts:
//     1. The first n elements belonging to the first part and their sum is sumfirst.
//     2. The next n elements belonging to the second part and their sum is sumsecond.

// The difference in sums of the two parts is denoted as sumfirst - sumsecond.
//     1. For example, if sumfirst = 3 and sumsecond = 2, their difference is 1.
//     2. Similarly, if sumfirst = 2 and sumsecond = 3, their difference is -1.

// Return the minimum difference possible between the sums of the two parts after the removal of n elements.

// Example 1:
// Input: nums = [3,1,2]
// Output: -1
// Explanation: Here, nums has 3 elements, so n = 1. 
// Thus we have to remove 1 element from nums and divide the array into two equal parts.
// - If we remove nums[0] = 3, the array will be [1,2]. The difference in sums of the two parts will be 1 - 2 = -1.
// - If we remove nums[1] = 1, the array will be [3,2]. The difference in sums of the two parts will be 3 - 2 = 1.
// - If we remove nums[2] = 2, the array will be [3,1]. The difference in sums of the two parts will be 3 - 1 = 2.
// The minimum difference between sums of the two parts is min(-1,1,2) = -1. 

// Example 2:
// Input: nums = [7,9,5,8,1,3]
// Output: 1
// Explanation: Here n = 2. So we must remove 2 elements and divide the remaining array into two parts containing two elements each.
// If we remove nums[2] = 5 and nums[3] = 8, the resultant array will be [7,9,1,3]. The difference in sums will be (7+9) - (1+3) = 12.
// To obtain the minimum difference, we should remove nums[1] = 9 and nums[4] = 1. The resultant array becomes [7,5,8,3]. The difference in sums of the two parts is (7+5) - (8+3) = 1.
// It can be shown that it is not possible to obtain a difference smaller than 1.

// Constraints:
//     nums.length == 3 * n
//     1 <= n <= 10^5
//     1 <= nums[i] <= 10^5

import "fmt"
import "sort"
import "math"
import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() (v interface{}) {
    v = (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return v
}

type MaxHeap []int
func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() (v interface{}) {
    v = (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return v
}

func minimumDifference(nums []int) int64 {
    n, leftSum, rightSum:= len(nums) / 3, 0, 0
    mxh, mnh := MaxHeap(make([]int, 0, n)), MinHeap(make([]int, 0, n))
    for i := 0; i < n; i++ {
        leftSum += nums[i]
        mxh = append(mxh, nums[i])
        j := len(nums) - i - 1
        rightSum += nums[j]
        mnh = append(mnh, nums[j])
    }
    heap.Init(&mxh)
    heap.Init(&mnh)
    diff := make([]int, n + 1)
    for i := range diff {
        diff[i] += leftSum
        diff[len(diff) - i - 1] -= rightSum
        l, r := n + i, len(nums) - n - i - 1
        heap.Push(&mxh, nums[l])
        heap.Push(&mnh, nums[r])
        leftSum += nums[l]
        leftSum -= heap.Pop(&mxh).(int)
        rightSum += nums[r]
        rightSum -= heap.Pop(&mnh).(int)
    }
    res := math.MaxInt64
    for _, d := range diff {
        if d < res {
            res = d
        }
    }
    return int64(res)
}

type MinHeap1 struct { sort.IntSlice }

func (h MinHeap1) Less(i, j int) bool { return h.IntSlice[i] < h.IntSlice[j] }
func (h *MinHeap1) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *MinHeap1) Pop() any {
    a := h.IntSlice
    v := a[len(a)-1]
    h.IntSlice = a[:len(a)-1]
    return v
}

func minimumDifference1(nums []int) int64 {
    m, sum := len(nums), 0
    n := m / 3
    prefix := make([]int, m + 1)
    pq1 := MinHeap1{}
    for i := 1; i <= n * 2; i++ {
        sum += nums[i-1]
        heap.Push(&pq1, -nums[i-1])
        if pq1.Len() > n {
            sum -= -heap.Pop(&pq1).(int)
        }
        prefix[i] = sum
    }
    sum = 0
    suffix:= make([]int, m + 1)
    pq2 := MinHeap1{}
    for i := m; i > n; i-- {
        sum += nums[i-1]
        heap.Push(&pq2, nums[i-1])
        if pq2.Len() > n {
            sum -= heap.Pop(&pq2).(int)
        }
        suffix[i] = sum
    }
    res := math.MaxInt64
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := n; i <= n * 2; i++ {
        res = min(res, prefix[i] - suffix[i+1])
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [3,1,2]
    // Output: -1
    // Explanation: Here, nums has 3 elements, so n = 1. 
    // Thus we have to remove 1 element from nums and divide the array into two equal parts.
    // - If we remove nums[0] = 3, the array will be [1,2]. The difference in sums of the two parts will be 1 - 2 = -1.
    // - If we remove nums[1] = 1, the array will be [3,2]. The difference in sums of the two parts will be 3 - 2 = 1.
    // - If we remove nums[2] = 2, the array will be [3,1]. The difference in sums of the two parts will be 3 - 1 = 2.
    // The minimum difference between sums of the two parts is min(-1,1,2) = -1. 
    fmt.Println(minimumDifference([]int{3,1,2})) // -1
    // Example 2:
    // Input: nums = [7,9,5,8,1,3]
    // Output: 1
    // Explanation: Here n = 2. So we must remove 2 elements and divide the remaining array into two parts containing two elements each.
    // If we remove nums[2] = 5 and nums[3] = 8, the resultant array will be [7,9,1,3]. The difference in sums will be (7+9) - (1+3) = 12.
    // To obtain the minimum difference, we should remove nums[1] = 9 and nums[4] = 1. The resultant array becomes [7,5,8,3]. The difference in sums of the two parts is (7+5) - (8+3) = 1.
    // It can be shown that it is not possible to obtain a difference smaller than 1.
    fmt.Println(minimumDifference([]int{7,9,5,8,1,3})) // 1

    fmt.Println(minimumDifference1([]int{3,1,2})) // -1
    fmt.Println(minimumDifference1([]int{7,9,5,8,1,3})) // 1
}