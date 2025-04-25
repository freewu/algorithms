package main

// 2444. Count Subarrays With Fixed Bounds
// You are given an integer array nums and two integers minK and maxK.
// A fixed-bound subarray of nums is a subarray that satisfies the following conditions:
//     The minimum value in the subarray is equal to minK.
//     The maximum value in the subarray is equal to maxK.

// Return the number of fixed-bound subarrays.
// A subarray is a contiguous part of an array.

// Example 1:
// Input: nums = [1,3,5,2,7,5], minK = 1, maxK = 5
// Output: 2
// Explanation: The fixed-bound subarrays are [1,3,5] and [1,3,5,2].

// Example 2:
// Input: nums = [1,1,1,1], minK = 1, maxK = 1
// Output: 10
// Explanation: Every subarray of nums is a fixed-bound subarray. There are 10 possible subarrays.
 
// Constraints:
//     2 <= nums.length <= 10^5
//     1 <= nums[i], minK, maxK <= 10^6

import "fmt"

// dp
func countSubarrays(nums []int, minK int, maxK int) int64 {
    // 以条件A表示最大值为maxK，条件B表示最小值为minK
    // dp[i][0]表示以nums[i]结尾的子数组中满足条件A和B的个数
    // dp[i][1]表示以nums[i]结尾的子数组中只满足条件A的个数
    // dp[i][2]表示以nums[i]结尾的子数组中只满足条件B的个数
    // dp[i][3]表示以nums[i]结尾的子数组中都不满足条件A和B，且没超出[minK,maxK]的个数
    dp, res := make([][]int, len(nums)), 0
    dp[0] = make([]int, 4)
    if minK == nums[0] && nums[0] == maxK {
        dp[0][0] = 1
    } else if nums[0] == maxK {
        dp[0][1] = 1
    } else if nums[0] == minK {
        dp[0][2] = 1
    } else if minK < nums[0] && nums[0] < maxK {
        dp[0][3] = 1
    }
    res += dp[0][0]
    for i := 1; i < len(nums); i++ {
        dp[i] = make([]int, 4)
        if nums[i] < minK || nums[i] > maxK {
            continue
        }
        if minK == nums[i] && nums[i] == maxK {
            dp[i][0] = dp[i-1][0] + dp[i-1][1] + dp[i-1][2] + dp[i-1][3] + 1
        } else if nums[i] == maxK {
            dp[i][0] = dp[i-1][0] + dp[i-1][2]
            dp[i][1] = dp[i-1][1] + dp[i-1][3] + 1
        } else if nums[i] == minK {
            dp[i][0] = dp[i-1][0] + dp[i-1][1]
            dp[i][2] = dp[i-1][2] + dp[i-1][3] + 1
        } else if minK < nums[i] && nums[i] < maxK {
            dp[i][0] = dp[i-1][0]
            dp[i][1] = dp[i-1][1]
            dp[i][2] = dp[i-1][2]
            dp[i][3] = dp[i-1][3] + 1
        }
        res += dp[i][0]
    }
    return int64(res)
}

// 
func countSubarrays1(nums []int, minK int, maxK int) int64 {
    res, mn, mx, begin, l, n := 0, 0, 0, 0, 0, len(nums)
    for i := 0; i < n; i++ {
        if nums[i] > maxK || nums[i] < minK {
            mn, mx, begin, l = 0, 0, i + 1, i + 1
        }
        if nums[i] == minK {
            mn++
        }
        if nums[i] == maxK {
            mx++
        }
        for l < n && ((nums[l] != minK && nums[l] != maxK) || 
                     (nums[l] == minK && mn > 1) || 
                     (nums[l] == maxK && mx > 1)) {
            if nums[l] == minK && mn > 1 {
                mn--
            }
            if nums[l] == maxK && mx > 1 {
                mx--
            }
            l++
        }
        if mn > 0 && mx > 0 {
            res += (l - begin + 1)
        }
    }
    return int64(res)
}

func countSubarrays2(nums []int, minK int, maxK int) int64 {
    res, border, mn, mx := 0, -1, -1, -1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for i := 0; i < len(nums); i++ {
        if nums[i] < minK || nums[i] > maxK {
            mx, mn, border = -1, -1, i
        }
        if nums[i] == minK {
            mn = i
        }
        if nums[i] == maxK {
            mx = i
        }
        if mn != -1 && mx != -1 {
            res += (min(mn, mx) - border)
        }
    }
    return int64(res)
}

func main() {
    // Example 1:
    // Input: nums = [1,3,5,2,7,5], minK = 1, maxK = 5
    // Output: 2
    // Explanation: The fixed-bound subarrays are [1,3,5] and [1,3,5,2].
    fmt.Println(countSubarrays([]int{1,3,5,2,7,5},1,5)) // 2
    // Example 2:
    // Input: nums = [1,1,1,1], minK = 1, maxK = 1
    // Output: 10
    // Explanation: Every subarray of nums is a fixed-bound subarray. There are 10 possible subarrays.
    fmt.Println(countSubarrays([]int{1,1,1,1}, 1,1)) // 10

    fmt.Println(countSubarrays([]int{1,2,3,4,5,6,7,8,9}, 1,1)) // 1
    fmt.Println(countSubarrays([]int{9,8,7,6,5,4,3,2,1}, 1,1)) // 1

    fmt.Println(countSubarrays1([]int{1,3,5,2,7,5},1,5)) // 2
    fmt.Println(countSubarrays1([]int{1,1,1,1}, 1,1)) // 10
    fmt.Println(countSubarrays1([]int{1,2,3,4,5,6,7,8,9}, 1,1)) // 1
    fmt.Println(countSubarrays1([]int{9,8,7,6,5,4,3,2,1}, 1,1)) // 1

    fmt.Println(countSubarrays2([]int{1,3,5,2,7,5},1,5)) // 2
    fmt.Println(countSubarrays2([]int{1,1,1,1}, 1,1)) // 10
    fmt.Println(countSubarrays2([]int{1,2,3,4,5,6,7,8,9}, 1,1)) // 1
    fmt.Println(countSubarrays2([]int{9,8,7,6,5,4,3,2,1}, 1,1)) // 1
}