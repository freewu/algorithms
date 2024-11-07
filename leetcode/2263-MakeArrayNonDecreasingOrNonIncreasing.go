package main 

// 2263. Make Array Non-decreasing or Non-increasing
// You are given a 0-indexed integer array nums. In one operation, you can:
//     Choose an index i in the range 0 <= i < nums.length
//     Set nums[i] to nums[i] + 1 or nums[i] - 1

// Return the minimum number of operations to make nums non-decreasing or non-increasing.

// Example 1:
// Input: nums = [3,2,4,5,0]
// Output: 4
// Explanation:
// One possible way to turn nums into non-increasing order is to:
// - Add 1 to nums[1] once so that it becomes 3.
// - Subtract 1 from nums[2] once so it becomes 3.
// - Subtract 1 from nums[3] twice so it becomes 3.
// After doing the 4 operations, nums becomes [3,3,3,3,0] which is in non-increasing order.
// Note that it is also possible to turn nums into [4,4,4,4,0] in 4 operations.
// It can be proven that 4 is the minimum number of operations needed.

// Example 2:
// Input: nums = [2,2,3,4]
// Output: 0
// Explanation: nums is already in non-decreasing order, so no operations are needed and we return 0.

// Example 3:
// Input: nums = [0]
// Output: 0
// Explanation: nums is already in non-decreasing order, so no operations are needed and we return 0.

// Constraints:
//     1 <= nums.length <= 1000
//     0 <= nums[i] <= 1000

// Follow up: Can you solve it in O(n*log(n)) time complexity?

import "fmt"

func convertArray(nums []int) int {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    min := func (x, y int) int { if x < y { return x; }; return y; }
    decrease := func (nums []int) int {
        dp, maximum:= make([]int, 1001), make([]int, 1001)
        for i := 1000; i >= 0; i-- {
            dp[i] = abs(nums[0] - i)
            if i == 1000 {
                maximum[i] = dp[i]
            } else {
                maximum[i] = min(dp[i], maximum[i+1])
            }
        }
        for i := 1; i < len(nums); i++ {
            for j := 1000; j >= 0; j-- {
                dp[j] = maximum[j] + abs(nums[i]-j)
                if j == 1000 {
                    maximum[j] = dp[j]
                } else {
                    maximum[j] = min(dp[j], maximum[j+1])
                }
            }
        }
        return maximum[0]
    }
    increase := func(nums []int) int {
        dp, minimum := make([]int, 1001), make([]int, 1001)
        for i := 0; i < 1001; i++ {
            dp[i] = abs(nums[0] - i)
            if i == 0 {
                minimum[i] = dp[i]
            } else {
                minimum[i] = min(dp[i], minimum[i-1])
            }
        }
        for i := 1; i < len(nums); i++ {
            for j := 0; j < 1001; j++ {
                dp[j] = minimum[j] + abs(nums[i]-j)
                if j == 0 {
                    minimum[j] = dp[j]
                } else {
                    minimum[j] = min(dp[j], minimum[j-1])
                }
            }
        }
        return minimum[1000]
    }
    return min(increase(nums), decrease(nums))
}


func main() {
    // Example 1:
    // Input: nums = [3,2,4,5,0]
    // Output: 4
    // Explanation:
    // One possible way to turn nums into non-increasing order is to:
    // - Add 1 to nums[1] once so that it becomes 3.
    // - Subtract 1 from nums[2] once so it becomes 3.
    // - Subtract 1 from nums[3] twice so it becomes 3.
    // After doing the 4 operations, nums becomes [3,3,3,3,0] which is in non-increasing order.
    // Note that it is also possible to turn nums into [4,4,4,4,0] in 4 operations.
    // It can be proven that 4 is the minimum number of operations needed.
    fmt.Println(convertArray([]int{3,2,4,5,0})) // 4
    // Example 2:
    // Input: nums = [2,2,3,4]
    // Output: 0
    // Explanation: nums is already in non-decreasing order, so no operations are needed and we return 0.
    fmt.Println(convertArray([]int{2,2,3,4})) // 0
    // Example 3:
    // Input: nums = [0]
    // Output: 0
    // Explanation: nums is already in non-decreasing order, so no operations are needed and we return 0.
    fmt.Println(convertArray([]int{0})) // 0
}