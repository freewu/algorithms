package main

// 416. Partition Equal Subset Sum
// Given an integer array nums, return true if you can partition the array into two subsets such that the sum of the elements in both subsets is equal or false otherwise.

// Example 1:
// Input: nums = [1,5,11,5]
// Output: true
// Explanation: The array can be partitioned as [1, 5, 5] and [11].

// Example 2:
// Input: nums = [1,2,3,5]
// Output: false
// Explanation: The array cannot be partitioned into equal sum subsets.

// Constraints:
//     1 <= nums.length <= 200
//     1 <= nums[i] <= 100

import "fmt"

//  dp 
func canPartition(nums []int) bool {
    sum := 0
    // 累加
    for _, v := range nums {
        sum += v
    }
    // 积为奇数直接返回
    if sum % 2 == 1 {
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