package main

// LCR 076. 数组中的第 K 个最大元素
// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

// 示例 1:
// 输入: [3,2,1,5,6,4] 和 k = 2
// 输出: 5

// 示例 2:
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
// 输出: 4

// 提示：
//     1 <= k <= nums.length <= 10^4
//     -10^4 <= nums[i] <= 10^4

import "fmt"
import "container/heap"
import "math/rand"
import "sort"

// 排序，排序的方法反而速度是最快的
func findKthLargest1(nums []int, k int) int {
    sort.Ints(nums) // 排序
    return nums[len(nums)-k] // 取 k 位
}

// 解法二 这个方法的理论依据是 partition 得到的点的下标就是最终排序之后的下标，根据这个下标，我们可以判断第 K 大的数在哪里
// 时间复杂度 O(n)，空间复杂度 O(log n)，最坏时间复杂度为 O(n^2)，空间复杂度 O(n)
func findKthLargest(nums []int, k int) int {
    m := len(nums) - k + 1 // mth smallest, from 1..len(nums)
    var selectSmallest func (nums []int, l, r, i int) int
    selectSmallest = func (nums []int, l, r, i int) int {
        if l >= r {
            return nums[l]
        }
        partition := func (nums []int, l, r int) int {
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
    return selectSmallest(nums, 0, len(nums)-1, m)
}

// 扩展题 剑指 Offer 40. 最小的 k 个数
func getLeastNumbers(arr []int, k int) []int {
    // 和 selectSmallest 实现完全一致，只是返回值不用再截取了，直接返回 nums 即可
    var selectSmallest func (nums []int, l, r, i int) []int
    selectSmallest = func (nums []int, l, r, i int) []int {
        if l >= r {
            return nums
        }
        partition := func (nums []int, l, r int) int {
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
        q := partition(nums, l, r)
        k := q - l + 1
        if k == i {
            return nums
        }
        if i < k {
            return selectSmallest(nums, l, q-1, i)
        } else {
            return selectSmallest(nums, q+1, r, i-k)
        }
    }
	return selectSmallest(arr, 0, len(arr)-1, k)[:k]
}

// min heap
func findKthLargest2(nums []int, k int) int {
    h := minHeap(nums[:k])
    heap.Init(&h)
    for _, num := range nums[k:] {
        if h[0] < num {
            h[0] = num
            heap.Fix(&h, 0)
        }
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

    fmt.Printf("findKthLargest2([]int{ 3,2,1,5,6,4}, 2) = %v\n",findKthLargest2([]int{ 3,2,1,5,6,4} ,2)) // 5
    fmt.Printf("findKthLargest2([]int{ 3,2,3,1,2,4,5,5,6}, 4) = %v\n",findKthLargest2([]int{ 3,2,3,1,2,4,5,5,6} ,4)) // 4
}
