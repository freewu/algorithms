package main

// 3691. Maximum Total Subarray Value II
// You are given an integer array nums of length n and an integer k.

// You must select exactly k distinct non-empty subarrays nums[l..r] of nums. 
// Subarrays may overlap, but the exact same subarray (same l and r) cannot be chosen more than once.

// The value of a subarray nums[l..r] is defined as: max(nums[l..r]) - min(nums[l..r]).

// The total value is the sum of the values of all chosen subarrays.

// Return the maximum possible total value you can achieve.

// Example 1:
// Input: nums = [1,3,2], k = 2
// Output: 4
// Explanation:
// One optimal approach is:
// Choose nums[0..1] = [1, 3]. The maximum is 3 and the minimum is 1, giving a value of 3 - 1 = 2.
// Choose nums[0..2] = [1, 3, 2]. The maximum is still 3 and the minimum is still 1, so the value is also 3 - 1 = 2.
// Adding these gives 2 + 2 = 4.

// Example 2:
// Input: nums = [4,2,5,1], k = 3
// Output: 12
// Explanation:
// One optimal approach is:
// Choose nums[0..3] = [4, 2, 5, 1]. The maximum is 5 and the minimum is 1, giving a value of 5 - 1 = 4.
// Choose nums[1..3] = [2, 5, 1]. The maximum is 5 and the minimum is 1, so the value is also 4.
// Choose nums[2..3] = [5, 1]. The maximum is 5 and the minimum is 1, so the value is again 4.
// Adding these gives 4 + 4 + 4 = 12.

// Constraints:
//     1 <= n == nums.length <= 5 * 10^​​​​​​​4
//     0 <= nums[i] <= 10^9
//     1 <= k <= min(10^5, n * (n + 1) / 2)

import "fmt"
import "container/heap"
import "math/bits"

type Pair struct{ 
    min, max int
}

func op(a, b Pair) Pair {
	return Pair{ min(a.min, b.min), max(a.max, b.max) }
}

type Tuple struct{ d, l, r int }
type MaxHeap []Tuple

func (h MaxHeap)  Len() int           { return len(h) }
func (h MaxHeap)  Less(i, j int) bool { return h[i].d > h[j].d } // 最大堆
func (h MaxHeap)  Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(v any)         { *h = append(*h, v.(Tuple)) }
func (h *MaxHeap) Pop() (_ any)       { return }

type SegmentTree [][16]Pair // 16 = bits.Len(5e4)

func newSegmentTree(a []int) SegmentTree {
    n := len(a)
    w := bits.Len(uint(n))
    st := make(SegmentTree, n)
    for i, x := range a {
        st[i][0] = Pair{x, x}
    }
    for j := 1; j < w; j++ {
        for i := range n - 1<<j + 1 {
            st[i][j] = op(st[i][j-1], st[i+1<<(j-1)][j-1])
        }
    }
    return st
}

// [l,r) 左闭右开
func (st SegmentTree) query(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	p := op(st[l][k], st[r-1<<k][k])
	return p.max - p.min
}

func maxTotalValue(nums []int, k int) int64 {
    res, n := int64(0), len(nums)
    st := newSegmentTree(nums)
    h := MaxHeap{ {st.query(0, n), 0, n}} // 子数组值，左端点，右端点加一
    for ; k > 0 && h[0].d > 0; k-- {
        res += int64(h[0].d)
        l, r := h[0].l, h[0].r
        h[0].r--
        h[0].d = st.query(h[0].l, h[0].r)
        heap.Fix(&h, 0)
        if r == n && l+1 < n {
            heap.Push(&h, Tuple{st.query(l+1, n), l + 1, n})
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,3,2], k = 2
    // Output: 4
    // Explanation:
    // One optimal approach is:
    // Choose nums[0..1] = [1, 3]. The maximum is 3 and the minimum is 1, giving a value of 3 - 1 = 2.
    // Choose nums[0..2] = [1, 3, 2]. The maximum is still 3 and the minimum is still 1, so the value is also 3 - 1 = 2.
    // Adding these gives 2 + 2 = 4.
    fmt.Println(maxTotalValue([]int{1,3,2}, 2)) // 4
    // Example 2:
    // Input: nums = [4,2,5,1], k = 3
    // Output: 12
    // Explanation:
    // One optimal approach is:
    // Choose nums[0..3] = [4, 2, 5, 1]. The maximum is 5 and the minimum is 1, giving a value of 5 - 1 = 4.
    // Choose nums[1..3] = [2, 5, 1]. The maximum is 5 and the minimum is 1, so the value is also 4.
    // Choose nums[2..3] = [5, 1]. The maximum is 5 and the minimum is 1, so the value is again 4.
    // Adding these gives 4 + 4 + 4 = 12.
    fmt.Println(maxTotalValue([]int{4,2,5,1}, 3)) // 12

    fmt.Println(maxTotalValue([]int{1,2,3,4,5,6,7,8,9}, 2)) // 15
    fmt.Println(maxTotalValue([]int{9,8,7,6,5,4,3,2,1}, 2)) // 15
}
