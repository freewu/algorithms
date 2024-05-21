package main

// 327. Count of Range Sum
// Given an integer array nums and two integers lower and upper, 
// return the number of range sums that lie in [lower, upper] inclusive.

// Range sum S(i, j) is defined as the sum of the elements in nums between indices i and j inclusive, where i <= j.

// Example 1:
// Input: nums = [-2,5,-1], lower = -2, upper = 2
// Output: 3
// Explanation: The three ranges are: [0,0], [2,2], and [0,2] and their respective sums are: -2, -1, 2.

// Example 2:
// Input: nums = [0], lower = 0, upper = 0
// Output: 1

// Constraints:
//     1 <= nums.length <= 10^5
//     -2^31 <= nums[i] <= 2^31 - 1
//     -10^5 <= lower <= upper <= 10^5
//     The answer is guaranteed to fit in a 32-bit integer.

import "fmt"
import "sort"
import "math/bits"

// 暴力 超出时间限制 62 / 67 
func countRangeSum1(nums []int, lower int, upper int) int {
    count, prefix := 0, make([]int, len(nums)+1)
    for i := 0; i < len(nums); i++ {
        prefix[i+1] = prefix[i] + nums[i]
    }
    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums)+1; j++ {
            if sum := prefix[j] - prefix[i]; lower <= sum && sum <= upper {
                count++
            }
        }
    }
    return count
}

// 线段树
func countRangeSum(nums []int, lower, upper int) int {
    prefix := make([]int, len(nums)+1)
    for i := 0; i < len(nums); i++ {
        prefix[i+1] = prefix[i] + nums[i]
    }
    sorted := make([]int, len(prefix))
    copy(sorted, prefix)
    sort.Ints(sorted)
    st := NewSegmentTree(len(sorted))
    var count int
    for i := 0; i < len(prefix); i++ {
        k := sort.SearchInts(sorted, prefix[i])
        st.Decrement(k)
        l, u := sort.SearchInts(sorted, prefix[i]+lower), sort.SearchInts(sorted, prefix[i]+upper+1) - 1
        count += st.SumRange(l, u)
    }
    return count
}

func NewSegmentTree(m int) *SegmentTree {
    size := uint32(m)
    if bits.OnesCount32(size) != 1 {
        zeros := bits.LeadingZeros32(size)
        size = 1 << (32 - zeros)
    }
    n := int(size)
    tree := make([]int, 2*n)
    for i := 0; i < m; i++ {
        tree[i+n] = 1
    }
    for i := n - 1; i > 0; i-- {
        tree[i] = tree[2*i] + tree[2*i+1]
    }
    return &SegmentTree{tree: tree, n: n}
}

type SegmentTree struct {
    tree []int
    n    int
}

func (t *SegmentTree) Decrement(index int) {
    t.tree[index+t.n]--
    for p := (index + t.n) / 2; p > 0; p = p / 2 {
        t.tree[p] = t.tree[2*p] + t.tree[2*p+1]
    }
}

func (t *SegmentTree) SumRange(ql, qr int) int {
    return t.find(1, 0, t.n-1, ql, qr)
}

func (t *SegmentTree) find(node, nl, nr, ql, qr int) int {
    if ql <= nl && nr <= qr {
        return t.tree[node]
    }
    if qr < nl || ql > nr {
        return 0
    }
    h := (nl + nr) / 2
    return t.find(2*node, nl, h, ql, qr) + t.find(2*node+1, h+1, nr, ql, qr)
}

// 归并排序
func countRangeSum2(nums []int, lower int, upper int) int {
    res, temp, preSum := 0, make([]int, len(nums) + 1), make([]int, len(nums) + 1)
    for i := 0; i < len(preSum)-1; i++ {
        preSum[i+1] = preSum[i] + nums[i]
    }
    var merge func(left, mid, right int, nums []int)
    merge = func(left, mid, right int, nums []int) {
        for i := left; i <= right; i++ {
            temp[i] = nums[i]
        }
        start, end := mid + 1, mid + 1
        for i := left; i <= mid; i++ { // 更新结果
            for start <= right && nums[start]-nums[i] < lower { start++; }
            for end <= right && nums[end]-nums[i] <= upper { end++; }
            res += end - start
        }
        i, j, p := left, mid + 1, left
        for p <= right { // 合并2个有序数组
            if i == mid + 1 {
                nums[p] = temp[j]
                j++
            } else if j == right+1 {
                nums[p] = temp[i]
                i++
            } else if temp[i] > temp[j] {
                nums[p] = temp[j]
                j++
            } else if temp[i] <= temp[j] {
                nums[p] = temp[i]
                i++
            }
            p++
        }
    }
    var sortm func(left, right int, nums []int)
    sortm = func(left, right int, nums []int) {
        if left >= right {
            return
        }
        mid := left + (right-left)/2
        sortm(left, mid, preSum)
        sortm(mid+1, right, preSum)
        merge(left, mid, right, preSum)
    }
    sortm(0, len(preSum)-1, preSum)
    return res
}

func main() {
    // Example 1:
    // Input: nums = [-2,5,-1], lower = -2, upper = 2
    // Output: 3
    // Explanation: The three ranges are: [0,0], [2,2], and [0,2] and their respective sums are: -2, -1, 2.
    fmt.Println(countRangeSum([]int{-2,5,-1}, -2, 2)) // 3
    // Example 2:
    // Input: nums = [0], lower = 0, upper = 0
    // Output: 1
    fmt.Println(countRangeSum([]int{0}, 0, 0)) // 1

    fmt.Println(countRangeSum1([]int{-2,5,-1}, -2, 2)) // 3
    fmt.Println(countRangeSum1([]int{0}, 0, 0)) // 1

    fmt.Println(countRangeSum2([]int{-2,5,-1}, -2, 2)) // 3
    fmt.Println(countRangeSum2([]int{0}, 0, 0)) // 1
}