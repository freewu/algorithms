package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
)

/**
215. Kth Largest Element in an Array
Given an integer array nums and an integer k, return the kth largest element in the array.
Note that it is the kth largest element in the sorted order, not the kth distinct element.


Example 1:

	Input: nums = [3,2,1,5,6,4], k = 2
	Output: 5

Example 2:

	Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
	Output: 4

Constraints:

	1 <= k <= nums.length <= 10^4
	-10^4 <= nums[i] <= 10^4
*/

// 解法一 排序，排序的方法反而速度是最快的
func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums) // 排序
	return nums[len(nums)-k] // 取 k 位
}

// 解法二 这个方法的理论依据是 partition 得到的点的下标就是最终排序之后的下标，根据这个下标，我们可以判断第 K 大的数在哪里
// 时间复杂度 O(n)，空间复杂度 O(log n)，最坏时间复杂度为 O(n^2)，空间复杂度 O(n)
func findKthLargest(nums []int, k int) int {
	m := len(nums) - k + 1 // mth smallest, from 1..len(nums)
	return selectSmallest(nums, 0, len(nums)-1, m)
}

func selectSmallest(nums []int, l, r, i int) int {
	if l >= r {
		return nums[l]
	}
	q := partition(nums, l, r)
	k := q - l + 1
	if k == i {
		return nums[q]
	}
	if i < k {
		return selectSmallest(nums, l, q-1, i)
	} else {
		return selectSmallest(nums, q+1, r, i-k)
	}
}

func partition(nums []int, l, r int) int {
	k := l + rand.Intn(r-l+1) // 此处为优化，使得时间复杂度期望降为 O(n)，最坏时间复杂度为 O(n^2)
	nums[k], nums[r] = nums[r], nums[k]
	i := l - 1
	// nums[l..i] <= nums[r]
	// nums[i+1..j-1] > nums[r]
	for j := l; j < r; j++ {
		if nums[j] <= nums[r] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

// 扩展题 剑指 Offer 40. 最小的 k 个数
func getLeastNumbers(arr []int, k int) []int {
	return selectSmallest1(arr, 0, len(arr)-1, k)[:k]
}

// 和 selectSmallest 实现完全一致，只是返回值不用再截取了，直接返回 nums 即可
func selectSmallest1(nums []int, l, r, i int) []int {
	if l >= r {
		return nums
	}
	q := partition(nums, l, r)
	k := q - l + 1
	if k == i {
		return nums
	}
	if i < k {
		return selectSmallest1(nums, l, q-1, i)
	} else {
		return selectSmallest1(nums, q+1, r, i-k)
	}
}

// best solution
func findKthLargestBest(nums []int, k int) int {
	fmt.Printf("nums = %v\n",nums)
	fmt.Printf("nums[:k] = %v\n",nums[:k])
	h := minHeap(nums[:k])
	heap.Init(&h)
	fmt.Printf("nums[k:] = %v\n",nums[k:])
	for _, num := range nums[k:] {
		fmt.Printf("h[0] = %v,num = %v\n",h[0],num)
		if h[0] < num {
			h[0] = num
			heap.Fix(&h, 0)
		}
		fmt.Printf("%v\n",h)
	}
	return heap.Pop(&h).(int)
}

type minHeap []int
func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	var x int
	x, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return x
}

func main() {
	fmt.Printf("findKthLargest1([]int{ 3,2,1,5,6,4}, 2) = %v\n",findKthLargest1([]int{ 3,2,1,5,6,4} ,2)) // 5
	fmt.Printf("findKthLargest1([]int{ 3,2,3,1,2,4,5,5,6}, 4) = %v\n",findKthLargest1([]int{ 3,2,3,1,2,4,5,5,6} ,4)) // 4

	fmt.Printf("findKthLargest([]int{ 3,2,1,5,6,4}, 2) = %v\n",findKthLargest([]int{ 3,2,1,5,6,4} ,2)) // 5
	fmt.Printf("findKthLargest([]int{ 3,2,3,1,2,4,5,5,6}, 4) = %v\n",findKthLargest([]int{ 3,2,3,1,2,4,5,5,6} ,4)) // 4

	fmt.Printf("findKthLargestBest([]int{ 3,2,1,5,6,4}, 2) = %v\n",findKthLargestBest([]int{ 3,2,1,5,6,4} ,2)) // 5
	fmt.Printf("findKthLargestBest([]int{ 3,2,3,1,2,4,5,5,6}, 4) = %v\n",findKthLargestBest([]int{ 3,2,3,1,2,4,5,5,6} ,4)) // 4
}
