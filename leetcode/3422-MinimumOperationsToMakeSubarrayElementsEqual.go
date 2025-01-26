package main

// 3422. Minimum Operations to Make Subarray Elements Equal
// You are given an integer array nums and an integer k. 
// You can perform the following operation any number of times:
//     Increase or decrease any element of nums by 1.

// Return the minimum number of operations required to ensure that at least one subarray of size k in nums has all elements equal.

// Example 1:
// Input: nums = [4,-3,2,1,-4,6], k = 3
// Output: 5
// Explanation:
// Use 4 operations to add 4 to nums[1]. The resulting array is [4, 1, 2, 1, -4, 6].
// Use 1 operation to subtract 1 from nums[2]. The resulting array is [4, 1, 1, 1, -4, 6].
// The array now contains a subarray [1, 1, 1] of size k = 3 with all elements equal. Hence, the answer is 5.

// Example 2:
// Input: nums = [-2,-2,3,1,4], k = 2
// Output: 0
// Explanation:
// The subarray [-2, -2] of size k = 2 already contains all equal elements, so no operations are needed. Hence, the answer is 0.

// Constraints:
//     2 <= nums.length <= 10^5
//     -10^6 <= nums[i] <= 10^6
//     2 <= k <= nums.length

import "fmt"
import "sort"
import "container/heap"
import "math"

// 超出时间限制 763 / 773
func minOperations1(nums []int, k int) int64 {
    n := len(nums)
    if n < k {
        return 0
    }
    // 初始化窗口
    sorted := make([]int, k)
    copy(sorted, nums[:k])
    sort.Ints(sorted)
    prefix := make([]int, k+1)
    for i := 0; i < k; i++ {
        prefix[i+1] = prefix[i] + sorted[i]
    }
    computeMinOps := func(sorted []int, prefix []int) int {
        k := len(sorted)
        if k == 0 { return 0 }
        if k%2 == 1 {
            mid := k / 2
            m := sorted[mid]
            leftSum := prefix[mid]
            rightSum := prefix[k] - prefix[mid+1]
            return m*mid - leftSum + (rightSum - m*(k-mid-1))
        } else {
            mid1 := k/2 - 1
            mid2 := k / 2
            m1 := sorted[mid1]
            leftSum1 := prefix[mid1]
            rightSum1 := prefix[k] - prefix[mid1+1]
            ops1 := m1*mid1 - leftSum1 + (rightSum1 - m1*(k-mid1-1))
            m2 := sorted[mid2]
            leftSum2 := prefix[mid2]
            rightSum2 := prefix[k] - prefix[mid2+1]
            ops2 := m2*mid2 - leftSum2 + (rightSum2 - m2*(k-mid2-1))
            if ops1 < ops2 {
                return ops1
            }
            return ops2
        }
    }
    res := computeMinOps(sorted, prefix)
    for i := k; i < n; i++ {
        // 移除左边的元素
        leftVal := nums[i-k]
        pos := sort.SearchInts(sorted, leftVal)
        if pos < len(sorted) && sorted[pos] == leftVal {
            sorted = append(sorted[:pos], sorted[pos+1:]...)
        }
        // 插入右边的元素
        rightVal := nums[i]
        insertPos := sort.SearchInts(sorted, rightVal)
        sorted = append(sorted[:insertPos], append([]int{rightVal}, sorted[insertPos:]...)...)
        // 重新计算prefix
        newK := len(sorted)
        prefix = make([]int, newK+1)
        for j := 0; j < newK; j++ {
            prefix[j+1] = prefix[j] + sorted[j]
        }
        // 计算当前的操作次数
        currentOps := computeMinOps(sorted, prefix)
        if currentOps < res {
            res = currentOps
        }
    }
    return int64(res)
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int))}
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type MedianMaintainer struct {
    lower      *MaxHeap
    higher     *MinHeap
    lowerSize  int
    higherSize int
    lowerSum   int64
    higherSum  int64
    removes    map[int]int
}

func (m *MedianMaintainer) addNum(num int) {
    if m.lower.Len() == 0 || num <= (*m.lower)[0] {
        heap.Push(m.lower, num)
        m.lowerSize++
        m.lowerSum += int64(num)
    } else {
        heap.Push(m.higher, num)
        m.higherSize++
        m.higherSum += int64(num)
    }
    m.adjustPriorityQueues()
}

func (m *MedianMaintainer) removeNum(num int) {
    m.removes[num]++
    if m.lower.Len() > 0 && num <= (*m.lower)[0] {
        m.lowerSize--
        m.lowerSum -= int64(num)
    } else {
        m.higherSize--
        m.higherSum -= int64(num)
    }
    m.adjustPriorityQueues()
}

func (m *MedianMaintainer) adjustPriorityQueues() {
    for m.lowerSize > m.higherSize+1 {
        num := heap.Pop(m.lower).(int)
        m.lowerSize--
        m.lowerSum -= int64(num)
        heap.Push(m.higher, num)
        m.higherSize++
        m.higherSum += int64(num)
    }
    for m.higherSize > m.lowerSize {
        num := heap.Pop(m.higher).(int)
        m.higherSize--
        m.higherSum -= int64(num)
        heap.Push(m.lower, num)
        m.lowerSize++
        m.lowerSum += int64(num)
    }
    for m.lower.Len() > 0 {
        top := (*m.lower)[0]
        if cnt, ok := m.removes[top]; ok && cnt > 0 {
            heap.Pop(m.lower)
            m.removes[top]--
            if m.removes[top] == 0 {
                delete(m.removes, top)
            }
        } else {
            break
        }
    }
    for m.higher.Len() > 0 {
        top := (*m.higher)[0]
        if cnt, ok := m.removes[top]; ok && cnt > 0 {
            heap.Pop(m.higher)
            m.removes[top]--
            if m.removes[top] == 0 {
                delete(m.removes, top)
            }
        } else {
            break
        }
    }
}

func (m *MedianMaintainer) findMedian() int {
    return (*m.lower)[0]
}

func minOperations(nums []int, k int) int64 {
    mm := &MedianMaintainer{
        lower:   &MaxHeap{},
        higher:  &MinHeap{},
        removes: make(map[int]int),
    }
    heap.Init(mm.lower)
    heap.Init(mm.higher)
    res, n := int64(math.MaxInt64), len(nums)
    for i := 0; i < n; i++ {
        if i >= k {
            mm.removeNum(nums[i-k])
        }
        mm.addNum(nums[i])
        if i >= k-1 {
            median := mm.findMedian()
            operations := mm.higherSum - mm.lowerSum + int64(median)*int64(mm.lowerSize-mm.higherSize)
            if operations < res {
                res = operations
            }
        }
    }
    if res == math.MaxInt64 { return 0 }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [4,-3,2,1,-4,6], k = 3
    // Output: 5
    // Explanation:
    // Use 4 operations to add 4 to nums[1]. The resulting array is [4, 1, 2, 1, -4, 6].
    // Use 1 operation to subtract 1 from nums[2]. The resulting array is [4, 1, 1, 1, -4, 6].
    // The array now contains a subarray [1, 1, 1] of size k = 3 with all elements equal. Hence, the answer is 5.
    fmt.Println(minOperations([]int{4,-3,2,1,-4,6}, 3)) // 5
    // Example 2:
    // Input: nums = [-2,-2,3,1,4], k = 2
    // Output: 0
    // Explanation:
    // The subarray [-2, -2] of size k = 2 already contains all equal elements, so no operations are needed. Hence, the answer is 0.
    fmt.Println(minOperations([]int{-2,-2,3,1,4}, 2)) // 0

    fmt.Println(minOperations([]int{1,2,3,4,5,6,7,8,9}, 2)) // 1
    fmt.Println(minOperations([]int{9,8,7,6,5,4,3,2,1}, 2)) // 1

    fmt.Println(minOperations1([]int{4,-3,2,1,-4,6}, 3)) // 5
    fmt.Println(minOperations1([]int{-2,-2,3,1,4}, 2)) // 0
    fmt.Println(minOperations1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 1
    fmt.Println(minOperations1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 1
}