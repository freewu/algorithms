package main

// 3953. Maximum Score with Co-Prime Element
// You are given an integer array nums of length n and an integer maxVal.

// You may change any element in nums to any positive integer less than or equal to maxVal. Each such change costs 1.

// Two integers are co-prime if their greatest common divisor (GCD) is 1.

// After all modifications, you must choose an index i such that, nums[i] is co-prime with every other element nums[j].

// Let:
//     1. selectedValue be the final value of nums[i] after modifications.
//     2. modificationCost be the total number of elements changed.

// The score is defined as score = selectedValue - modificationCost.

// Return the maximum possible score.

// Example 1:
// Input: nums = [3,4,6], maxVal = 5
// Output: 4
// Explanation:
// Change nums[2] from 6 to 5, which costs 1. Choose nums[2] = 5, since it is co-prime with 3 and 4.
// selectedValue = 5
// modificationCost = 1
// The score is 5 - 1 = 4

// Example 2:
// Input: nums = [1,2,3], maxVal = 4
// Output: 3
// Explanation:
// No modifications are required. Choose nums[2] = 3, since it is co-prime with 1 and 2.
// selectedValue = 3
// modificationCost = 0
// The score is 3 - 0 = 3

// Example 3:
// Input: nums = [2,2], maxVal = 1
// Output: 1
// Explanation:
// Change nums[0] from 2 to 1, which costs 1. Choose nums[1] = 2, since it is co-prime with 1.
// selectedValue = 2
// modificationCost = 1
// The score is ​​​​​​​2 - 1 = 1
 
// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     1 <= maxVal <= 10​^​​​​​​5

import "fmt"
import "slices"

const MX = 100_001
var mobius = [MX]int{1: 1}     // 莫比乌斯函数
var divisors = [MX][]int{} // 不含平方因子的因子列表，用于容斥

func init() {
    // 预处理莫比乌斯函数
    // 当 n > 1 时，sum_{d|n} mu[d] = 0
    // 所以 mu[n] = -sum_{d|n ∧ d<n} mu[d]
    for i := 1; i < MX; i++ {
        for j := i * 2; j < MX; j += i {
            mobius[j] -= mobius[i] // i 是 j 的真因子
        }
    }
    // 预处理不含平方因子的因子列表
    // 本题不需要因子 1
    for i := 2; i < MX; i++ {
        if mobius[i] == 0 {
            continue
        }
        for j := i; j < MX; j += i {
            divisors[j] = append(divisors[j], i) // i 是 j 的因子，且 mobius[i] != 0
        }
    }
}

func maxScore(nums []int, maxVal int) int {
    res, mx := 0, slices.Max(nums)
    count := make([]int, mx + 1)
    for _, x := range nums {
        count[x]++
    }
    multi := make([]int, mx +1)
    for i := 2; i <= mx; i++ {
        for j := i; j <= mx; j += i {
            multi[i] += count[j] // 统计 nums 中有多少个数是 i 的倍数
        }
    }
    if count[1] > 0 {
        res = 1 // selectedValue = 1 时，无需修改，得分为 1
    }
    // 从大到小枚举 selectedValue
    // 优化：如果 selectedValue <= res，那么 res 不会变大，跳出循环
    for val := max(mx, maxVal); val > res; val-- {
        if val > maxVal && count[val] == 0 {
            continue // 无法改成 val
        }
        // 与 val 不互质的数，其中一个数改成 val，其余数都改成 1
        cost := 0
        for _, d := range divisors[val] {
            if d > mx {
                break
            }
            cost -= mobius[d] * multi[d]
        }
        if val <= mx && count[val] > 0 {
            cost-- // 如果某个 nums[i] 恰好等于 val，可以少改一次
        } else if cost == 0 {
            cost = 1 // 至少要有一个数改成 val
        }
        res = max(res, val - cost)
    }
    return res
}

func maxScore1(nums []int, maxVal int) int {
    res, mx, n := 0, maxVal, len(nums)
    for i := 0; i < n; i++ {
        if nums[i] > mx {
            mx = nums[i]
        }
    }
    freq := make([]int, mx + 1)
    for _, v := range nums {
        freq[v]++
    }
    primes, mu, lp := make([]int, 0), make([]int, mx + 1), make([]int, mx+1)
    mu[1] = 1
    for i := 2; i <= mx; i++ {
        if lp[i] == 0 {
            lp[i] = i
            primes = append(primes, i)
            mu[i] = -1
        }
        for _, p := range primes {
            v := i * p
            if v > mx {
                break
            }
            lp[v] = p
            if i%p == 0 {
                mu[v] = 0
                break
            } else {
                mu[v] = -mu[i]
            }
        }
    }
    multCnt := make([]int, mx + 1)
    for d := 1; d <= mx; d++ {
        for k := d; k <= mx; k += d {
            multCnt[d] += freq[k]
        }
    }
    cop := make([]int, mx+ 1)
    for d := 1; d <= mx; d++ {
        if mu[d] == 0 {
            continue
        }
        add := mu[d] * multCnt[d]
        for x := d; x <= mx; x += d {
            cop[x] += add
        }
    }
    for x := 1; x <= mx; x++ {
        bad := n - cop[x]
        if freq[x] > 0 {
            if x == 1 {
                res = max(res, 1)
            } else {
                res = max(res, x-bad+1)
            }
        } else if x <= maxVal {
            if bad > 0 {
                res = max(res, x-bad)
            } else {
                res = max(res, x-1)
            }
        }
    }
    return res  
}

func main() {
    // Example 1:
    // Input: nums = [3,4,6], maxVal = 5
    // Output: 4
    // Explanation:
    // Change nums[2] from 6 to 5, which costs 1. Choose nums[2] = 5, since it is co-prime with 3 and 4.
    // selectedValue = 5
    // modificationCost = 1
    // The score is 5 - 1 = 4
    fmt.Println(maxScore([]int{3,4,6}, 5)) // 4
    // Example 2:
    // Input: nums = [1,2,3], maxVal = 4
    // Output: 3
    // Explanation:
    // No modifications are required. Choose nums[2] = 3, since it is co-prime with 1 and 2.
    // selectedValue = 3
    // modificationCost = 0
    // The score is 3 - 0 = 3
    fmt.Println(maxScore([]int{1,2,3}, 4)) // 3
    // Example 3:
    // Input: nums = [2,2], maxVal = 1
    // Output: 1
    // Explanation:
    // Change nums[0] from 2 to 1, which costs 1. Choose nums[1] = 2, since it is co-prime with 1.
    // selectedValue = 2
    // modificationCost = 1
    // The score is ​​​​​​​2 - 1 = 1
    fmt.Println(maxScore([]int{2,2}, 1)) // 1

    fmt.Println(maxScore([]int{1,2,3,4,5,6,7,8,9}, 2)) // 7
    fmt.Println(maxScore([]int{9,8,7,6,5,4,3,2,1}, 2)) // 7

    fmt.Println(maxScore1([]int{3,4,6}, 5)) // 4
    fmt.Println(maxScore1([]int{1,2,3}, 4)) // 3
    fmt.Println(maxScore1([]int{2,2}, 1)) // 1
    fmt.Println(maxScore1([]int{1,2,3,4,5,6,7,8,9}, 2)) // 7
    fmt.Println(maxScore1([]int{9,8,7,6,5,4,3,2,1}, 2)) // 7
}