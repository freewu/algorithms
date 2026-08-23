package main

// 4033. Valid K-Unique Subarrays I
// You are given an integer array nums and an integer k.

// You are also given a 2D integer array queries, where queries[i] = [li, ri] represents the subarray nums[li..ri].

// For each query, the subarray nums[li..ri] is considered valid if:
//     1. It contains exactly k distinct numbers, and
//     2. The frequency of every number in the subarray is even.

// Return a boolean array ans, where ans[i] is true if nums[li..ri] is valid, and false otherwise.

// Example 1:
// Input: nums = [1,2,2,1], k = 2, queries = [[0,1],[0,3],[1,2]]
// Output: [false,true,false]
// Explanation:
// i | [li, ri] | Subarray     | Unique numbers | Frequency        | Validity check
// 0 | [0, 1]   | [1, 2]       | {1, 2} → 2     | {1: 1, 2: 1}	    | false: Element counts are not even.
// 1 | [0, 3]   | [1, 2, 2, 1] | {1, 2} → 2     | {1: 2, 2: 2}	    | true: Exactly k = 2 distinct elements, all appear an even number of times.
// 2 | [1, 2]   | [2, 2]       | {2} → 1        | {2: 2}	        | false: Number of distinct elements is less than k = 2.
// Thus, ans = [false, true, false].

// Example 2:
// Input: nums = [3,3,3], k = 1, queries = [[1,2],[0,2]]
// Output: [true,false]
// Explanation:
// i | [li, ri] | Subarray     | Unique numbers | Frequency        | Validity check
// 0 | [1, 2]   | [3, 3]       | {3} → 1        | {3: 2}	        | true: Exactly k = 1 distinct element, appears an even number of times.
// 1 | [0, 2]   | [3, 3, 3]    | {3} → 1        | {3: 3}	        | false: 3 does not appear an even number of times.
// Thus, ans = [true, false].

// Constraints:
//     2 <= n == nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     1 <= k <= n
//     1 <= queries.length <= 10^5
//     queries[i] == [li, ri]
//     0 <= li < ri <= n - 1

import "fmt"
import "math/rand"

type Fenwick []int

func NewFenwickTree(n int) Fenwick {
    return make(Fenwick, n + 1) // 使用下标 1 到 n
}

// a[i] 增加 val
// 时间复杂度 O(log n)
func (f Fenwick) update(i, val int) {
    for i++; i < len(f); i += i & -i {
        f[i] += val
    }
}

// 求前缀和 a[1] + ... + a[i]
// 时间复杂度 O(log n)
func (f Fenwick) pre(i int) (res int) {
    for i++; i > 0; i &= i - 1 {
        res += f[i]
    }
    return
}

// 求区间和 a[l] + ... + a[r]
// 时间复杂度 O(log n)
func (f Fenwick) query(l, r int) int {
    return f.pre(r) - f.pre(l-1)
}

func validSubarrays(nums []int, k int, queries [][]int) []bool {
    n := len(nums)
    sum := make([]uint64, n+1)
    mp := map[int]uint64{}
    for i, x := range nums {
        // 把不同的 nums[i] 映射成一个随机的 uint64
        if _, ok := mp[x]; !ok {
            mp[x] = rand.Uint64()
        }
        sum[i+1] = sum[i] ^ mp[x]
    }
    // 离线询问：按照右端点分组
    type Pair struct{ l, qid int }
    groups := make([][]Pair, n)
    for i, q := range queries {
        groups[q[1]] = append(groups[q[1]], Pair{q[0], i})
    }
    t := NewFenwickTree(n)
    last := make(map[int]int, len(mp)) // 预分配空间    
    res := make([]bool, len(queries))
    for r, x := range nums {
        if i, ok := last[x]; ok {
            t.update(i, -1)
        }
        last[x] = r
        t.update(r, 1)
        for _, p := range groups[r] {
            res[p.qid] = sum[r+1] == sum[p.l] && t.query(p.l, r) == k
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [1,2,2,1], k = 2, queries = [[0,1],[0,3],[1,2]]
    // Output: [false,true,false]
    // Explanation:
    // i | [li, ri] | Subarray     | Unique numbers | Frequency        | Validity check
    // 0 | [0, 1]   | [1, 2]       | {1, 2} → 2     | {1: 1, 2: 1}	    | false: Element counts are not even.
    // 1 | [0, 3]   | [1, 2, 2, 1] | {1, 2} → 2     | {1: 2, 2: 2}	    | true: Exactly k = 2 distinct elements, all appear an even number of times.
    // 2 | [1, 2]   | [2, 2]       | {2} → 1        | {2: 2}	        | false: Number of distinct elements is less than k = 2.
    // Thus, ans = [false, true, false].
    fmt.Println(validSubarrays([]int{1,2,2,1}, 2, [][]int{{0,1},{0,3},{1,2}})) // [false,true,false]
    // Example 2:
    // Input: nums = [3,3,3], k = 1, queries = [[1,2],[0,2]]
    // Output: [true,false]
    // Explanation:
    // i | [li, ri] | Subarray     | Unique numbers | Frequency        | Validity check
    // 0 | [1, 2]   | [3, 3]       | {3} → 1        | {3: 2}	        | true: Exactly k = 1 distinct element, appears an even number of times.
    // 1 | [0, 2]   | [3, 3, 3]    | {3} → 1        | {3: 3}	        | false: 3 does not appear an even number of times.
    // Thus, ans = [true, false].
    fmt.Println(validSubarrays([]int{3,3,3}, 1, [][]int{{1,2},{0,2}})) // [true,false]

    fmt.Println(validSubarrays([]int{1,2,3,4,5,6,7,8,9}, 1, [][]int{{1,2},{0,2}})) // [false,false]
    fmt.Println(validSubarrays([]int{9,8,7,6,5,4,3,2,1}, 1, [][]int{{1,2},{0,2}})) // [false,false]
}