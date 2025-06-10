package main

// 3578. Count Partitions With Max-Min Difference at Most K
// You are given an integer array nums and an integer k. 
// Your task is to partition nums into one or more non-empty contiguous segments such that in each segment, 
// the difference between its maximum and minimum elements is at most k.

// Return the total number of ways to partition nums under this condition.

// Since the answer may be too large, return it modulo 10^9 + 7.

// Example 1:
// Input: nums = [9,4,1,3,7], k = 4
// Output: 6
// Explanation:
// There are 6 valid partitions where the difference between the maximum and minimum elements in each segment is at most k = 4:
// [[9], [4], [1], [3], [7]]
// [[9], [4], [1], [3, 7]]
// [[9], [4], [1, 3], [7]]
// [[9], [4, 1], [3], [7]]
// [[9], [4, 1], [3, 7]]
// [[9], [4, 1, 3], [7]]

// Example 2:
// Input: nums = [3,3,4], k = 0
// Output: 2
// Explanation:
// There are 2 valid partitions that satisfy the given conditions:
// [[3], [3], [4]]
// [[3, 3], [4]]

// Constraints:
//     2 <= nums.length <= 5 * 10^4
//     1 <= nums[i] <= 10^9
//     0 <= k <= 10^9

import "fmt"

func countPartitions(nums []int, k int) int {
    sum, left, n, mod := 0, 0, len(nums), 1_000_000_007
    f, mn, mx := make([]int, n + 1), []int{}, []int{}
    f[0] = 1
    for i, v:= range nums {
        // 1. 入
        sum += f[i]
        for len(mn) > 0 && v <= nums[mn[len(mn) - 1]] {
            mn = mn[:len(mn)-1]
        }
        mn = append(mn, i)
        for len(mx) > 0 && v >= nums[mx[len(mx) - 1]] {
            mx = mx[:len(mx) - 1]
        }
        mx = append(mx, i)
        // 2. 出
        for nums[mx[0]] - nums[mn[0]] > k {
            sum -= f[left]
            left++
            if mn[0] < left {
                mn = mn[1:]
            }
            if mx[0] < left {
                mx = mx[1:]
            }
        }
        // 3. 更新答案
        f[i+1] = sum % mod
    }
    return f[n]
}

func countPartitions1(nums []int, k int) int {
    n, j, mod := len(nums), 0, 1_000_000_007
    if n == 0 { return 0 }
    f, mn, mx := make([]int, n + 10), make([]int, 0, n), make([]int, 0, n)
    f[0] = 1
    for i := 0; i < n; i++ {
        for len(mn) > 0 && nums[i] < nums[mn[len(mn) - 1]] {
            mn = mn[:len(mn) - 1] // mn pop
        }
        mn = append(mn, i) // mn push
        for len(mx) > 0 && nums[i] > nums[mx[len(mx) - 1]] {
            mx = mx[:len(mx) - 1] // mx pop
        }
        mx = append(mx, i) // mx push
        for j <= i {
            if len(mn) == 0 || len(mx) == 0 { break }
            mnv := nums[mn[0]]
            mxv := nums[mx[0]]
            if mxv - mnv > k {
                if mn[0] == j {
                    mn = mn[1:] // mn shift
                }
                if mx[0] == j {
                    mx = mx[1:] // mx shift
                }
                j++
            } else {
                break
            }
        }
        ways := 0
        if j <= i {
            if j == 0 {
                ways = f[i]
            } else {
                ways = (f[i] - f[j - 1] + mod) % mod
            }
        }
        if i == n  -1 {
            return ways % mod
        }
        f[i + 1] = (f[i] + ways) % mod
    }
    return 0
}

func main() {
    // Example 1:
    // Input: nums = [9,4,1,3,7], k = 4
    // Output: 6
    // Explanation:
    // There are 6 valid partitions where the difference between the maximum and minimum elements in each segment is at most k = 4:
    // [[9], [4], [1], [3], [7]]
    // [[9], [4], [1], [3, 7]]
    // [[9], [4], [1, 3], [7]]
    // [[9], [4, 1], [3], [7]]
    // [[9], [4, 1], [3, 7]]
    // [[9], [4, 1, 3], [7]]
    fmt.Println(countPartitions([]int{9,4,1,3,7}, 4)) // 6
    // Example 2:
    // Input: nums = [3,3,4], k = 0
    // Output: 2
    // Explanation:
    // There are 2 valid partitions that satisfy the given conditions:
    // [[3], [3], [4]]
    // [[3, 3], [4]]
    fmt.Println(countPartitions([]int{3,3,4}, 0)) // 2

    fmt.Println(countPartitions([]int{1,2,3,4,5,6,7,8,9}, 0)) // 1
    fmt.Println(countPartitions([]int{9,8,7,6,5,4,3,2,1}, 0)) // 1

    fmt.Println(countPartitions1([]int{9,4,1,3,7}, 4)) // 6
    fmt.Println(countPartitions1([]int{3,3,4}, 0)) // 2
    fmt.Println(countPartitions1([]int{1,2,3,4,5,6,7,8,9}, 0)) // 1
    fmt.Println(countPartitions1([]int{9,8,7,6,5,4,3,2,1}, 0)) // 1
}