package main

// 3671. Sum of Beautiful Subsequences
// You are given an integer array nums of length n.

// For every positive integer g, we define the beauty of g as the product of g and the number of strictly increasing subsequences of nums whose greatest common divisor (GCD) is exactly g.

// Return the sum of beauty values for all positive integers g.

// Since the answer could be very large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [1,2,3]
// Output: 10
// Explanation:
// All strictly increasing subsequences and their GCDs are:
// Subsequence	GCD
// [1]	1
// [2]	2
// [3]	3
// [1,2]	1
// [1,3]	1
// [2,3]	1
// [1,2,3]	1
// Calculating beauty for each GCD:
// GCD	Count of subsequences	Beauty (GCD × Count)
// 1	5	1 × 5 = 5
// 2	1	2 × 1 = 2
// 3	1	3 × 1 = 3
// Total beauty is 5 + 2 + 3 = 10.

// Example 2:
// Input: nums = [4,6]
// Output: 12
// Explanation:
// All strictly increasing subsequences and their GCDs are:
// Subsequence	GCD
// [4]	4
// [6]	6
// [4,6]	2
// Calculating beauty for each GCD:
// GCD	Count of subsequences	Beauty (GCD × Count)
// 2	1	2 × 1 = 2
// 4	1	4 × 1 = 4
// 6	1	6 × 1 = 6
// Total beauty is 2 + 4 + 6 = 12.

// Constraints:
//     1 <= n == nums.length <= 10^4
//     1 <= nums[i] <= 7 * 10^4

import "fmt"
import "slices"

const mod = 1_000_000_007
const mx = 70_001

var divisors [mx][]int

func init() {
    // 预处理每个数的因子
    for i := 1; i < mx; i++ {
        for j := i; j < mx; j += i { // 枚举 i 的倍数 j
            divisors[j] = append(divisors[j], i) // i 是 j 的因子
        }
    }
}

func totalBeauty(nums []int) int {
    res, now, mx := 0, 0, slices.Max(nums)
    // 树状数组（时间戳优化）
    tree, time := make([]int, mx + 1), make([]int, mx + 1) // 避免反复初始化树状数组
    update := func(i, val int) {
        for ; i <= mx; i += i & -i {
            if time[i] < now {
                time[i] = now
                tree[i] = 0 // 懒重置
            }
            tree[i] += val
        }
    }
    pre := func(i int) int {
        sum := 0
        for ; i > 0; i &= i - 1 {
            if time[i] == now {
                sum += tree[i]
            }
        }
        return sum % mod
    }
    // 计算 b 的严格递增子序列的个数
    countIncreasingSubsequence := func(b []int) (res int) {
        now++ // 重置树状数组（懒重置）
        for _, x := range b {
            // cnt 表示以 x 结尾的严格递增子序列的个数
            cnt := pre(x-1) + 1 // +1 是因为 x 可以一个数组成一个子序列
            res += cnt
            update(x, cnt) // 更新以 x 结尾的严格递增子序列的个数
        }
        return res % mod
    }
    groups := make([][]int, mx + 1)
    for _, x := range nums {
        for _, d := range divisors[x] {
            groups[d] = append(groups[d], x)
        }
    }
    f := make([]int, mx + 1)
    for i := mx; i > 0; i-- {
        f[i] = countIncreasingSubsequence(groups[i])
        // 倍数容斥
        for j := i * 2; j <= mx; j += i {
            f[i] -= f[j]
        }
        // 注意 |f[i]| * i < mod * (m / i) * i = mod * m
        // m 个 mod * m 相加，至多为 mod * m * m，不会超过 64 位整数最大值
        res += f[i] * i
    }
    // 保证结果非负
    return (res % mod + mod) % mod
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3]
    // Output: 10
    // Explanation:
    // All strictly increasing subsequences and their GCDs are:
    // Subsequence	GCD
    // [1]	1
    // [2]	2
    // [3]	3
    // [1,2]	1
    // [1,3]	1
    // [2,3]	1
    // [1,2,3]	1
    // Calculating beauty for each GCD:
    // GCD	Count of subsequences	Beauty (GCD × Count)
    // 1	5	1 × 5 = 5
    // 2	1	2 × 1 = 2
    // 3	1	3 × 1 = 3
    // Total beauty is 5 + 2 + 3 = 10.
    fmt.Println(totalBeauty([]int{1,2,3})) // 10
    // Example 2:
    // Input: nums = [4,6]
    // Output: 12
    // Explanation:
    // All strictly increasing subsequences and their GCDs are:
    // Subsequence	GCD
    // [4]	4
    // [6]	6
    // [4,6]	2
    // Calculating beauty for each GCD:
    // GCD	Count of subsequences	Beauty (GCD × Count)
    // 2	1	2 × 1 = 2
    // 4	1	4 × 1 = 4
    // 6	1	6 × 1 = 6
    // Total beauty is 2 + 4 + 6 = 12.
    fmt.Println(totalBeauty([]int{4,6})) // 12

    fmt.Println(totalBeauty([]int{1,2,3,4,5,6,7,8,9})) // 568
    fmt.Println(totalBeauty([]int{9,8,7,6,5,4,3,2,1})) // 45
}