package main

// 3915. Maximum Sum of Alternating Subsequence With Distance at Least K
// You are given an integer array nums of length n and an integer k.

// Pick a subsequence with indices 0 <= i1 < i2 < ... < im < n such that:
//     1. For every 1 <= t < m, it+1 - it >= k.
//     2. The selected values form a strictly alternating sequence. In other words, either:
//         2.1 nums[i1] < nums[i2] > nums[i3] < ..., or
//         2.2 nums[i1] > nums[i2] < nums[i3] > ...

// A subsequence of length 1 is also considered strictly alternating. 
// The score of a valid subsequence is the sum of its selected values.

// Return an integer denoting the maximum possible score of a valid subsequence.

// Example 1:
// Input: nums = [5,4,2], k = 2
// Output: 7
// Explanation:
// An optimal choice is indices [0, 2], which gives values [5, 2].
// The distance condition holds because 2 - 0 = 2 >= k.
// The values are strictly alternating because 5 > 2.
// The score is 5 + 2 = 7.

// Example 2:
// Input: nums = [3,5,4,2,4], k = 1
// Output: 14
// Explanation:
// An optimal choice is indices [0, 1, 3, 4], which gives values [3, 5, 2, 4].
// The distance condition holds because each pair of consecutive chosen indices differs by at least k = 1.
// The values are strictly alternating since 3 < 5 > 2 < 4.
// The score is 3 + 5 + 2 + 4 = 14.

// Example 3:
// Input: nums = [5], k = 1
// Output: 5
// Explanation:
// The only valid subsequence is [5]. A subsequence with 1 element is always strictly alternating, so the score is 5.

// Constraints:
//     1 <= n == nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     1 <= k <= n

import "fmt"
import "sort"
import "slices"

type Fenwick []int64

func (f Fenwick) update(i int, val int64) {
    for ; i < len(f); i += i & -i {
        f[i] = max(f[i], val)
    }
}

// [1, i] 中的最大值
func (f Fenwick) preMax(i int) (res int64) {
    for ; i > 0; i &= i - 1 {
        res = max(res, f[i])
    }
    return
}

func maxAlternatingSum(nums []int, k int) int64 {
    // 离散化 nums
    sorted := slices.Clone(nums)
    slices.Sort(sorted)
    sorted = slices.Compact(sorted)
    res, n := int64(0), len(nums)
    fInc := make([]int64, n) // fInc[i] 表示以 nums[i] 结尾且最后两项递增的交替子序列的最大和
    fDec := make([]int64, n) // fDec[i] 表示以 nums[i] 结尾且最后两项递减的交替子序列的最大和
    // 值域树状数组
    m := len(sorted)
    inc := make(Fenwick, m + 1) // 维护 fInc[i] 的最大值
    dec := make(Fenwick, m + 1) // 维护 fDec[i] 的最大值
    for i, x := range nums {
        if i >= k {
            // 在这个时候才把 fInc[i-k] 和 fDec[i-k] 添加到值域树状数组中，从而保证转移来源的下标 <= i-k
            j := nums[i-k]
            inc.update(m-j, fInc[i-k]) // m-j 可以把后缀变成前缀
            dec.update(j+1, fDec[i-k])
        }
        j := sort.SearchInts(sorted, x)
        nums[i] = j // 注意这里修改了 nums[i]，这样上面的 nums[i-k] 无需二分
        fInc[i] = dec.preMax(j) + int64(x)     // 计算满足 nums[i'] < x 的 fDec[i'] 的最大值
        fDec[i] = inc.preMax(m-1-j) + int64(x) // 计算满足 nums[i'] > x 的 fInc[i'] 的最大值
        res = max(res, fInc[i], fDec[i])       // 枚举子序列以 nums[i] 结尾
    }
    return res
}

func maxAlternatingSum1(nums []int, k int) int64 {
    res, mx, n := int64(0), int64(0), len(nums)
    for _, v := range nums {
        if int64(v) > mx {
            mx = int64(v)
        }
    }
    bitDown,bitUp := make([]int64, mx + 2), make([]int64, mx + 2)
    update := func(tree []int64, i int64, val int64) {
        for ; i < int64(len(tree)); i += i & -i {
            if val > tree[i] {
                tree[i] = val
            }
        }
    }
    query := func(tree []int64, i int64) int64 {
        res := int64(0)
        for ; i > 0; i -= i & -i {
            if tree[i] > res {
                res = tree[i]
            }
        }
        return res
    }
    dpDown, dpUp := make([]int64, n), make([]int64, n)
    for i := 0; i < n; i++ {
        if i >= k {
            prevIndex := i - k
            prevVal := int64(nums[prevIndex])
            update(bitDown, prevVal, dpDown[prevIndex])
            update(bitUp, mx-prevVal+1, dpUp[prevIndex])
        }
        val := int64(nums[i])
        dpUp[i] = val + query(bitDown, val - 1)
        dpDown[i] = val + query(bitUp, mx - val)
        if dpUp[i] > res {
            res = dpUp[i]
        }
        if dpDown[i] > res {
            res = dpDown[i]
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [5,4,2], k = 2
    // Output: 7
    // Explanation:
    // An optimal choice is indices [0, 2], which gives values [5, 2].
    // The distance condition holds because 2 - 0 = 2 >= k.
    // The values are strictly alternating because 5 > 2.
    // The score is 5 + 2 = 7.
    fmt.Println(maxAlternatingSum([]int{5,4,2}, 2)) // 7
    // Example 2:
    // Input: nums = [3,5,4,2,4], k = 1
    // Output: 14
    // Explanation:
    // An optimal choice is indices [0, 1, 3, 4], which gives values [3, 5, 2, 4].
    // The distance condition holds because each pair of consecutive chosen indices differs by at least k = 1.
    // The values are strictly alternating since 3 < 5 > 2 < 4.
    // The score is 3 + 5 + 2 + 4 = 14.
    fmt.Println(maxAlternatingSum([]int{3,5,4,2,4}, 1)) // 14
    // Example 3:
    // Input: nums = [5], k = 1
    // Output: 5
    // Explanation:
    // The only valid subsequence is [5]. A subsequence with 1 element is always strictly alternating, so the score is 5.
    fmt.Println(maxAlternatingSum([]int{5}, 1)) // 5

    fmt.Println(maxAlternatingSum([]int{1,2,3,4,5,6,7,8,9}, 2)) // 16
    fmt.Println(maxAlternatingSum([]int{9,8,7,6,5,4,3,2,1}, 2)) // 16

    fmt.Println(maxAlternatingSum1([]int{5,4,2}, 2)) // 7
    fmt.Println(maxAlternatingSum1([]int{3,5,4,2,4}, 1)) // 14
    fmt.Println(maxAlternatingSum1([]int{5}, 1)) // 5
    fmt.Println(maxAlternatingSum1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 16
    fmt.Println(maxAlternatingSum1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 16
}