package main

// 2407. Longest Increasing Subsequence II
// You are given an integer array nums and an integer k.

// Find the longest subsequence of nums that meets the following requirements:
//     1. The subsequence is strictly increasing and
//     2. The difference between adjacent elements in the subsequence is at most k.

// Return the length of the longest subsequence that meets the requirements.

// A subsequence is an array that can be derived from another array by deleting some or no elements without changing the order of the remaining elements.

// Example 1:
// Input: nums = [4,2,1,4,3,4,5,8,15], k = 3
// Output: 5
// Explanation:
// The longest subsequence that meets the requirements is [1,3,4,5,8].
// The subsequence has a length of 5, so we return 5.
// Note that the subsequence [1,3,4,5,8,15] does not meet the requirements because 15 - 8 = 7 is larger than 3.

// Example 2:
// Input: nums = [7,4,5,1,8,12,4,7], k = 5
// Output: 4
// Explanation:
// The longest subsequence that meets the requirements is [4,5,8,12].
// The subsequence has a length of 4, so we return 4.

// Example 3:
// Input: nums = [1,5], k = 1
// Output: 1
// Explanation:
// The longest subsequence that meets the requirements is [1].
// The subsequence has a length of 1, so we return 1.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i], k <= 10^5

import "fmt"
import "slices"

// // 83 / 84 个通过的测试用例
// func lengthOfLIS(nums []int, k int) int {
//     list := make([]int, 1e5 + 1)
//     getMax := func(arr ...int) int {
//         res := arr[0]
//         for i := range arr {
//             if arr[i] > res {
//                 res = arr[i]
//             }
//         }
//         return res
//     }
//     for i := range nums {
//         if nums[i] >= k {
//             list[nums[i]] = getMax(list[nums[i] - k:nums[i]]...) + 1
//         } else {
//             list[nums[i]] = getMax(list[:nums[i]]...) + 1
//         }
//     }
//     return getMax(list...)
// }

func lengthOfLIS(nums []int, k int) int {
    res, mx := 1, slices.Max(nums)
    tree := newSegmentTree(mx)
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for _, v := range nums {
        t := tree.query(1, v-k, v-1) + 1
        res = max(res, t)
        tree.modify(1, v, t)
    }
    return res
}

type Node struct {
    l int
    r int
    v int
}

type SegmentTree struct {
    tr []*Node
}

func newSegmentTree(n int) *SegmentTree {
    tr := make([]*Node, n<<2)
    for i := range tr {
        tr[i] = &Node{}
    }
    t := &SegmentTree{tr}
    t.build(1, 1, n)
    return t
}

func (t *SegmentTree) build(u, l, r int) {
    t.tr[u].l, t.tr[u].r = l, r
    if l == r { return }
    mid := (l + r) >> 1
    t.build(u<<1, l, mid)
    t.build(u<<1|1, mid+1, r)
    t.pushup(u)
}

func (t *SegmentTree) modify(u, x, v int) {
    if t.tr[u].l == x && t.tr[u].r == x {
        t.tr[u].v = v
        return
    }
    mid := (t.tr[u].l + t.tr[u].r) >> 1
    if x <= mid {
        t.modify(u<<1, x, v)
    } else {
        t.modify(u<<1|1, x, v)
    }
    t.pushup(u)
}

func (t *SegmentTree) query(u, l, r int) int {
    if t.tr[u].l >= l && t.tr[u].r <= r {
        return t.tr[u].v
    }
    mid := (t.tr[u].l + t.tr[u].r) >> 1
    v := 0
    if l <= mid {
        v = t.query(u<<1, l, r)
    }
    if r > mid {
        v = max(v, t.query(u<<1|1, l, r))
    }
    return v
}

func (t *SegmentTree) pushup(u int) {
    t.tr[u].v = max(t.tr[u<<1].v, t.tr[u<<1|1].v)
}

func lengthOfLIS1(nums []int, k int) int {
    res, mx := 0, 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    for i := 0; i < len(nums); i++ {
        mx = max(mx, nums[i])
    }
    dp := make([]int, 2 * (mx + 1))
    query := func(l int, r int, n int) int {
        res := 0
        for l, r = l + n, r + n; l < r;{
            if l % 2 == 1 {
                res = max(res, dp[l])
                l++
            }
            if r % 2 == 1 {
                r--
                res = max(res, dp[r])
            }
            l >>= 1
            r >>= 1
        }
        return res
    }
    update := func (i int, val int, mx int){
        i += mx
        dp[i] = val
        for i > 1 {
            i >>= 1
            dp[i] = max(dp[i * 2], dp[i * 2 + 1])
        }
    }
    for i := 0; i < len(nums); i++ {
        cur := 1 + query(max(1, nums[i] - k), nums[i], mx)
        update(nums[i], cur, mx)
        res = max(res, cur)
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [4,2,1,4,3,4,5,8,15], k = 3
    // Output: 5
    // Explanation:
    // The longest subsequence that meets the requirements is [1,3,4,5,8].
    // The subsequence has a length of 5, so we return 5.
    // Note that the subsequence [1,3,4,5,8,15] does not meet the requirements because 15 - 8 = 7 is larger than 3.
    fmt.Println(lengthOfLIS([]int{4,2,1,4,3,4,5,8,15}, 3)) // 5
    // Example 2:
    // Input: nums = [7,4,5,1,8,12,4,7], k = 5
    // Output: 4
    // Explanation:
    // The longest subsequence that meets the requirements is [4,5,8,12].
    // The subsequence has a length of 4, so we return 4.
    fmt.Println(lengthOfLIS([]int{7,4,5,1,8,12,4,7}, 5)) // 4
    // Example 3:
    // Input: nums = [1,5], k = 1
    // Output: 1
    // Explanation:
    // The longest subsequence that meets the requirements is [1].
    // The subsequence has a length of 1, so we return 1.
    fmt.Println(lengthOfLIS([]int{1,5}, 1)) // 1

    fmt.Println(lengthOfLIS1([]int{4,2,1,4,3,4,5,8,15}, 3)) // 5
    fmt.Println(lengthOfLIS1([]int{7,4,5,1,8,12,4,7}, 5)) // 4
    fmt.Println(lengthOfLIS1([]int{1,5}, 1)) // 1
}