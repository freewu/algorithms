package main

// 2826. Sorting Three Groups
// You are given an integer array nums. 
// Each element in nums is 1, 2 or 3. In each operation, you can remove an element from nums. 
// Return the minimum number of operations to make nums non-decreasing.

// Example 1:
// Input: nums = [2,1,3,2,1]
// Output: 3
// Explanation:
// One of the optimal solutions is to remove nums[0], nums[2] and nums[3].

// Example 2:
// Input: nums = [1,3,2,1,3,3]
// Output: 2
// Explanation:
// One of the optimal solutions is to remove nums[1] and nums[2].

// Example 3:
// Input: nums = [2,2,2,2,3,3]
// Output: 0
// Explanation:
// nums is already non-decreasing.

// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 3

// Follow-up: Can you come up with an algorithm that runs in O(n) time complexity?

import "fmt"

func minimumOperations(nums []int) int {
    dp := make([][]int, len(nums)) 
    for i := range dp {
        dp[i] = make([]int, 4)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var helper func(i, s int) int
    helper =  func(i, s int) int {
        if i == -1 { return 0 }
        if s == -1 { return 1 << 31 }
        if dp[i][s] != -1 { return dp[i][s] }
        res := 0
        if nums[i] != s {
            res = min(helper(i, s - 1), 1 + helper(i - 1, s))
        } else {
            res = helper(i - 1, s)
        }
        dp[i][s] = res
        return res
    }
    return helper(len(nums) - 1, 3)
}

func minimumOperations1(nums []int) int {
    dp := make([]int, len(nums) + 1)
    dp[1] = nums[0]
    r := 1
    search := func(nums []int, l, r, target int) int {
        for l <= r {
            mid := (l + r) / 2
            if nums[mid] < target {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
        return l
    }
    for i := 1 ; i < len(nums) ; i ++ {
        index := search(dp, 1, r, nums[i] + 1)
        dp[index] = nums[i]
        if index == r + 1 {
            r = r + 1
        }
    }
    return len(nums) - r
}

func minimumOperations2(nums []int) int {
    f := make([]int, 3)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    for _, v := range nums {
        g := make([]int, 3)
        if v == 1 {
            g[0], g[1], g[2] = f[0], min(f[0], f[1]) + 1, min(f[0], min(f[1], f[2])) + 1
        } else if v == 2 {
            g[0], g[1], g[2] = f[0] + 1, min(f[0], f[1]), min(f[0], min(f[1], f[2])) + 1
        } else {
            g[0], g[1], g[2] = f[0] + 1,  min(f[0], f[1]) + 1, min(f[0], min(f[1], f[2]))
        }
        f = g
    }
    return min(f[0], min(f[1], f[2]))
}

func main() {
    // Example 1:
    // Input: nums = [2,1,3,2,1]
    // Output: 3
    // Explanation:
    // One of the optimal solutions is to remove nums[0], nums[2] and nums[3].
    fmt.Println(minimumOperations([]int{2,1,3,2,1})) // 3
    // Example 2:
    // Input: nums = [1,3,2,1,3,3]
    // Output: 2
    // Explanation:
    // One of the optimal solutions is to remove nums[1] and nums[2].
    fmt.Println(minimumOperations([]int{1,3,2,1,3,3})) // 2
    // Example 3:
    // Input: nums = [2,2,2,2,3,3]
    // Output: 0
    // Explanation:
    // nums is already non-decreasing.
    fmt.Println(minimumOperations([]int{2,2,2,2,3,3})) // 0

    fmt.Println(minimumOperations1([]int{2,1,3,2,1})) // 3
    fmt.Println(minimumOperations1([]int{1,3,2,1,3,3})) // 2
    fmt.Println(minimumOperations1([]int{2,2,2,2,3,3})) // 0

    fmt.Println(minimumOperations2([]int{2,1,3,2,1})) // 3
    fmt.Println(minimumOperations2([]int{1,3,2,1,3,3})) // 2
    fmt.Println(minimumOperations2([]int{2,2,2,2,3,3})) // 0
}