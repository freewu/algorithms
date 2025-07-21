package main

// 3624. Number of Integers With Popcount-Depth Equal to K II
// You are given an integer array nums.

// For any positive integer x, define the following sequence:
//     1. p0 = x
//     2. pi+1 = popcount(pi) for all i >= 0, 
//        where popcount(y) is the number of set bits (1's) in the binary representation of y.

// This sequence will eventually reach the value 1.

// The popcount-depth of x is defined as the smallest integer d >= 0 such that pd = 1.

// For example, if x = 7 (binary representation "111"). 
// Then, the sequence is: 7 → 3 → 2 → 1, so the popcount-depth of 7 is 3.

// You are also given a 2D integer array queries, where each queries[i] is either:
//     1. [1, l, r, k] - Determine the number of indices j such that l <= j <= r and the popcount-depth of nums[j] is equal to k.
//     2. [2, idx, val] - Update nums[idx] to val.

// Return an integer array answer, where answer[i] is the number of indices for the ith query of type [1, l, r, k].

// Example 1:
// Input: nums = [2,4], queries = [[1,0,1,1],[2,1,1],[1,0,1,0]]
// Output: [2,1]
// Explanation:
// i	queries[i]	nums	binary(nums)	popcount-
// depth	[l, r]	k	Valid
// nums[j]	updated
// nums	Answer
// 0	[1,0,1,1]	[2,4]	[10, 100]	[1, 1]	[0, 1]	1	[0, 1]	—	2
// 1	[2,1,1]	[2,4]	[10, 100]	[1, 1]	—	—	—	[2,1]	—
// 2	[1,0,1,0]	[2,1]	[10, 1]	[1, 0]	[0, 1]	0	[1]	—	1
// Thus, the final answer is [2, 1].

// Example 2:
// Input: nums = [3,5,6], queries = [[1,0,2,2],[2,1,4],[1,1,2,1],[1,0,1,0]]
// Output: [3,1,0]
// Explanation:
// i	queries[i]	nums	binary(nums)	popcount-
// depth	[l, r]	k	Valid
// nums[j]	updated
// nums	Answer
// 0	[1,0,2,2]	[3, 5, 6]	[11, 101, 110]	[2, 2, 2]	[0, 2]	2	[0, 1, 2]	—	3
// 1	[2,1,4]	[3, 5, 6]	[11, 101, 110]	[2, 2, 2]	—	—	—	[3, 4, 6]	—
// 2	[1,1,2,1]	[3, 4, 6]	[11, 100, 110]	[2, 1, 2]	[1, 2]	1	[1]	—	1
// 3	[1,0,1,0]	[3, 4, 6]	[11, 100, 110]	[2, 1, 2]	[0, 1]	0	[]	—	0
// Thus, the final answer is [3, 1, 0].

// Example 3:
// Input: nums = [1,2], queries = [[1,0,1,1],[2,0,3],[1,0,0,1],[1,0,0,2]]
// Output: [1,0,1]
// Explanation:
// i	queries[i]	nums	binary(nums)	popcount-
// depth	[l, r]	k	Valid
// nums[j]	updated
// nums	Answer
// 0	[1,0,1,1]	[1, 2]	[1, 10]	[0, 1]	[0, 1]	1	[1]	—	1
// 1	[2,0,3]	[1, 2]	[1, 10]	[0, 1]	—	—	—	[3, 2]	 
// 2	[1,0,0,1]	[3, 2]	[11, 10]	[2, 1]	[0, 0]	1	[]	—	0
// 3	[1,0,0,2]	[3, 2]	[11, 10]	[2, 1]	[0, 0]	2	[0]	—	1
// Thus, the final answer is [1, 0, 1].
 
// Constraints:
//     1 <= n == nums.length <= 10^5
//     1 <= nums[i] <= 10^15
//     1 <= queries.length <= 10^5
//     queries[i].length == 3 or 4
//     queries[i] == [1, l, r, k] or,
//     queries[i] == [2, idx, val]
//     0 <= l <= r <= n - 1
//     0 <= k <= 5
//     0 <= idx <= n - 1
//     1 <= val <= 10^15

import "fmt"
import "math/bits"

type Fenwick []int

func newFenwickTree(n int) Fenwick {
    return make(Fenwick, n + 1) // 使用下标 1 到 n
}

// a[i] 增加 val
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f Fenwick) update(i int, val int) {
    for ; i < len(f); i += i & -i {
        f[i] += val
    }
}

// 求前缀和 a[1] + ... + a[i]
// 1 <= i <= n
// 时间复杂度 O(log n)
func (f Fenwick) pre(i int) (res int) {
    for ; i > 0; i &= i - 1 {
        res += f[i]
    }
    return
}

// 求区间和 a[l] + ... + a[r]
// 1 <= l <= r <= n
// 时间复杂度 O(log n)
func (f Fenwick) query(l, r int) int {
    return f.pre(r) - f.pre(l-1)
}

func popcountDepth(nums []int64, queries [][]int64) []int  {
    n := len(nums)
    res, f := []int{}, [6]Fenwick{}
    for i := range f {
        f[i] = newFenwickTree(n)
    }
    popDepth := func(x uint64) (res int) { // 不写记忆化更快，直接迭代
        for x > 1 {
            res++
            x = uint64(bits.OnesCount64(x))
        }
        return
    }
    update := func(i, delta int) {
        d := popDepth(uint64(nums[i]))
        f[d].update(i+1, delta)
    }
    for i := 0; i < n; i++ {
        update(i, 1) // 添加
    }
    for _, q := range queries {
        if q[0] == 1 {
            res = append(res, f[q[3]].query(int(q[1])+1, int(q[2])+1))
        } else {
            i := int(q[1])
            update(i, -1) // 撤销旧的
            nums[i] = q[2]
            update(i, 1) // 添加新的
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [2,4], queries = [[1,0,1,1],[2,1,1],[1,0,1,0]]
    // Output: [2,1]
    // Explanation:
    // i	queries[i]	nums	binary(nums)	popcount-
    // depth	[l, r]	k	Valid
    // nums[j]	updated
    // nums	Answer
    // 0	[1,0,1,1]	[2,4]	[10, 100]	[1, 1]	[0, 1]	1	[0, 1]	—	2
    // 1	[2,1,1]	[2,4]	[10, 100]	[1, 1]	—	—	—	[2,1]	—
    // 2	[1,0,1,0]	[2,1]	[10, 1]	[1, 0]	[0, 1]	0	[1]	—	1
    // Thus, the final answer is [2, 1].
    fmt.Println(popcountDepth([]int64{2,4}, [][]int64{{1,0,1,1},{2,1,1},{1,0,1,0}})) // [2,1]
    // Example 2:
    // Input: nums = [3,5,6], queries = [[1,0,2,2],[2,1,4],[1,1,2,1],[1,0,1,0]]
    // Output: [3,1,0]
    // Explanation:
    // i	queries[i]	nums	binary(nums)	popcount-
    // depth	[l, r]	k	Valid
    // nums[j]	updated
    // nums	Answer
    // 0	[1,0,2,2]	[3, 5, 6]	[11, 101, 110]	[2, 2, 2]	[0, 2]	2	[0, 1, 2]	—	3
    // 1	[2,1,4]	[3, 5, 6]	[11, 101, 110]	[2, 2, 2]	—	—	—	[3, 4, 6]	—
    // 2	[1,1,2,1]	[3, 4, 6]	[11, 100, 110]	[2, 1, 2]	[1, 2]	1	[1]	—	1
    // 3	[1,0,1,0]	[3, 4, 6]	[11, 100, 110]	[2, 1, 2]	[0, 1]	0	[]	—	0
    // Thus, the final answer is [3, 1, 0].
    fmt.Println(popcountDepth([]int64{3,5,6}, [][]int64{{1,0,2,2},{2,1,4},{1,1,2,1},{1,0,1,0}})) // [3,1,0]
    // Example 3:
    // Input: nums = [1,2], queries = [[1,0,1,1],[2,0,3],[1,0,0,1],[1,0,0,2]]
    // Output: [1,0,1]
    // Explanation:
    // i	queries[i]	nums	binary(nums)	popcount-
    // depth	[l, r]	k	Valid
    // nums[j]	updated
    // nums	Answer
    // 0	[1,0,1,1]	[1, 2]	[1, 10]	[0, 1]	[0, 1]	1	[1]	—	1
    // 1	[2,0,3]	[1, 2]	[1, 10]	[0, 1]	—	—	—	[3, 2]	 
    // 2	[1,0,0,1]	[3, 2]	[11, 10]	[2, 1]	[0, 0]	1	[]	—	0
    // 3	[1,0,0,2]	[3, 2]	[11, 10]	[2, 1]	[0, 0]	2	[0]	—	1
    // Thus, the final answer is [1, 0, 1].
    fmt.Println(popcountDepth([]int64{1,2}, [][]int64{{1,0,1,1},{2,0,3},{1,0,0,1},{1,0,0,2}})) // [1,0,1]
}