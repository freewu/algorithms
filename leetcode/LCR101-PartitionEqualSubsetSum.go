package main

// LCR 101. 分割等和子集
// 给定一个非空的正整数数组 nums ，请判断能否将这些数字分成元素和相等的两部分。

// 示例 1：
// 输入：nums = [1,5,11,5]
// 输出：true
// 解释：nums 可以分割成 [1, 5, 5] 和 [11] 。

// 示例 2：
// 输入：nums = [1,2,3,5]
// 输出：false
// 解释：nums 不可以分为和相等的两部分

// 提示：
//     1 <= nums.length <= 200
//     1 <= nums[i] <= 100

import "fmt"

//  dp 
func canPartition(nums []int) bool {
    sum := 0
    for _, v := range nums { // 累加
        sum += v
    }
    if sum % 2 == 1 { // 积为奇数直接返回
        return false
    }
    half := sum / 2
    dp := make([]bool,half, half) // array to mark reachable numbers
    dp[0] = true
    for _, n := range nums {
        if n <= half {   // to skip too big numbers
            if dp[half - n] == true { // we found our sum
                return true
            }      
            for j:= half - n - 1; j >= 0; j-- { // we loop in opposite direction, because we don't want to check index and then loop over it
                if dp[j] == true  {
                    dp[j+n] = true
                }
            }    
        }
    }          
    return false
}

func main() {
    // Explanation: The array can be partitioned as [1, 5, 5] and [11].
    fmt.Println(canPartition([]int{1,5,11,5})) // true
    // Explanation: The array cannot be partitioned into equal sum subsets.
    fmt.Println(canPartition([]int{1,2,3,5})) // false 
    fmt.Println(canPartition([]int{1,2,3,4,5})) // false 
    fmt.Println(canPartition([]int{6,2,3,4,5})) // true 
}