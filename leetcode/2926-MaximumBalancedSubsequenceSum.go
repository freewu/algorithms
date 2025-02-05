package main

// 2926. Maximum Balanced Subsequence Sum
// You are given a 0-indexed integer array nums.

// A subsequence of nums having length k and consisting of indices i0 < i1 < ... < ik-1 is balanced if the following holds:
//     nums[ij] - nums[ij-1] >= ij - ij-1, for every j in the range [1, k - 1].

// A subsequence of nums having length 1 is considered balanced.

// Return an integer denoting the maximum possible sum of elements in a balanced subsequence of nums.

// A subsequence of an array is a new non-empty array that is formed from the original array by deleting some (possibly none) of the elements without disturbing the relative positions of the remaining elements.

// Example 1:
// Input: nums = [3,3,5,6]
// Output: 14
// Explanation: In this example, the subsequence [3,5,6] consisting of indices 0, 2, and 3 can be selected.
// nums[2] - nums[0] >= 2 - 0.
// nums[3] - nums[2] >= 3 - 2.
// Hence, it is a balanced subsequence, and its sum is the maximum among the balanced subsequences of nums.
// The subsequence consisting of indices 1, 2, and 3 is also valid.
// It can be shown that it is not possible to get a balanced subsequence with a sum greater than 14.

// Example 2:
// Input: nums = [5,-1,-3,8]
// Output: 13
// Explanation: In this example, the subsequence [5,8] consisting of indices 0 and 3 can be selected.
// nums[3] - nums[0] >= 3 - 0.
// Hence, it is a balanced subsequence, and its sum is the maximum among the balanced subsequences of nums.
// It can be shown that it is not possible to get a balanced subsequence with a sum greater than 13.

// Example 3:
// Input: nums = [-2,-1]
// Output: -1
// Explanation: In this example, the subsequence [-1] can be selected.
// It is a balanced subsequence, and its sum is the maximum among the balanced subsequences of nums.

// Constraints:
//     1 <= nums.length <= 10^5
//     -10^9 <= nums[i] <= 10^9

import "fmt"
import "sort"
import "slices"

type SegmentTree struct {
    vec []int64
}

func (st *SegmentTree) build (n int) {
    for n & (n - 1) != 0 {
        n += 1
    }
    st.vec = make([]int64, 2 * n)
}

func (st *SegmentTree) query_rec (dum int, tl int, tr int, l int, r int) int64 {
    if (tl >= l && tr <= r) {
        return st.vec[dum]
    }
    if tl > r || tr < l {
        return -1e15
    }
    return max(st.query_rec(2 * dum + 1, tl, (tl + tr)/2, l, r), st.query_rec(2 * dum + 2, (tl + tr)/2 + 1, tr, l, r))
}

func (st *SegmentTree) query (l int, r int) int64 {
    return st.query_rec(0, 0, len(st.vec)/2 - 1, l, r)
}

func (st *SegmentTree) update (x int, y int64) {
    x += len(st.vec)/2 - 1
    st.vec[x] = y
    for x != 0 {
        x = (x - 1)/2
        st.vec[x] = max(st.vec[2 * x + 1] , st.vec[2 * x + 2])
    }
}

func maxBalancedSubsequenceSum(nums []int) int64 {
    n := len(nums)
    vec := make([][]int, len(nums))
    for i := 0; i < len(nums); i++ {
        vec[i] = []int{ nums[i] - i, i }
    }
    sort.Slice (vec, func(i int, j int) bool {
        if vec[i][0] != vec[j][0] { return (vec[i][0] < vec[j][0]) }
        return vec[i][1] < vec[j][1]
    })
    mp := make(map[int]int)
    for i, v := range vec {
        mp[v[0]] = i + 1
    }
    st := SegmentTree{}
    st.build(n + 10)
    for i := 0; i < len(st.vec) / 2 ; i++ {
        st.update(i, -1e15)
        st.query(i, i)
    }
    st.update(0, 0)
    for i, v := range nums {
        st.update(mp[v - i], st.query(0, mp[v - i]) + int64(v))
    }
    return st.query(1, n + 1)
}

// 树状数组模板（维护前缀最大值）
type Fenwick []int

func (f Fenwick) update(i, val int) {
    for ; i < len(f); i += i & -i {
        f[i] = max(f[i], val)
    }
}

func (f Fenwick) preMax(i int) int {
    mx := -1 << 31
    for ; i > 0; i &= i - 1 {
        mx = max(mx, f[i])
    }
    return mx
}

func maxBalancedSubsequenceSum1(nums []int) int64 {
    // 离散化 nums[i] - i
    b := slices.Clone(nums)
    for i := range b {
        b[i] -= i
    }
    slices.Sort(b)
    b = slices.Compact(b) // 去重
    // 初始化树状数组
    t := make(Fenwick, len(b)+1)
    for i := range t {
        t[i] = -1 << 31
    }
    for i, v := range nums {
        j := sort.SearchInts(b, v - i) + 1 // nums[i]-i 离散化后的值（从 1 开始）
        f := max(t.preMax(j), 0) + v
        t.update(j, f)
    }
    return int64(t.preMax(len(b)))
}

func main() {
    // Example 1:
    // Input: nums = [3,3,5,6]
    // Output: 14
    // Explanation: In this example, the subsequence [3,5,6] consisting of indices 0, 2, and 3 can be selected.
    // nums[2] - nums[0] >= 2 - 0.
    // nums[3] - nums[2] >= 3 - 2.
    // Hence, it is a balanced subsequence, and its sum is the maximum among the balanced subsequences of nums.
    // The subsequence consisting of indices 1, 2, and 3 is also valid.
    // It can be shown that it is not possible to get a balanced subsequence with a sum greater than 14.
    fmt.Println(maxBalancedSubsequenceSum([]int{3,3,5,6})) // 14
    // Example 2:
    // Input: nums = [5,-1,-3,8]
    // Output: 13
    // Explanation: In this example, the subsequence [5,8] consisting of indices 0 and 3 can be selected.
    // nums[3] - nums[0] >= 3 - 0.
    // Hence, it is a balanced subsequence, and its sum is the maximum among the balanced subsequences of nums.
    // It can be shown that it is not possible to get a balanced subsequence with a sum greater than 13.
    fmt.Println(maxBalancedSubsequenceSum([]int{5,-1,-3,8})) // 13
    // Example 3:
    // Input: nums = [-2,-1]
    // Output: -1
    // Explanation: In this example, the subsequence [-1] can be selected.
    // It is a balanced subsequence, and its sum is the maximum among the balanced subsequences of nums.
    fmt.Println(maxBalancedSubsequenceSum([]int{-2,-1})) // -1

    fmt.Println(maxBalancedSubsequenceSum([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxBalancedSubsequenceSum([]int{9,8,7,6,5,4,3,2,1})) // 9

    fmt.Println(maxBalancedSubsequenceSum1([]int{3,3,5,6})) // 14
    fmt.Println(maxBalancedSubsequenceSum1([]int{5,-1,-3,8})) // 13
    fmt.Println(maxBalancedSubsequenceSum1([]int{-2,-1})) // -1
    fmt.Println(maxBalancedSubsequenceSum1([]int{1,2,3,4,5,6,7,8,9})) // 45
    fmt.Println(maxBalancedSubsequenceSum1([]int{9,8,7,6,5,4,3,2,1})) // 9
}