package main

// 1099. Two Sum Less Than K
// Given an array nums of integers and integer k, 
// return the maximum sum such that there exists i < j with nums[i] + nums[j] = sum and sum < k. 
// If no i, j exist satisfying this equation, return -1.

// Example 1:
// Input: nums = [34,23,1,24,75,33,54,8], k = 60
// Output: 58
// Explanation: We can use 34 and 24 to sum 58 which is less than 60.

// Example 2:
// Input: nums = [10,20,30], k = 15
// Output: -1
// Explanation: In this case it is not possible to get a pair sum less that 15.
 
// Constraints:
//     1 <= nums.length <= 100
//     1 <= nums[i] <= 1000
//     1 <= k <= 2000

import "fmt"
import "sort"

// sort + two pointer
func twoSumLessThanK(nums []int, k int) int {
    sort.Ints(nums)
    res := -1
    max := func (x, y int) int { if x > y { return x; }; return y; }
    i , j := 0, len(nums) - 1
    for i < j {
        sum := nums[i] + nums[j]
        if sum >= k { // 右指针左移使和变小 
            j-- 
        } else { // 左指针右移使和变大
            res = max(res, sum)
            i++
        }
    }
    return res
}

func main() {
    // Explanation: We can use 34 and 24 to sum 58 which is less than 60.
    fmt.Println(twoSumLessThanK([]int{34,23,1,24,75,33,54,8}, 60)) // 58
    // Explanation: In this case it is not possible to get a pair sum less that 15.
    fmt.Println(twoSumLessThanK([]int{10,20,30}, 15)) // -1

    fmt.Println(twoSumLessThanK([]int{1,2,3,4,5,6,7,8,9}, 15)) // 14
    fmt.Println(twoSumLessThanK([]int{9,8,7,6,5,4,3,2,1}, 15)) // 14
}