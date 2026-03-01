package main 

// 3859. Count Subarrays With K Distinct Integers
// You are given an integer array nums and two integers k and m.

// Return an integer denoting the count of subarrays of nums such that:
//     1. The subarray contains exactly k distinct integers.
//     2. Within the subarray, each distinct integer appears at least m times.

// Example 1:
// Input: nums = [1,2,1,2,2], k = 2, m = 2
// Output: 2
// Explanation:
// The possible subarrays with k = 2 distinct integers, each appearing at least m = 2 times are:
// Subarray        | Distinct numbers      | Frequency
// [1, 2, 1, 2]	| {1, 2} → 2            | {1: 2, 2: 2}
// [1, 2, 1, 2, 2]	| {1, 2} → 2            | {1: 2, 2: 3}
// Thus, the answer is 2.

// Example 2:
// Input: nums = [3,1,2,4], k = 2, m = 1
// Output: 3
// Explanation:
// The possible subarrays with k = 2 distinct integers, each appearing at least m = 1 times are:
// Subarray | Distinct numbers      | Frequency
// [3, 1]   | {3, 1} → 2            | {3: 1, 1: 1}
// [1, 2]   | {1, 2} → 2            | {1: 1, 2: 1}
// [2, 4]   | {2, 4} → 2            | {2: 1, 4: 1}
// Thus, the answer is 3.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     1 <= k, m <= nums.length

import "fmt"

func countSubarrays(nums []int, k, m int) int64 {
    calc := func(distinctLimit int) int64 {
        cnt := map[int]int{}
        geM := 0 // 窗口中的出现次数 >= m 的元素个数
        res, left := 0, 0
        for _, x := range nums {
            // 1. 入
            cnt[x]++
            if cnt[x] == m {
                geM++
            }
            // 2. 出
            for len(cnt) >= distinctLimit && geM >= k {
                out := nums[left]
                if cnt[out] == m {
                    geM--
                }
                cnt[out]--
                if cnt[out] == 0 {
                    delete(cnt, out)
                }
                left++
            }
            // 3. 更新答案
            res += left
        }
        return int64(res)
    }
    return calc(k) - calc(k+1)
}

func countSubarrays1(nums []int, k int, m int) int64 {
    mp := make(map[int]int)
    res, l, prefix, good, dist := 0, 0, 0, 0, 0
    for r := 0; r < len(nums); r++ {
        if mp[nums[r]] == 0 {
            dist++
        }
        mp[nums[r]]++
        if mp[nums[r]] == m {
            good++
        }
        if dist > k {
            for dist > k {
                if mp[nums[l]] == m {
                    good--
                }
                mp[nums[l]]--
                if mp[nums[l]] == 0 {
                    dist--
                }
                l++
            }
            prefix = 0
        }
        for dist == k && mp[nums[l]] > m {
            mp[nums[l]]--
            l++
            prefix++
        }
        if dist == k && good == k {
            res += (prefix + 1)
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,2,1,2,2], k = 2, m = 2
    // Output: 2
    // Explanation:
    // The possible subarrays with k = 2 distinct integers, each appearing at least m = 2 times are:
    // Subarray        | Distinct numbers      | Frequency
    // [1, 2, 1, 2]	| {1, 2} → 2            | {1: 2, 2: 2}
    // [1, 2, 1, 2, 2]	| {1, 2} → 2            | {1: 2, 2: 3}
    // Thus, the answer is 2.
    fmt.Println(countSubarrays([]int{1,2,1,2,2}, 2, 2)) // 2
    // Example 2:
    // Input: nums = [3,1,2,4], k = 2, m = 1
    // Output: 3
    // Explanation:
    // The possible subarrays with k = 2 distinct integers, each appearing at least m = 1 times are:
    // Subarray | Distinct numbers      | Frequency
    // [3, 1]   | {3, 1} → 2            | {3: 1, 1: 1}
    // [1, 2]   | {1, 2} → 2            | {1: 1, 2: 1}
    // [2, 4]   | {2, 4} → 2            | {2: 1, 4: 1}
    // Thus, the answer is 3.
    fmt.Println(countSubarrays([]int{3,1,2,4}, 2, 1)) // 3

    fmt.Println(countSubarrays([]int{1,2,3,4,5,6,7,8,9}, 2, 1)) // 8
    fmt.Println(countSubarrays([]int{9,8,7,6,5,4,3,2,1}, 2, 1)) // 8

    fmt.Println(countSubarrays1([]int{1,2,1,2,2}, 2, 2)) // 2
    fmt.Println(countSubarrays1([]int{3,1,2,4}, 2, 1)) // 3
    fmt.Println(countSubarrays1([]int{1,2,3,4,5,6,7,8,9}, 2, 1)) // 8
    fmt.Println(countSubarrays1([]int{9,8,7,6,5,4,3,2,1}, 2, 1)) // 8
}