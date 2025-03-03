package main

// 3469. Find Minimum Cost to Remove Array Elements
// You are given an integer array nums. 
// Your task is to remove all elements from the array by performing one of the following operations at each step until nums is empty:
//     1. Choose any two elements from the first three elements of nums and remove them.  
//        The cost of this operation is the maximum of the two elements removed.
//     2. If fewer than three elements remain in nums, remove all the remaining elements in a single operation. 
//        The cost of this operation is the maximum of the remaining elements.

// Return the minimum cost required to remove all the elements.

// Example 1:
// Input: nums = [6,2,8,4]
// Output: 12
// Explanation:
// Initially, nums = [6, 2, 8, 4].
// In the first operation, remove nums[0] = 6 and nums[2] = 8 with a cost of max(6, 8) = 8. Now, nums = [2, 4].
// In the second operation, remove the remaining elements with a cost of max(2, 4) = 4.
// The cost to remove all elements is 8 + 4 = 12. This is the minimum cost to remove all elements in nums. Hence, the output is 12.

// Example 2:
// Input: nums = [2,1,3,3]
// Output: 5
// Explanation:
// Initially, nums = [2, 1, 3, 3].
// In the first operation, remove nums[0] = 2 and nums[1] = 1 with a cost of max(2, 1) = 2. Now, nums = [3, 3].
// In the second operation remove the remaining elements with a cost of max(3, 3) = 3.
// The cost to remove all elements is 2 + 3 = 5. This is the minimum cost to remove all elements in nums. Hence, the output is 5.

// Constraints:
//     1 <= nums.length <= 1000
//     1 <= nums[i] <= 10^6

import "fmt"

func minCost(nums []int) int {
    var memo [1001][1001]int
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(int, int) int
    dfs = func(remain int, start int) int {
        if memo[remain][start] != 0 { return memo[remain][start]  }
        if start == len(nums) { return nums[remain] }
        if start == len(nums) - 1 { return max(nums[remain], nums[start]) }
        c1 := max(nums[start],  nums[start + 1]) + dfs(remain, start + 2)
        c2 := max(nums[remain], nums[start])     + dfs(start + 1, start+ 2)
        c3 := max(nums[remain], nums[start + 1]) + dfs(start, start + 2)
        memo[remain][start] = min(min(c1, c2), c3)
        return memo[remain][start]
    }
    return dfs(0, 1)
}

func minCost1(nums []int) int {
    n := len(nums)
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    if n < 3 {
        res := 0
        for _, v := range nums {
            res = max(res, v)
        }
        return res 
    }
    dp := make([][]int, 2)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    res := -1
    update := func(u int, v int) int {
        if u == -1 {
            return v
        } else {
            return min(u, v)
        }
    }
    now, pre := 0, 1 
    for i := 1; i < n; i += 2 {
        res = -1
        now, pre = pre, now
        if i + 2 > n {
            for j := 0; j < i; j ++ {
                tmp := dp[pre][j] + max(nums[j], nums[i])
                res = update(res, tmp)
            }
        } else {
            for j := 0; j < i; j ++ {
                dp[now][j] = dp[pre][j] + max(nums[i], nums[i + 1])
                res = update(res, nums[j] + dp[now][j])
                if j == 0 {
                    dp[now][i]     = dp[pre][j] + max(nums[j], nums[i + 1])
                    dp[now][i + 1] = dp[pre][j] + max(nums[j], nums[i])
                } else {
                    dp[now][i]     = min(dp[pre][j] + max(nums[j], nums[i + 1]), dp[now][i])
                    dp[now][i + 1] = min(dp[pre][j] + max(nums[j], nums[i]), dp[now][i + 1])
                }
                res = update(res, nums[i] + dp[now][i])
                res = update(res, nums[i + 1] + dp[now][i + 1])
            }
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: nums = [6,2,8,4]
    // Output: 12
    // Explanation:
    // Initially, nums = [6, 2, 8, 4].
    // In the first operation, remove nums[0] = 6 and nums[2] = 8 with a cost of max(6, 8) = 8. Now, nums = [2, 4].
    // In the second operation, remove the remaining elements with a cost of max(2, 4) = 4.
    // The cost to remove all elements is 8 + 4 = 12. This is the minimum cost to remove all elements in nums. Hence, the output is 12.
    fmt.Println(minCost([]int{6,2,8,4})) // 12
    // Example 2:
    // Input: nums = [2,1,3,3]
    // Output: 5
    // Explanation:
    // Initially, nums = [2, 1, 3, 3].
    // In the first operation, remove nums[0] = 2 and nums[1] = 1 with a cost of max(2, 1) = 2. Now, nums = [3, 3].
    // In the second operation remove the remaining elements with a cost of max(3, 3) = 3.
    // The cost to remove all elements is 2 + 3 = 5. This is the minimum cost to remove all elements in nums. Hence, the output is 5.
    fmt.Println(minCost([]int{2,1,3,3})) // 5

    fmt.Println(minCost([]int{1,2,3,4,5,6,7,8,9})) // 25
    fmt.Println(minCost([]int{9,8,7,6,5,4,3,2,1})) // 25

    fmt.Println(minCost1([]int{6,2,8,4})) // 12
    fmt.Println(minCost1([]int{2,1,3,3})) // 5
    fmt.Println(minCost1([]int{1,2,3,4,5,6,7,8,9})) // 25
    fmt.Println(minCost1([]int{9,8,7,6,5,4,3,2,1})) // 25
}