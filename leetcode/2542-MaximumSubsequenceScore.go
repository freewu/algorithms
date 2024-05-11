package main

// 2542. Maximum Subsequence Score
// You are given two 0-indexed integer arrays nums1 and nums2 of equal length n and a positive integer k. 
// You must choose a subsequence of indices from nums1 of length k.
// For chosen indices i0, i1, ..., ik - 1, your score is defined as:
//     The sum of the selected elements from nums1 multiplied with the minimum of the selected elements from nums2.
//     It can defined simply as: (nums1[i0] + nums1[i1] +...+ nums1[ik - 1]) * min(nums2[i0] , nums2[i1], ... ,nums2[ik - 1]).

// Return the maximum possible score.
// A subsequence of indices of an array is a set that can be derived from the set {0, 1, ..., n-1} by deleting some or no elements.

// Example 1:
// Input: nums1 = [1,3,3,2], nums2 = [2,1,3,4], k = 3
// Output: 12
// Explanation: 
// The four possible subsequence scores are:
// - We choose the indices 0, 1, and 2 with score = (1+3+3) * min(2,1,3) = 7.
// - We choose the indices 0, 1, and 3 with score = (1+3+2) * min(2,1,4) = 6. 
// - We choose the indices 0, 2, and 3 with score = (1+3+2) * min(2,3,4) = 12. 
// - We choose the indices 1, 2, and 3 with score = (3+3+2) * min(1,3,4) = 8.
// Therefore, we return the max score, which is 12.

// Example 2:
// Input: nums1 = [4,2,3,1,1], nums2 = [7,5,10,9,6], k = 1
// Output: 30
// Explanation: 
// Choosing index 2 is optimal: nums1[2] * nums2[2] = 3 * 10 = 30 is the maximum possible score.
 
// Constraints:
//     n == nums1.length == nums2.length
//     1 <= n <= 10^5
//     0 <= nums1[i], nums2[j] <= 10^5
//     1 <= k <= n

import "fmt"
import "sort"
import "container/heap"
import "slices"

// Define the pair sorter
type PairSorter struct {
    nums1, nums2 []int
}
// Define the methods required by the Sort interface
func (p PairSorter) Len() int {return len(p.nums1)}
func (p PairSorter) Less(i, j int) bool {return p.nums2[i] > p.nums2[j]}
func (p PairSorter) Swap(i, j int) {
    p.nums1[i], p.nums1[j] = p.nums1[j], p.nums1[i]
    p.nums2[i], p.nums2[j] = p.nums2[j], p.nums2[i]
}

// Define the min-heap for integer
type IntHeap []int
// Define the methods required by the heap interface
func (h IntHeap) Len() int {return len(h)}
func (h IntHeap) Less(i, j int) bool {return h[i] < h[j]}
func (h IntHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *IntHeap) Push(x any) {
    *h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
    n := len(*h)
    x := (*h)[n - 1]
    *h = (*h)[:n - 1]
    return x
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
    sort.Sort(PairSorter{nums1, nums2}) // Sort the array pair
    h := &IntHeap{} // Create the min-heap with the first k elements of nums1
    curSum := 0
    for i := 0; i < k; i++ {
        heap.Push(h, nums1[i])
        curSum = curSum + nums1[i]
    }
    maxScore := curSum * nums2[k - 1]
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for j := k; j < len(nums2); j++ { // Traverse the nums1 array to find the max score
        curSum = curSum - heap.Pop(h).(int)
        curSum = curSum + nums1[j]
        heap.Push(h, nums1[j])
        maxScore = max(maxScore, curSum * nums2[j])
    }
    return int64(maxScore)
}

func maxScore1(nums1, nums2 []int, k int) int64 {
    ids := make([]int, len(nums1))
    for i := range ids {
        ids[i] = i
    }
    slices.SortFunc(ids, func(i, j int) int { return nums2[j] - nums2[i] }) // 对下标排序，不影响原数组的顺序
    max := func (x, y int) int { if x > y { return x; }; return y; }
    h := MinHeap{make([]int, k)}
    sum := 0
    for i, idx := range ids[:k] {
        sum += nums1[idx]
        h.IntSlice[i] = nums1[idx]
    }
    heap.Init(&h)
    res := sum * nums2[ids[k-1]]
    for _, i := range ids[k:] {
        x := nums1[i]
        if x > h.IntSlice[0] {
            sum += x - h.replace(x)
            res = max(res, sum * nums2[i])
        }
    }
    return int64(res)
}

type MinHeap struct{ sort.IntSlice }
func (MinHeap) Push(any)            {}
func (MinHeap) Pop() (_ any)        { return }
func (h MinHeap) replace(v int) int { top := h.IntSlice[0]; h.IntSlice[0] = v; heap.Fix(&h, 0); return top }

func main() {
    // Example 1:
    // Input: nums1 = [1,3,3,2], nums2 = [2,1,3,4], k = 3
    // Output: 12
    // Explanation: 
    // The four possible subsequence scores are:
    // - We choose the indices 0, 1, and 2 with score = (1+3+3) * min(2,1,3) = 7.
    // - We choose the indices 0, 1, and 3 with score = (1+3+2) * min(2,1,4) = 6. 
    // - We choose the indices 0, 2, and 3 with score = (1+3+2) * min(2,3,4) = 12. 
    // - We choose the indices 1, 2, and 3 with score = (3+3+2) * min(1,3,4) = 8.
    // Therefore, we return the max score, which is 12.
    fmt.Println(maxScore([]int{1,3,3,2},[]int{2,1,3,4}, 3)) // 12
    // Example 2:
    // Input: nums1 = [4,2,3,1,1], nums2 = [7,5,10,9,6], k = 1
    // Output: 30
    // Explanation: 
    // Choosing index 2 is optimal: nums1[2] * nums2[2] = 3 * 10 = 30 is the maximum possible score.
    fmt.Println(maxScore([]int{4,2,3,1,1},[]int{7,5,10,9,6}, 1)) // 30

    fmt.Println(maxScore1([]int{1,3,3,2},[]int{2,1,3,4}, 3)) // 12
    fmt.Println(maxScore1([]int{4,2,3,1,1},[]int{7,5,10,9,6}, 1)) // 30
}