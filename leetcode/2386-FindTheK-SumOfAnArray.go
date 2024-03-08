package main

// 2386. Find the K-Sum of an Array
// You are given an integer array nums and a positive integer k. You can choose any subsequence of the array and sum all of its elements together.
// We define the K-Sum of the array as the kth largest subsequence sum that can be obtained (not necessarily distinct).
// Return the K-Sum of the array.
// A subsequence is an array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.
// Note that the empty subsequence is considered to have a sum of 0.

// Example 1:
// Input: nums = [2,4,-2], k = 5
// Output: 2
// Explanation: All the possible subsequence sums that we can obtain are the following sorted in decreasing order:
// - 6, 4, 4, 2, 2, 0, 0, -2.
// The 5-Sum of the array is 2.

// Example 2:
// Input: nums = [1,-2,3,4,-10,12], k = 16
// Output: 10
// Explanation: The 16-Sum of the array is 10.

// Constraints:
//         n == nums.length
//         1 <= n <= 10^5
//         -10^9 <= nums[i] <= 10^9
//         1 <= k <= min(2000, 2^n)

import "fmt"
import "sort"
import "container/heap"

type KSumItem struct {
	Sum   int64
	Index int
}

type PQ []KSumItem

func (m PQ) Len() int { return len(m) }
func (m PQ) Less(i, j int) bool {
	return m[i].Sum > m[j].Sum
}
func (m PQ) Swap(i, j int)       { m[i], m[j] = m[j], m[i] }
func (m *PQ) Push(x interface{}) { *m = append(*m, x.(KSumItem)) }
func (m *PQ) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

// Priority Queue
func kSum(nums []int, k int) int64 {
	maxSum := int64(0)
	for i := range nums {
		if nums[i] >= 0 {
			maxSum += int64(nums[i])
		} else {
			nums[i] = -nums[i]
		}
	}
	sort.Ints(nums)

	maxHeap := make(PQ, 0)
	var nextSum int64
	heap.Push(&maxHeap, KSumItem{maxSum - int64(nums[0]), 0})
	nextSum = maxSum
	for ; k > 1; k-- {
		num := heap.Pop(&maxHeap).(KSumItem)
		nextSum = num.Sum
		i := num.Index
		if i+1 < len(nums) {
			heap.Push(&maxHeap, KSumItem{nextSum + int64(nums[i]) - int64(nums[i+1]), i + 1})
			heap.Push(&maxHeap, KSumItem{nextSum - int64(nums[i+1]), i + 1})
		}
	}
	return nextSum
}

// best solution
type Pair struct {
    sum int
    idx int
}

type mheap []Pair

func (h mheap) Len() int {
	return len(h)
}

func (h mheap) Less(i, j int) bool {
	return h[i].sum < h[j].sum
}

func (h mheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *mheap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *mheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSum1(nums []int, k int) int64 {
    maxSum := 0
    n := len(nums)

    for i := 0; i < n; i++ {
        if nums[i] < 0 {
            nums[i] = -nums[i]
        } else {
            maxSum += nums[i]
        }
    }

    sort.Ints(nums)
    topk := make([]int, 0)
    topk = append(topk, 0)

    h := &mheap{}
	heap.Init(h)
    heap.Push(h, Pair{sum: nums[0], idx: 1})

    for len(topk) < k {
        pair := heap.Pop(h).(Pair)

        topk = append(topk, pair.sum)
        if pair.idx < n {
            heap.Push(h, Pair{sum: pair.sum+nums[pair.idx], idx: pair.idx+1})
            heap.Push(h, Pair{sum: pair.sum-nums[pair.idx-1]+nums[pair.idx], idx: pair.idx+1})
        }
    }
    return int64(maxSum - topk[k-1])
}

func main() {
    // All the possible subsequence sums that we can obtain are the following sorted in decreasing order:
    //      - 6, 4, 4, 2, 2, 0, 0, -2.
    // The 5-Sum of the array is 2.
    fmt.Println(kSum([]int{2,4,-2},5)) // 2
    // The 16-Sum of the array is 10.
    fmt.Println(kSum([]int{1,-2,3,4,-10,12},16)) // 10

    fmt.Println(kSum1([]int{2,4,-2},5)) // 2
    fmt.Println(kSum1([]int{1,-2,3,4,-10,12},16)) // 10
}