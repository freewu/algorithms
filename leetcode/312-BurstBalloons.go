package main

// 312. Burst Balloons
// You are given n balloons, indexed from 0 to n - 1. 
// Each balloon is painted with a number on it represented by an array nums. 
// You are asked to burst all the balloons.

// If you burst the ith balloon, you will get nums[i - 1] * nums[i] * nums[i + 1] coins. 
// If i - 1 or i + 1 goes out of bounds of the array, then treat it as if there is a balloon with a 1 painted on it.

// Return the maximum coins you can collect by bursting the balloons wisely.

// Example 1:
// Input: nums = [3,1,5,8]
// Output: 167
// Explanation:
// nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
// coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167

// Example 2:
// Input: nums = [1,5]
// Output: 10
 
// Constraints:
//     n == nums.length
//     1 <= n <= 300
//     0 <= nums[i] <= 100

import "fmt"

// Top down approach to Dynamic Programmin.
func maxCoins(nums []int) int {
    // Sentinel elements at the beginning and the end
    nums = append([]int{1}, nums...)
    nums = append(nums, 1)

    memo := make(map[[2]int]int)
    var dp func(int, int) int
    dp = func(left, right int) int {
        if left > right {
            return 0
        }
        if v, ok := memo[[2]int{left,right}]; ok {
            return v
        }
        res := 0
        for i := left; i <= right; i++ {
            gain := nums[left-1] * nums[i] * nums[right+1]
            gain += dp(left, i-1) + dp(i+1, right)
            if gain > res {
                res = gain
            }
        }
        memo[[2]int{left, right}] = res
        return res
    }
    return dp(1, len(nums)-2)
}

// dp
func maxCoins1(nums []int) int {
    nums = append([]int{1}, nums...)
    nums = append(nums, 1)
    dp := make([][]int, len(nums))
    for i := 0; i < len(nums); i++ {
        dp[i] = make([]int, len(nums))
    }
    for i := 1; i < len(nums); i++ {
        for j := 0; j < len(nums)-2; j++ {
            right := i + j + 1
            if right >= len(nums) {
                break
            }
            mx := 0
            for k := j + 1; k < right; k++ {
                tmp := dp[j][k] + nums[j] * nums[k] * nums[right] + dp[k][right]
                if tmp > mx {
                    mx = tmp
                }
            }
            dp[j][j+i+1] = mx
        }
    }
    return dp[0][len(nums)-1]
}

func main() {
    // nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
    // coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
    fmt.Println(maxCoins([]int{3,1,5,8})) // 167
    fmt.Println(maxCoins([]int{1,5})) // 10

    fmt.Println(maxCoins1([]int{3,1,5,8})) // 167
    fmt.Println(maxCoins1([]int{1,5})) // 10
}